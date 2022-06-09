package core

import (
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := mockContext()
	assert.Equal(t, "excel.xlsx(sheet)", ctx.LogName())
	assert.Equal(t, "path/excel.xlsx", ctx.ExcelFilePath())
	assert.Equal(t, "excelSheet.json", ctx.JsonFileName())
	assert.Equal(t, "json/excelSheet.json", ctx.JsonFilePath())
	assert.Equal(t, "excelSheet.hpp", ctx.JsonCppFileName())
	assert.Equal(t, "jsonCpp/excelSheet.hpp", ctx.JsonCppFilePath())
	assert.Equal(t, "excelSheet.cs", ctx.JsonCsFileName())
	assert.Equal(t, "jsonCs/excelSheet.cs", ctx.JsonCsFilePath())
	assert.Equal(t, "excelSheet.go", ctx.JsonGoFileName())
	assert.Equal(t, "jsonGo/excelSheet.go", ctx.JsonGoFilePath())
	assert.Equal(t, "excelSheet.proto", ctx.ProtoFileName())
	assert.Equal(t, "proto/excelSheet.proto", ctx.ProtoFilePath())
	assert.Equal(t, "excelSheet.bytes", ctx.ProtoBytesFileName())
	assert.Equal(t, "proto/excelSheet.bytes", ctx.ProtoBytesFilePath())
	assert.Equal(t, "excelSheet.hpp", ctx.ProtoCppFileName())
	assert.Equal(t, "protoCpp/excelSheet.hpp", ctx.ProtoCppFilePath())
	assert.Equal(t, "excelSheet.cs", ctx.ProtoCsFileName())
	assert.Equal(t, "protoCs/excelSheet.cs", ctx.ProtoCsFilePath())
	assert.Equal(t, "excelSheet.go", ctx.ProtoGoFileName())
	assert.Equal(t, "protoGo/excelSheet.go", ctx.ProtoGoFilePath())
	assert.Equal(t, "ExcelSheet", ctx.StructName())
	assert.Equal(t, "Sheeter", ctx.CppNamespace())
	assert.Equal(t, "Sheeter", ctx.CsNamespace())
	assert.Equal(t, "sheeter", ctx.GoPackage())
	assert.Equal(t, "excelSheet.test", ctx.fileName("test"))
	assert.Equal(t, "excel", ctx.excelName())

	ctx = mockContext()
	ctx.Element.Sheet = testdata.SheetName
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	rows := ctx.GetRows(1)
	assert.NotNil(t, rows)
	util.SilentClose(rows)
	rows = ctx.GetRows(0)
	assert.Nil(t, rows)
	rows = ctx.GetRows(10)
	assert.Nil(t, rows)
	ctx.Element.Sheet = "?????"
	rows = ctx.GetRows(1)
	assert.Nil(t, rows)
	util.SilentClose(ctx.Excel)

	ctx = mockContext()
	ctx.Element.Sheet = testdata.SheetName
	ctx.Excel = testdata.GetTestExcel(testdata.Defect1Excel)
	cols := ctx.GetCols(1)
	assert.Equal(t, []string{"name0#pkey", "name1#bool", "name2#int", "name3#text"}, cols)
	cols = ctx.GetCols(2)
	assert.Equal(t, []string{}, cols)
	cols = ctx.GetCols(10)
	assert.Nil(t, cols)
	util.SilentClose(ctx.Excel)

	ctx = mockContext()
	ctx.Element.Sheet = testdata.SheetName
	ctx.Excel = testdata.GetTestExcel(testdata.Defect1Excel)
	assert.True(t, ctx.IsSheetExists())
	ctx.Element.Sheet = "?????"
	assert.False(t, ctx.IsSheetExists())
	util.SilentClose(ctx.Excel)

	ctx = mockContext()
	result, err := ctx.GenerateCode("{{.StructName}}#{{.CppNamespace}}")
	assert.Nil(t, err)
	assert.Equal(t, "ExcelSheet#Sheeter", string(result[:]))
	result, err = ctx.GenerateCode("{{setline .Columns}}{{newline}}{{newline}}{{newline}}{{newline}}")
	assert.Nil(t, err)
	assert.Equal(t, "\n\n", string(result[:]))
	result, err = ctx.GenerateCode("{{{.Unknown}}")
	assert.NotNil(t, err)
}

func TestColumn(t *testing.T) {
	ctx := mockContext()
	assert.Equal(t, "Name0", ctx.Columns[0].ColumnName())
	assert.Equal(t, "Name1", ctx.Columns[1].ColumnName())
	assert.Equal(t, "Name2", ctx.Columns[2].ColumnName())
	assert.Equal(t, "Name3", ctx.Columns[3].ColumnName())
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
		Columns: []*Column{
			{Name: "name0", Note: "note0", Field: &FieldInt{}},
			{Name: "name1", Note: "note1", Field: &FieldInt{}},
			{Name: "name2", Note: "note2", Field: &FieldInt{}},
			{Name: "name3", Note: "note3", Field: &FieldEmpty{}},
		},
	}
}
