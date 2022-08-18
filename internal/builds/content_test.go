package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vbauerster/mpb/v7"

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
		Path:        testdata.RootPath,
		Bom:         true,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
		Progress:    mpb.New(mpb.WithOutput(nil)),
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

	target = this.target()
	target.Progress = nil
	assert.NotNil(this.T(), target.Check())
}

func (this *SuiteContent) TestExcelFilePath() {
	assert.Equal(this.T(), "testdata/real.xlsx", this.target().ExcelFilePath())
}

func (this *SuiteContent) TestSchemaFilePath() {
	assert.Equal(this.T(), "schema/realData.schema", this.target().SchemaFilePath())
}

func (this *SuiteContent) TestJsonFileName() {
	assert.Equal(this.T(), "realData.json", this.target().JsonFileName())
}

func (this *SuiteContent) TestJsonFilePath() {
	assert.Equal(this.T(), "json/realData.json", this.target().JsonFilePath())
}

func (this *SuiteContent) TestJsonCsFilePath() {
	assert.Equal(this.T(), "json-cs/realData.cs", this.target().JsonCsFilePath())
}

func (this *SuiteContent) TestJsonCsReaderFilePath() {
	assert.Equal(this.T(), "json-cs/realData.reader.cs", this.target().JsonCsReaderFilePath())
}

func (this *SuiteContent) TestJsonGoFilePath() {
	assert.Equal(this.T(), "json-go/realData.go", this.target().JsonGoFilePath())
}

func (this *SuiteContent) TestJsonGoReaderFilePath() {
	assert.Equal(this.T(), "json-go/realData.reader.go", this.target().JsonGoReaderFilePath())
}

func (this *SuiteContent) TestNamespace() {
	assert.Equal(this.T(), "sheeter", this.target().Namespace())
}

func (this *SuiteContent) TestTargetName() {
	assert.Equal(this.T(), "real.xlsx(Data)", this.target().TargetName())
}

func (this *SuiteContent) TestStructName() {
	assert.Equal(this.T(), "RealData", this.target().StructName())
}

func (this *SuiteContent) TestReaderName() {
	assert.Equal(this.T(), "RealDataReader", this.target().ReaderName())
}

func (this *SuiteContent) TestExcelName() {
	assert.Equal(this.T(), "real", this.target().ExcelName())
}

func (this *SuiteContent) TestFileName() {
	assert.Equal(this.T(), "realData.test1.test2.test3", this.target().fileName("test1", "test2", "test3"))
}

func (this *SuiteContent) TestGetRows() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	rows, err := target.getRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = target.getRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = target.getRows(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.getRows(1)
	assert.NotNil(this.T(), err)

	target.close()
}

func (this *SuiteContent) TestGetColumns() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := target.getColumns(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"name0#pkey", "empty#empty", "name1#bool", "name2#int", "name3#text", "name2#int", "name3#text", "name2#int", "name3#text"}, cols)

	_, err = target.getColumns(10)
	assert.NotNil(this.T(), err)

	_, err = target.getColumns(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.getColumns(1)
	assert.NotNil(this.T(), err)

	target.close()
}
