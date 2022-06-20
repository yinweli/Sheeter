package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskPath(t *testing.T) {
	task := mockTaskPath()
	assert.Equal(t, "path/real.xlsx", task.excelFilePath())
	assert.Equal(t, "schema/realData.proto", task.schemaProtoFilePath())
	assert.Equal(t, "realData.json", task.jsonFileName())
	assert.Equal(t, "json/realData.json", task.jsonFilePath())
	assert.Equal(t, "realData.bytes", task.bytesFileName())
	assert.Equal(t, "bytes/realData.bytes", task.bytesFilePath())
	assert.Equal(t, "lua/realData.lua", task.luaFilePath())
	assert.Equal(t, "jsonCs/realData.cs", task.jsonCsFilePath())
	assert.Equal(t, "jsonGo/realData.go", task.jsonGoFilePath())
	assert.Equal(t, "protoCs/realData.cs", task.protoCsFilePath())
	assert.Equal(t, "protoGo/realData.go", task.protoGoFilePath())
	assert.Equal(t, "realData.test", task.fileName("test"))
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
