package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/mixeds"
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

func (this *SuiteEncodingJson) target(excel string) *encodingData {
	context := &Context{
		Config: &Config{
			Global: Global{
				ExportJson:  true,
				LineOfField: 1,
				LineOfLayer: 2,
				LineOfNote:  3,
				LineOfData:  4,
			},
		},
		Sector: []*ContextSector{
			{
				Element: Element{
					Excel: excel,
					Sheet: testdata.SheetName,
				},
			},
		},
	}
	sector := context.Sector[0]

	assert.Nil(this.T(), initializeSector(context, sector))

	target := &encodingData{
		Global:     &context.Global,
		Element:    &sector.Element,
		Mixed:      mixeds.NewMixed(excel, testdata.SheetName),
		excel:      sector.excel,
		layoutJson: sector.layoutJson,
	}
	return target
}

func (this *SuiteEncodingJson) TestEncodingJson() {
	empty, err := utils.JsonMarshal(testdata.GetExcelContentEmpty())
	assert.Nil(this.T(), err)
	data, err := utils.JsonMarshal(testdata.GetExcelContentReal())
	assert.Nil(this.T(), err)

	target := this.target(testdata.ExcelNameReal)
	assert.Nil(this.T(), encodingJson(target))
	testdata.CompareFile(this.T(), target.JsonDataPath(), data)

	target = this.target(testdata.ExcelNameEmpty)
	assert.Nil(this.T(), encodingJson(target))
	testdata.CompareFile(this.T(), target.JsonDataPath(), empty)

	target = this.target(testdata.ExcelNameInvalidData)
	assert.NotNil(this.T(), encodingJson(target))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelNameReal)
		target.Mixed = mixeds.NewMixed(testdata.UnknownStr, target.Sheet)
		assert.NotNil(this.T(), encodingJson(target))

		target = this.target(testdata.ExcelNameReal)
		target.Mixed = mixeds.NewMixed(target.Excel, testdata.UnknownStr)
		assert.NotNil(this.T(), encodingJson(target))
	} // if
}
