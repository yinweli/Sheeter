package core

import (
	"fmt"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

const extProto = "proto"      // 副檔名: proto
const extJson = "json"        // 副檔名: json
const extBytes = "bytes"      // 副檔名: bytes
const extLua = "lua"          // 副檔名: lua
const extCs = "cs"            // 副檔名: c#
const extGo = "go"            // 副檔名: go
const pathSchema = "schema"   // 輸出路徑: 架構
const pathLua = "lua"         // 輸出路徑: lua
const pathJson = "json"       // 輸出路徑: json
const pathBytes = "bytes"     // 輸出路徑: bytes
const pathJsonCs = "jsonCs"   // 輸出路徑: json/c#
const pathJsonGo = "jsonGo"   // 輸出路徑: json/go
const pathProtoCs = "protoCs" // 輸出路徑: proto/c#
const pathProtoGo = "protoGo" // 輸出路徑: proto/go

// excelFilePath 取得excel檔名路徑
func (this *Task) excelFilePath() string {
	return path.Join(this.global.ExcelPath, this.element.Excel)
}

// schemaProtoFilePath 取得proto架構檔名路徑
func (this *Task) schemaProtoFilePath() string {
	return path.Join(pathSchema, this.fileName(extProto))
}

// jsonFileName 取得json檔名
func (this *Task) jsonFileName() string {
	return this.fileName(extJson)
}

// jsonFilePath 取得json檔名路徑
func (this *Task) jsonFilePath() string {
	return path.Join(pathJson, this.jsonFileName())
}

// bytesFileName 取得bytes檔名
func (this *Task) bytesFileName() string {
	return this.fileName(extBytes)
}

// bytesFilePath 取得bytes檔名路徑
func (this *Task) bytesFilePath() string {
	return path.Join(pathBytes, this.bytesFileName())
}

// luaFilePath 取得lua檔名路徑
func (this *Task) luaFilePath() string {
	return path.Join(pathLua, this.fileName(extLua))
}

// jsonCsFilePath 取得json/c#檔名路徑
func (this *Task) jsonCsFilePath() string {
	return path.Join(pathJsonCs, this.fileName(extCs))
}

// jsonGoFilePath 取得json/go檔名路徑
func (this *Task) jsonGoFilePath() string {
	return path.Join(pathJsonGo, this.fileName(extGo))
}

// protoCsFilePath 取得proto/c#檔名路徑
func (this *Task) protoCsFilePath() string {
	return path.Join(pathProtoCs, this.fileName(extCs))
}

// protoGoFilePath 取得proto/go檔名路徑
func (this *Task) protoGoFilePath() string {
	return path.Join(pathProtoGo, this.fileName(extGo))
}

// fileName 取得檔案名稱
func (this *Task) fileName(ext string) string {
	excelName := util.FirstLower(this.excelName())
	sheetName := util.FirstUpper(this.element.Sheet)

	return fmt.Sprintf("%s%s.%s", excelName, sheetName, ext)
}
