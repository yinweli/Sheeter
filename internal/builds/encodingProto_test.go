package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/nameds"
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

func (this *SuiteEncodingProto) target(excel string) *encodingProto {
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
	generate := []any{}

	for _, itor := range element.layoutType.TypeNames() {
		types := element.layoutType.Types(itor)
		depend := element.layoutDepend.Depends(itor)
		generate = append(generate, &generateProto{
			Global: element.Global,
			Named:  &nameds.Named{ExcelName: types.Excel, SheetName: types.Sheet},
			Field:  &nameds.Field{},
			Proto:  &nameds.Proto{ExcelName: types.Excel, SheetName: types.Sheet},
			Type:   types,
			Depend: depend,
		})
	} // for

	for _, itor := range generate {
		assert.Nil(this.T(), GenerateProtoSchema(itor))
	} // for

	target := &encodingProto{
		Global:     element.Global,
		Named:      element.Named,
		Proto:      &nameds.Proto{ExcelName: excel, SheetName: testdata.SheetData},
		excel:      element.excel,
		layoutData: element.layoutData,
	}
	return target
}

func (this *SuiteEncodingProto) TestEncodingProto() {
	target := this.target(testdata.ExcelReal)
	assert.Nil(this.T(), EncodingProto(target))
	assert.FileExists(this.T(), target.ProtoDataPath())

	target = this.target(testdata.ExcelEmpty)
	assert.Nil(this.T(), EncodingProto(target))
	assert.FileExists(this.T(), target.ProtoDataPath())

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelReal)
		target.Proto.ExcelName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingProto(target))

		target = this.target(testdata.ExcelReal)
		target.Proto.SheetName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingProto(target))
	} // if
}
