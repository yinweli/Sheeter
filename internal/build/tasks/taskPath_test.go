package tasks

import (
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskPath(t *testing.T) {
	suite.Run(t, new(SuiteTaskPath))
}

type SuiteTaskPath struct {
	suite.Suite
}

func (this *SuiteTaskPath) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{
		ExcelPath: "path",
	}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteTaskPath) TestExcelFilePath() {
	assert.Equal(this.T(), "path/real.xlsx", this.target().excelFilePath())
}

func (this *SuiteTaskPath) TestJsonSchemaFilePath() {
	assert.Equal(this.T(), "schema/realData.json.schema", this.target().jsonSchemaFilePath())
}

func (this *SuiteTaskPath) TestJsonFileName() {
	assert.Equal(this.T(), "realData.json", this.target().jsonFileName())
}

func (this *SuiteTaskPath) TestJsonFilePath() {
	assert.Equal(this.T(), "json/realData.json", this.target().jsonFilePath())
}

func (this *SuiteTaskPath) TestJsonCsFilePath() {
	assert.Equal(this.T(), "jsonCs/realData.cs", this.target().jsonCsFilePath())
}

func (this *SuiteTaskPath) TestJsonCsReaderFilePath() {
	assert.Equal(this.T(), "jsonCs/realData.reader.cs", this.target().jsonCsReaderFilePath())
}

func (this *SuiteTaskPath) TestJsonGoFilePath() {
	assert.Equal(this.T(), "jsonGo/realData.go", this.target().jsonGoFilePath())
}

func (this *SuiteTaskPath) TestJsonGoReaderFilePath() {
	assert.Equal(this.T(), "jsonGo/realData.reader.go", this.target().jsonGoReaderFilePath())
}

func (this *SuiteTaskPath) TestLuaFilePath() {
	assert.Equal(this.T(), "lua/realData.lua", this.target().luaFilePath())
}

func (this *SuiteTaskPath) TestFileName() {
	assert.Equal(this.T(), "realData.test1.test2.test3", this.target().fileName("test1", "test2", "test3"))
}
