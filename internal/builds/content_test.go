package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestContent(t *testing.T) {
	suite.Run(t, new(SuiteContent))
}

type SuiteContent struct {
	suite.Suite
}

func (this *SuiteContent) target() *Content {
	target := &Content{
		Bom:         true,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.Path(testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteContent) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.Check())

	target = this.target()
	target.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfLayer = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfField = 4
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfLayer = 4
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfNote = 4
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Excel = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Sheet = ""
	assert.NotNil(this.T(), target.Check())
}

func (this *SuiteContent) TestShowName() {
	assert.Equal(this.T(), "real#data", this.target().ShowName())
}

func (this *SuiteContent) TestSchemaPath() {
	assert.Equal(this.T(), filepath.Join(pathSchema, "realData.schema"), this.target().SchemaPath())
}

func (this *SuiteContent) TestJsonPath() {
	assert.Equal(this.T(), filepath.Join(pathJson, "realData.json"), this.target().JsonPath())
}

func (this *SuiteContent) TestJsonCsPath() {
	assert.Equal(this.T(), filepath.Join(pathJsonCs, "realData", "realData.cs"), this.target().JsonCsPath())
}

func (this *SuiteContent) TestJsonCsReaderPath() {
	assert.Equal(this.T(), filepath.Join(pathJsonCs, "realData", "realDataReader.cs"), this.target().JsonCsReaderPath())
}

func (this *SuiteContent) TestJsonGoPath() {
	assert.Equal(this.T(), filepath.Join(pathJsonGo, "realData", "realData.go"), this.target().JsonGoPath())
}

func (this *SuiteContent) TestJsonGoReaderPath() {
	assert.Equal(this.T(), filepath.Join(pathJsonGo, "realData", "realDataReader.go"), this.target().JsonGoReaderPath())
}

func (this *SuiteContent) TestAppName() {
	assert.Equal(this.T(), "sheeter", this.target().AppName())
}

func (this *SuiteContent) TestNamespace() {
	assert.Equal(this.T(), "realdata", this.target().Namespace())
}

func (this *SuiteContent) TestStructName() {
	assert.Equal(this.T(), "RealData", this.target().StructName())
}

func (this *SuiteContent) TestReaderName() {
	assert.Equal(this.T(), "RealDataReader", this.target().ReaderName())
}

func (this *SuiteContent) TestGetRows() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	rows, err := target.GetRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = target.GetRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = target.GetRows(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.GetRows(1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteContent) TestGetColumns() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := target.GetColumns(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{"name0#pkey", "empty#empty", "name1#bool", "name2#int", "name3#text", "name2#int", "name3#text",
			"name2#int", "name3#text"}, cols)

	_, err = target.GetColumns(10)
	assert.NotNil(this.T(), err)

	_, err = target.GetColumns(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.GetColumns(1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteContent) TestFileName() {
	target := this.target()

	assert.Equal(this.T(), "Realdata", target.combine(params{excelUpper: true}))
	assert.Equal(this.T(), "realdata", target.combine(params{excelUpper: false}))
	assert.Equal(this.T(), "realData", target.combine(params{sheetUpper: true}))
	assert.Equal(this.T(), "realdata", target.combine(params{sheetUpper: false}))
	assert.Equal(this.T(), "real#data", target.combine(params{middle: "#"}))
	assert.Equal(this.T(), "realdata#", target.combine(params{last: "#"}))
	assert.Equal(this.T(), "realdata.#", target.combine(params{ext: "#"}))
}
