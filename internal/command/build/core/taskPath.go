package core

import (
	"fmt"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

const extJsonSchema = "json.schema" // 副檔名: json架構
const extJson = "json"              // 副檔名: json
const extCs = "cs"                  // 副檔名: c#
const extGo = "go"                  // 副檔名: go
const extLua = "lua"                // 副檔名: lua
const pathSchema = "schema"         // 輸出路徑: 架構
const pathJson = "json"             // 輸出路徑: json
const pathJsonCs = "jsonCs"         // 輸出路徑: json-c#
const pathJsonGo = "jsonGo"         // 輸出路徑: json-go
const pathLua = "lua"               // 輸出路徑: lua

// excelFilePath 取得excel檔名路徑
func (this *Task) excelFilePath() string {
	return path.Join(this.global.ExcelPath, this.element.Excel)
}

// jsonSchemaFilePath 取得json架構檔名路徑
func (this *Task) jsonSchemaFilePath() string {
	return path.Join(pathSchema, this.fileName(extJsonSchema))
}

// jsonFileName 取得json檔名
func (this *Task) jsonFileName() string {
	return this.fileName(extJson)
}

// jsonFilePath 取得json檔名路徑
func (this *Task) jsonFilePath() string {
	return path.Join(pathJson, this.jsonFileName())
}

// jsonCsFilePath 取得json-c#檔名路徑
func (this *Task) jsonCsFilePath() string {
	return path.Join(pathJsonCs, this.fileName(extCs))
}

// jsonGoFilePath 取得json-go檔名路徑
func (this *Task) jsonGoFilePath() string {
	return path.Join(pathJsonGo, this.fileName(extGo))
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
