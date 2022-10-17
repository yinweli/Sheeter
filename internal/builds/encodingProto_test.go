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

func TestEncodingProto(t *testing.T) {
	suite.Run(t, new(SuiteEncodingProto))
}

type SuiteEncodingProto struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingProto) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingProto) target(excel string) *encodingData {
	context := &Context{
		Config: &Config{
			Global: Global{
				ExportProto: true,
				LineOfName:  1,
				LineOfNote:  2,
				LineOfField: 3,
				LineOfLayer: 4,
				LineOfData:  5,
			},
		},
		Sector: []*ContextSector{
			{
				Element: Element{
					Excel: excel,
					Sheet: testdata.SheetData,
				},
			},
		},
	}
	sector := context.Sector[0]

	assert.Nil(this.T(), Initialize(context))
	assert.Nil(this.T(), Generate(context))

	target := &encodingData{
		Global:     &context.Global,
		Element:    &sector.Element,
		Mixed:      mixeds.NewMixed(excel, testdata.SheetData),
		excel:      sector.excel,
		layoutJson: sector.layoutJson,
	}
	return target
}

func (this *SuiteEncodingProto) TestEncodingProto() {
	target := this.target(testdata.ExcelNameReal)
	assert.Nil(this.T(), encodingProto(target))
	assert.True(this.T(), utils.FileExist(target.ProtoDataPath()))

	target = this.target(testdata.ExcelNameEmpty)
	assert.Nil(this.T(), encodingProto(target))
	assert.True(this.T(), utils.FileExist(target.ProtoDataPath()))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelNameReal)
		target.Mixed = mixeds.NewMixed(testdata.UnknownStr, target.Sheet)
		assert.NotNil(this.T(), encodingProto(target))

		target = this.target(testdata.ExcelNameReal)
		target.Mixed = mixeds.NewMixed(target.Excel, testdata.UnknownStr)
		assert.NotNil(this.T(), encodingProto(target))
	} // if
}
