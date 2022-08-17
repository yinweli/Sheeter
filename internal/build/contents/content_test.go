package contents

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/suite"
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
		Excel:       testdata.RealExcel,
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
