package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingJson(t *testing.T) {
	suite.Run(t, new(SuiteEncodingJson))
}

type SuiteEncodingJson struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteEncodingJson) SetupSuite() {
	this.Change("test-encodingJson")
}

func (this *SuiteEncodingJson) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteEncodingJson) target(excel string) *encodingJson {
	sheet := &initializeSheetData{
		Global: &Global{
			LineOfTag:   1,
			LineOfName:  2,
			LineOfNote:  3,
			LineOfField: 4,
			LineOfLayer: 5,
			LineOfData:  6,
			Tags:        "A",
		},
		Named: &nameds.Named{ExcelName: excel, SheetName: testdata.SheetData},
	}
	result := make(chan any, 1)

	assert.Nil(this.T(), InitializeSheetData(sheet, result))
	assert.Empty(this.T(), result)
	assert.NotNil(this.T(), sheet.excel)
	assert.NotNil(this.T(), sheet.layoutData)
	assert.NotNil(this.T(), sheet.layoutType)
	assert.NotNil(this.T(), sheet.layoutDepend)

	target := &encodingJson{
		Global:     sheet.Global,
		Named:      sheet.Named,
		Json:       &nameds.Json{ExcelName: sheet.ExcelName, SheetName: sheet.SheetName},
		excel:      sheet.excel,
		layoutData: sheet.layoutData,
	}
	return target
}

func (this *SuiteEncodingJson) TestEncodingJson() {
	emptyBytes, _ := utils.JsonMarshal(map[string]interface{}{
		"Datas": map[sheeter.PkeyType]interface{}{},
	})
	realBytes, _ := utils.JsonMarshal(map[string]interface{}{
		"Datas": map[sheeter.PkeyType]interface{}{
			1: map[string]interface{}{
				"Name0": 1,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
					},
					"Name1": true,
				},
			},
			2: map[string]interface{}{
				"Name0": 2,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
					},
					"Name1": false,
				},
			},
			3: map[string]interface{}{
				"Name0": 3,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
					},
					"Name1": true,
				},
			},
		},
	})

	target := this.target(testdata.ExcelReal)
	assert.Nil(this.T(), EncodingJson(target, nil))
	testdata.CompareFile(this.T(), target.JsonDataPath(), realBytes)

	target = this.target(testdata.ExcelEmpty)
	assert.Nil(this.T(), EncodingJson(target, nil))
	testdata.CompareFile(this.T(), target.JsonDataPath(), emptyBytes)

	target = this.target(testdata.ExcelInvalidData)
	assert.NotNil(this.T(), EncodingJson(target, nil))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelReal)
		target.Json.ExcelName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingJson(target, nil))

		target = this.target(testdata.ExcelReal)
		target.Json.SheetName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingJson(target, nil))
	} // if
}
