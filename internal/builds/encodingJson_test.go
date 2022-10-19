package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingJson(t *testing.T) {
	suite.Run(t, new(SuiteEncodingJson))
}

type SuiteEncodingJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingJson) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingJson) target(excel string) *encodingJson {
	element := &initializeElement{
		Global: &Global{
			LineOfName:  1,
			LineOfNote:  2,
			LineOfField: 3,
			LineOfLayer: 4,
			LineOfData:  5,
		},
		Named: &nameds.Named{ExcelName: excel, SheetName: testdata.SheetData},
	}

	assert.Nil(this.T(), InitializeElement(element))

	target := &encodingJson{
		Global:     element.Global,
		Named:      element.Named,
		Json:       &nameds.Json{ExcelName: excel, SheetName: testdata.SheetData},
		excel:      element.excel,
		layoutJson: element.layoutJson,
	}
	return target
}

func (this *SuiteEncodingJson) TestEncodingJson() {
	empty, err := utils.JsonMarshal(testdata.GetExcelContentEmpty())
	assert.Nil(this.T(), err)
	data, err := utils.JsonMarshal(testdata.GetExcelContentReal())
	assert.Nil(this.T(), err)

	target := this.target(testdata.ExcelReal)
	assert.Nil(this.T(), EncodingJson(target))
	testdata.CompareFile(this.T(), target.JsonDataPath(), data)

	target = this.target(testdata.ExcelEmpty)
	assert.Nil(this.T(), EncodingJson(target))
	testdata.CompareFile(this.T(), target.JsonDataPath(), empty)

	target = this.target(testdata.ExcelInvalidData)
	assert.NotNil(this.T(), EncodingJson(target))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelReal)
		target.Json.ExcelName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingJson(target))

		target = this.target(testdata.ExcelReal)
		target.Json.SheetName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingJson(target))
	} // if
}
