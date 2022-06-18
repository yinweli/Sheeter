package core

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/yinweli/Sheeter/internal/util"
)

const pathSchema = "schema"   // 輸出路徑: 架構
const pathJson = "json"       // 輸出路徑: json
const pathJsonCs = "jsonCs"   // 輸出路徑: json/c#
const pathJsonGo = "jsonGo"   // 輸出路徑: json/go
const pathProto = "proto"     // 輸出路徑: proto
const pathProtoCs = "protoCs" // 輸出路徑: proto/c#
const pathProtoGo = "protoGo" // 輸出路徑: proto/go
const pathLua = "lua"         // 輸出路徑: lua
const extSchema = "schema"    // 副檔名: 架構
const extJson = "json"        // 副檔名: json
const extProto = "proto"      // 副檔名: proto
const extBytes = "bytes"      // 副檔名: bytes
const extCs = "cs"            // 副檔名: c#
const extGo = "go"            // 副檔名: go
const extLua = "lua"          // 副檔名: lua

// originalName 取得原始名稱
func (this *Task) originalName() string {
	return fmt.Sprintf("%s(%s)", this.element.Excel, this.element.Sheet)
}

// excelFilePath 取得excel檔名路徑
func (this *Task) excelFilePath() string {
	return path.Join(this.global.ExcelPath, this.element.Excel)
}

// schemaFilePath 取得架構檔名路徑
func (this *Task) schemaFilePath() string {
	return path.Join(pathSchema, this.fileName(extSchema))
}

// jsonFileName 取得json檔名
func (this *Task) jsonFileName() string {
	return this.fileName(extJson)
}

// jsonFilePath 取得json檔名路徑
func (this *Task) jsonFilePath() string {
	return path.Join(pathJson, this.jsonFileName())
}

// jsonCsFilePath 取得json/c#檔名路徑
func (this *Task) jsonCsFilePath() string {
	return path.Join(pathJsonCs, this.fileName(extCs))
}

// jsonGoFilePath 取得json/go檔名路徑
func (this *Task) jsonGoFilePath() string {
	return path.Join(pathJsonGo, this.fileName(extGo))
}

// protoFilePath 取得proto檔名路徑
func (this *Task) protoFilePath() string {
	return path.Join(pathProto, this.fileName(extProto))
}

// protoBytesFilePath 取得proto資料檔名路徑
func (this *Task) protoBytesFilePath() string {
	return path.Join(pathProto, this.fileName(extBytes))
}

// protoCsFilePath 取得proto/c#檔名路徑
func (this *Task) protoCsFilePath() string {
	return path.Join(pathProtoCs, this.fileName(extCs))
}

// protoGoFilePath 取得proto/go檔名路徑
func (this *Task) protoGoFilePath() string {
	return path.Join(pathProtoGo, this.fileName(extGo))
}

// luaFilePath 取得lua檔名路徑
func (this *Task) luaFilePath() string {
	return path.Join(pathLua, this.fileName(extLua))
}

// fileName 取得檔案名稱
func (this *Task) fileName(ext string) string {
	excelName := util.FirstLower(this.excelName())
	sheetName := util.FirstUpper(this.element.Sheet)

	return fmt.Sprintf("%s%s.%s", excelName, sheetName, ext)
}

// excelName 取得沒有副檔名的excel名稱
func (this *Task) excelName() string {
	return strings.TrimSuffix(this.element.Excel, filepath.Ext(this.element.Excel))
}

// structName 取得結構名稱
func (this *Task) structName() string {
	excelName := util.FirstUpper(this.excelName())
	sheetName := util.FirstUpper(this.element.Sheet)

	return excelName + sheetName
}

// sheetExists 表格是否存在
func (this *Task) sheetExists() bool {
	return this.excel.GetSheetIndex(this.element.Sheet) != -1
}

// getRows 取得表格行資料, line從1起算
func (this *Task) getRows(line int) *excelize.Rows {
	if line <= 0 {
		return nil
	} // if

	rows, err := this.excel.Rows(this.element.Sheet)

	if err != nil {
		return nil
	} // if

	for l := 0; l < line; l++ {
		if rows.Next() == false { // 注意! 最少要一次才能定位到第1行; 所以若line=0, 則取不到資料
			_ = rows.Close()
			return nil
		} // if
	} // for

	return rows
}

// getRowContent 取得表格行內容, line從1起算
func (this *Task) getRowContent(line int) []string {
	rows := this.getRows(line)

	if rows == nil {
		return nil
	} // if

	defer func() { _ = rows.Close() }()
	cols, err := rows.Columns()

	if err != nil {
		return nil
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols
}
