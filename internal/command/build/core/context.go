package core

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"Sheeter/internal/util"

	"github.com/xuri/excelize/v2"
)

const PathJson = "json"         // 輸出路徑: json
const PathJsonCpp = "jsonCpp"   // 輸出路徑: json/c++
const PathJsonCs = "jsonCs"     // 輸出路徑: json/c#
const PathJsonGo = "jsonGo"     // 輸出路徑: json/go
const PathProto = "proto"       // 輸出路徑: proto
const PathProtoCpp = "protoCpp" // 輸出路徑: proto/c++
const PathProtoCs = "protoCs"   // 輸出路徑: proto/c#
const PathProtoGo = "protoGo"   // 輸出路徑: proto/go
const ExtJson = "json"          // 副檔名: json
const ExtProto = "proto"        // 副檔名: proto
const ExtBytes = "bytes"        // 副檔名: bytes
const ExtCpp = "hpp"            // 副檔名: c++
const ExtCs = "cs"              // 副檔名: c#
const ExtGo = "go"              // 副檔名: go
const CppNamespace = "Sheeter"  // c++命名空間名稱
const CsNamespace = "Sheeter"   // c#命名空間名稱
const GoPackage = "sheeter"     // go包名

// Context 工作資料
type Context struct {
	Global  *Global        // 全域設定
	Element *Element       // 項目設定
	Excel   *excelize.File // excel物件
	Columns []*Column      // 欄位列表
}

// LogName 取得紀錄名稱
func (this *Context) LogName() string {
	return fmt.Sprintf("%s(%s)", this.Element.Excel, this.Element.Sheet)
}

// ExcelFilePath 取得excel檔名路徑
func (this *Context) ExcelFilePath() string {
	return path.Join(this.Global.ExcelPath, this.Element.Excel)
}

// JsonFileName 取得json檔名
func (this *Context) JsonFileName() string {
	return this.fileName(ExtJson)
}

// JsonFilePath 取得json檔名路徑
func (this *Context) JsonFilePath() string {
	return path.Join(PathJson, this.JsonFileName())
}

// JsonCppFileName 取得json/c++檔名
func (this *Context) JsonCppFileName() string {
	return this.fileName(ExtCpp)
}

// JsonCppFilePath 取得json/c++檔名路徑
func (this *Context) JsonCppFilePath() string {
	return path.Join(PathJsonCpp, this.JsonCppFileName())
}

// JsonCsFileName 取得json/c#檔名
func (this *Context) JsonCsFileName() string {
	return this.fileName(ExtCs)
}

// JsonCsFilePath 取得json/c#檔名路徑
func (this *Context) JsonCsFilePath() string {
	return path.Join(PathJsonCs, this.JsonCsFileName())
}

// JsonGoFileName 取得json/go檔名
func (this *Context) JsonGoFileName() string {
	return this.fileName(ExtGo)
}

// JsonGoFilePath 取得json/go檔名路徑
func (this *Context) JsonGoFilePath() string {
	return path.Join(PathJsonGo, this.JsonGoFileName())
}

// ProtoFileName 取得proto檔名
func (this *Context) ProtoFileName() string {
	return this.fileName(ExtProto)
}

// ProtoFilePath 取得proto檔名路徑
func (this *Context) ProtoFilePath() string {
	return path.Join(PathProto, this.ProtoFileName())
}

// ProtoBytesFileName 取得proto資料檔名
func (this *Context) ProtoBytesFileName() string {
	return this.fileName(ExtBytes)
}

// ProtoBytesFilePath 取得proto資料檔名路徑
func (this *Context) ProtoBytesFilePath() string {
	return path.Join(PathProto, this.ProtoBytesFileName())
}

// ProtoCppFileName 取得proto/c++檔名 // TODO: 搞不好不需要QQ
func (this *Context) ProtoCppFileName() string {
	return this.fileName(ExtCpp)
}

// ProtoCppFilePath 取得proto/c++檔名路徑 // TODO: 搞不好不需要QQ
func (this *Context) ProtoCppFilePath() string {
	return path.Join(PathProtoCpp, this.ProtoCppFileName())
}

// ProtoCsFileName 取得proto/c#檔名 // TODO: 搞不好不需要QQ
func (this *Context) ProtoCsFileName() string {
	return this.fileName(ExtCs)
}

// ProtoCsFilePath 取得proto/c#檔名路徑 // TODO: 搞不好不需要QQ
func (this *Context) ProtoCsFilePath() string {
	return path.Join(PathProtoCs, this.ProtoCsFileName())
}

// ProtoGoFileName 取得proto/go檔名 // TODO: 搞不好不需要QQ
func (this *Context) ProtoGoFileName() string {
	return this.fileName(ExtGo)
}

// ProtoGoFilePath 取得proto/go檔名路徑 // TODO: 搞不好不需要QQ
func (this *Context) ProtoGoFilePath() string {
	return path.Join(PathProtoGo, this.ProtoGoFileName())
}

// StructName 取得結構名稱
func (this *Context) StructName() string {
	excelName := util.FirstUpper(this.excelName())
	sheetName := util.FirstUpper(this.Element.Sheet)

	return excelName + sheetName
}

// CppNamespace 取得c++命名空間名稱
func (this *Context) CppNamespace() string {
	return CppNamespace
}

// CsNamespace 取得c#命名空間名稱
func (this *Context) CsNamespace() string {
	return CsNamespace
}

// GoPackage 取得go包名
func (this *Context) GoPackage() string {
	return GoPackage
}

// GetRows 取得表格行資料, line從1起算
func (this *Context) GetRows(line int) *excelize.Rows {
	if line <= 0 {
		return nil
	} // if

	rows, err := this.Excel.Rows(this.Element.Sheet)

	if err != nil {
		return nil
	} // if

	for l := 0; l < line; l++ {
		if rows.Next() == false { // 注意! 最少要一次才能定位到第1行; 所以若line=0, 則取不到資料
			util.SilentClose(rows)
			return nil
		} // if
	} // for

	return rows
}

// GetCols 取得表格行內容, line從1起算
func (this *Context) GetCols(line int) []string {
	rows := this.GetRows(line)

	if rows == nil {
		return nil
	} // if

	defer util.SilentClose(rows)
	cols, err := rows.Columns()

	if err != nil {
		return nil
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols
}

// IsSheetExists 表格是否存在
func (this *Context) IsSheetExists() bool {
	return this.Excel.GetSheetIndex(this.Element.Sheet) != -1
}

// GenerateCode 產生程式碼
func (this *Context) GenerateCode(code string) (results []byte, err error) {
	maxline := 0
	curline := 0
	temp, err := template.New("generateCode").Funcs(template.FuncMap{
		"setline": func(columns []*Column) string {
			maxline = 0

			for _, itor := range columns {
				if itor.Field.IsShow() {
					maxline++
				} // if
			} // for

			maxline = maxline - 1 // 減一是為了避免多換一次新行
			curline = 0
			return ""
		},
		"newline": func() string {
			result := ""

			if maxline > curline {
				result = "\n"
			} // if

			curline++
			return result
		},
	}).Parse(code)

	if err != nil {
		return nil, err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, this)

	if err != nil {
		return nil, err
	} // if

	return buffer.Bytes(), nil
}

// fileName 取得檔案名稱
func (this *Context) fileName(ext string) string {
	excelName := util.FirstLower(this.excelName())
	sheetName := util.FirstUpper(this.Element.Sheet)

	return fmt.Sprintf("%s%s.%s", excelName, sheetName, ext)
}

// excelName 取得沒有副檔名的excel名稱
func (this *Context) excelName() string {
	return strings.TrimSuffix(this.Element.Excel, filepath.Ext(this.Element.Excel))
}

// Column 欄位資料
type Column struct {
	Name  string // 欄位名稱
	Note  string // 欄位註解
	Field Field  // 欄位類型
}

// ColumnName 取得欄位名稱
func (this *Column) ColumnName() string {
	return util.FirstUpper(this.Name)
}
