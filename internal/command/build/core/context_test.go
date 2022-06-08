package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	context := mockContext()
	assert.Equal(t, "excel.xlsx(sheet)", context.LogName())
	assert.Equal(t, "path/excel.xlsx", context.ExcelFilePath())
	assert.Equal(t, "excelSheet.json", context.JsonFileName())
	assert.Equal(t, "json/excelSheet.json", context.JsonFilePath())
	assert.Equal(t, "excelSheet.hpp", context.JsonCppFileName())
	assert.Equal(t, "jsonCpp/excelSheet.hpp", context.JsonCppFilePath())
	assert.Equal(t, "excelSheet.cs", context.JsonCsFileName())
	assert.Equal(t, "jsonCs/excelSheet.cs", context.JsonCsFilePath())
	assert.Equal(t, "excelSheet.go", context.JsonGoFileName())
	assert.Equal(t, "jsonGo/excelSheet.go", context.JsonGoFilePath())
	assert.Equal(t, "excelSheet.proto", context.ProtoFileName())
	assert.Equal(t, "proto/excelSheet.proto", context.ProtoFilePath())
	assert.Equal(t, "excelSheet.bytes", context.ProtoBytesFileName())
	assert.Equal(t, "proto/excelSheet.bytes", context.ProtoBytesFilePath())
	assert.Equal(t, "protoCpp/excelSheet.hpp", context.ProtoCppFilePath())
	assert.Equal(t, "excelSheet.hpp", context.ProtoCppFileName())
	assert.Equal(t, "protoCs/excelSheet.cs", context.ProtoCsFilePath())
	assert.Equal(t, "excelSheet.go", context.ProtoGoFileName())
	assert.Equal(t, "protoGo/excelSheet.go", context.ProtoGoFilePath())
	assert.Equal(t, "ExcelSheet", context.StructName())
	assert.Equal(t, "Sheeter", context.CppNamespace())
	assert.Equal(t, "Sheeter", context.CsNamespace())
	assert.Equal(t, "sheeter", context.GoPackage())
	assert.Equal(t, "excelSheet.test", context.fileName("test"))
	assert.Equal(t, "excel", context.excelName())
}

func TestSheets(t *testing.T) {
	context := mockContext()
	assert.Equal(t, 12, context.Sheets.Size())
}

func TestColumn(t *testing.T) {
	context := mockContext()
	assert.Equal(t, "Name0", context.Columns[0].ColumnName())
	assert.Equal(t, "Name1", context.Columns[1].ColumnName())
	assert.Equal(t, "Name2", context.Columns[2].ColumnName())
	assert.Equal(t, "Name3", context.Columns[3].ColumnName())
}

func mockContext() *Context {
	return &Context{
		Global: &Global{
			ExcelPath: "path",
		},
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Sheets: Sheets{
			{"1", "2", "3", "4", "5"},
			{"1", "2", "3", "4"},
			{"1", "2", "3"},
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
			{Note: "note3", Name: "name3", Field: &FieldEmpty{}},
		},
	}
}
