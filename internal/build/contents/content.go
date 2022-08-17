package contents

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/util"

	"github.com/vbauerster/mpb/v7"
)

const pathSchema = "schema"  // 輸出路徑: json架構
const pathJson = "json"      // 輸出路徑: json
const pathJsonCs = "json-cs" // 輸出路徑: json-c#
const pathJsonGo = "json-go" // 輸出路徑: json-go
const midReader = "reader"   // 中間名: 讀取器
const extSchema = "schema"   // 副檔名: json架構
const extJson = "json"       // 副檔名: json
const extCs = "cs"           // 副檔名: c#
const extGo = "go"           // 副檔名: go

// Content 內容資料
type Content struct {
	Path        string        // 來源excel路徑
	Bom         bool          // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int           // 欄位行號(1為起始行)
	LineOfLayer int           // 階層行號(1為起始行)
	LineOfNote  int           // 註解行號(1為起始行)
	LineOfData  int           // 資料起始行號(1為起始行)
	Excel       string        // excel檔案名稱
	Sheet       string        // excel表單名稱
	Progress    *mpb.Progress // 進度條產生器
}

// Check 檢查工作
func (this *Content) Check() error {
	if this.LineOfField <= 0 {
		return fmt.Errorf("lineOfField <= 0")
	} // if

	if this.LineOfLayer <= 0 {
		return fmt.Errorf("lineOfLayer <= 0")
	} // if

	if this.LineOfNote <= 0 {
		return fmt.Errorf("lineOfNote <= 0")
	} // if

	if this.LineOfData <= 0 {
		return fmt.Errorf("lineOfData <= 0")
	} // if

	if this.LineOfField >= this.LineOfData {
		return fmt.Errorf("lineOfField(%d) >= lineOfData(%d)", this.LineOfField, this.LineOfData)
	} // if

	if this.LineOfLayer >= this.LineOfData {
		return fmt.Errorf("lineOfLayer(%d) >= lineOfData(%d)", this.LineOfLayer, this.LineOfData)
	} // if

	if this.LineOfNote >= this.LineOfData {
		return fmt.Errorf("lineOfNote(%d) >= lineOfData(%d)", this.LineOfNote, this.LineOfData)
	} // if

	if this.Excel == "" {
		return fmt.Errorf("excel empty")
	} // if

	if this.Sheet == "" {
		return fmt.Errorf("sheet empty")
	} // if

	if this.Progress == nil {
		return fmt.Errorf("progress nil")
	} // if

	return nil
}

// ExcelFilePath 取得excel檔名路徑
func (this *Content) ExcelFilePath() string {
	return path.Join(this.Path, this.Excel)
}

// SchemaFilePath 取得json架構檔名路徑
func (this *Content) SchemaFilePath() string {
	return path.Join(pathSchema, this.fileName(extSchema))
}

// JsonFileName 取得json檔名
func (this *Content) JsonFileName() string {
	return this.fileName(extJson)
}

// JsonFilePath 取得json檔名路徑
func (this *Content) JsonFilePath() string {
	return path.Join(pathJson, this.JsonFileName())
}

// JsonCsFilePath 取得json-c#檔名路徑
func (this *Content) JsonCsFilePath() string {
	return path.Join(pathJsonCs, this.fileName(extCs))
}

// JsonCsReaderFilePath 取得json-c#讀取器檔名路徑
func (this *Content) JsonCsReaderFilePath() string {
	return path.Join(pathJsonCs, this.fileName(midReader, extCs))
}

// JsonGoFilePath 取得json-go檔名路徑
func (this *Content) JsonGoFilePath() string {
	return path.Join(pathJsonGo, this.fileName(extGo))
}

// JsonGoReaderFilePath 取得json-go讀取器檔名路徑
func (this *Content) JsonGoReaderFilePath() string {
	return path.Join(pathJsonGo, this.fileName(midReader, extGo))
}

// Namespace 取得命名空間名稱
func (this *Content) Namespace() string {
	return internal.Title
}

// TargetName 取得目標名稱
func (this *Content) TargetName() string {
	return fmt.Sprintf("%s(%s)", this.Excel, this.Sheet)
}

// StructName 取得結構名稱
func (this *Content) StructName() string {
	excelName := util.FirstUpper(this.ExcelName())
	sheetName := util.FirstUpper(this.Sheet)

	return excelName + sheetName
}

// ReaderName 取得讀取器名稱
func (this *Content) ReaderName() string {
	return this.StructName() + midReader
}

// ExcelName 取得沒有副檔名的excel名稱
func (this *Content) ExcelName() string {
	return strings.TrimSuffix(this.Excel, filepath.Ext(this.Excel))
}

// fileName 取得檔案名稱
func (this *Content) fileName(ext ...string) string {
	excelName := util.FirstLower(this.ExcelName())
	sheetName := util.FirstUpper(this.Sheet)

	fileNames := []string{}
	fileNames = append(fileNames, excelName+sheetName)
	fileNames = append(fileNames, ext...)

	return strings.Join(fileNames, ".")
}
