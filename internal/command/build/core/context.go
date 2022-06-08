package core

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"Sheeter/internal/util"

	"github.com/schollz/progressbar/v3"
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
	Global   *Global                  // 全域設定
	Element  *Element                 // 項目設定
	Progress *progressbar.ProgressBar // 進度條
	Sheets   Sheets                   // 表格列表
	Columns  []*Column                // 欄位列表
	Pkey     *Column                  // 主要索引欄位
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

// Sheets 表格列表
type Sheets [][]string

// Size 取得表格數量
func (this *Sheets) Size() int {
	count := 0

	for _, itor := range *this {
		count = count + len(itor)
	} // for

	return count
}

// Column 欄位資料
type Column struct {
	Note  string   // 欄位註解
	Name  string   // 欄位名稱
	Field Field    // 欄位類型
	Datas []string // 資料列表
}

// ColumnName 取得欄位名稱
func (this *Column) ColumnName() string {
	return util.FirstUpper(this.Name)
}
