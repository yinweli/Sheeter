package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingProto(t *testing.T) {
	suite.Run(t, new(SuiteEncodingProto))
}

type SuiteEncodingProto struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteEncodingProto) SetupSuite() {
	this.Change("test-encodingProto")
}

func (this *SuiteEncodingProto) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteEncodingProto) target(excel string) *encodingProto {
	sheet := &initializeSheetData{
		Global: &Global{
			LineOfName:  1,
			LineOfNote:  2,
			LineOfField: 3,
			LineOfLayer: 4,
			LineOfTag:   5,
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

	generate := []any{}

	for _, itor := range sheet.layoutType.TypeNames() {
		type_ := sheet.layoutType.Type(itor)
		generate = append(generate, &generateProto{
			Global: sheet.Global,
			Named:  &nameds.Named{ExcelName: type_.Excel, SheetName: type_.Sheet},
			Field:  &nameds.Field{},
			Proto:  &nameds.Proto{ExcelName: type_.Excel, SheetName: type_.Sheet},
			Reader: type_.Reader,
			Fields: sheet.layoutType.Fields(itor),
			Depend: sheet.layoutDepend.Depends(itor),
		})
	} // for

	for _, itor := range generate {
		assert.Nil(this.T(), GenerateProtoSchema(itor, nil))
	} // for

	target := &encodingProto{
		Global:     sheet.Global,
		Named:      sheet.Named,
		Proto:      &nameds.Proto{ExcelName: sheet.ExcelName, SheetName: sheet.SheetName},
		excel:      sheet.excel,
		layoutData: sheet.layoutData,
	}
	return target
}

func (this *SuiteEncodingProto) TestEncodingProto() {
	target := this.target(testdata.ExcelReal)
	assert.Nil(this.T(), EncodingProto(target, nil))
	assert.FileExists(this.T(), target.ProtoDataPath())

	target = this.target(testdata.ExcelEmpty)
	assert.Nil(this.T(), EncodingProto(target, nil))
	assert.FileExists(this.T(), target.ProtoDataPath())

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target(testdata.ExcelReal)
		target.Proto.ExcelName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingProto(target, nil))

		target = this.target(testdata.ExcelReal)
		target.Proto.SheetName = testdata.UnknownStr
		assert.NotNil(this.T(), EncodingProto(target, nil))
	} // if
}
