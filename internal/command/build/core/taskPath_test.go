package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskPath(t *testing.T) {
	task := mockTaskPath()
	assert.Equal(t, "path/real.xlsx", task.excelFilePath())
	assert.Equal(t, "schema/realData.json.schema", task.jsonSchemaFilePath())
	assert.Equal(t, "realData.json", task.jsonFileName())
	assert.Equal(t, "json/realData.json", task.jsonFilePath())
	assert.Equal(t, "jsonCs/realData.cs", task.jsonCsFilePath())
	assert.Equal(t, "jsonCs/realData.reader.cs", task.jsonCsReaderFilePath())
	assert.Equal(t, "jsonGo/realData.go", task.jsonGoFilePath())
	assert.Equal(t, "lua/realData.lua", task.luaFilePath())
	assert.Equal(t, "realData.test1.test2.test3", task.fileName("test1", "test2", "test3"))
}

func mockTaskPath() *Task {
	return &Task{
		global: &Global{
			ExcelPath: "path",
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
