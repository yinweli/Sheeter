package tasks

import (
	"path"
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

const pathSchema = "schema"         // 輸出路徑: 架構
const pathJson = "json"             // 輸出路徑: json
const pathJsonCs = "jsonCs"         // 輸出路徑: json-c#
const pathJsonGo = "jsonGo"         // 輸出路徑: json-go
const midReader = "reader"          // 中間名: 讀取器
const extJsonSchema = "json.schema" // 副檔名: json架構
const extJson = "json"              // 副檔名: json
const extCs = "cs"                  // 副檔名: c#
const extGo = "go"                  // 副檔名: go

// excelFilePath 取得excel檔名路徑
func (this *Task) excelFilePath() string {
	return path.Join(this.Path, this.Excel)
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

// jsonCsReaderFilePath 取得json-c#讀取器檔名路徑
func (this *Task) jsonCsReaderFilePath() string {
	return path.Join(pathJsonCs, this.fileName(midReader, extCs))
}

// jsonGoFilePath 取得json-go檔名路徑
func (this *Task) jsonGoFilePath() string {
	return path.Join(pathJsonGo, this.fileName(extGo))
}

// jsonGoReaderFilePath 取得json-go讀取器檔名路徑
func (this *Task) jsonGoReaderFilePath() string {
	return path.Join(pathJsonGo, this.fileName(midReader, extGo))
}

// fileName 取得檔案名稱
func (this *Task) fileName(ext ...string) string {
	excelName := util.FirstLower(this.excelName())
	sheetName := util.FirstUpper(this.Sheet)

	fileNames := []string{}
	fileNames = append(fileNames, excelName+sheetName)
	fileNames = append(fileNames, ext...)

	return strings.Join(fileNames, ".")
}
