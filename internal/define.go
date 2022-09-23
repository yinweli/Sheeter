package internal

/* 應用程式相關 */

const AppName = "sheeter"   // 程式名稱
const Version = "0.4.0"     // 版本字串
const BarWidth = 40         // 進度條寬度
const SeparateField = "#"   // 欄位字串以'#'符號分割為名稱與欄位
const SeparateElement = "#" // 項目字串以'#'符號分割為檔案名稱與表單名稱

/* 類型相關 */

const TokenBoolCs = "bool"        // 類型字串: 布林值(cs)
const TokenBoolGo = "bool"        // 類型字串: 布林值(go)
const TokenBoolProto = "bool"     // 類型字串: 布林值(proto)
const TokenIntCs = "long"         // 類型字串: 整數(cs)
const TokenIntGo = "int64"        // 類型字串: 整數(go)
const TokenIntProto = "int64"     // 類型字串: 整數(proto)
const TokenFloatCs = "double"     // 類型字串: 浮點數(cs)
const TokenFloatGo = "float64"    // 類型字串: 浮點數(go)
const TokenFloatProto = "double"  // 類型字串: 浮點數(proto)
const TokenStringCs = "string"    // 類型字串: 字串(cs)
const TokenStringGo = "string"    // 類型字串: 字串(go)
const TokenStringProto = "string" // 類型字串: 字串(proto)
const TokenArray = "[]"           // 類型字串: 陣列
const TokenOptional = "optional"  // 類型字串: optional(proto)
const TokenRepeated = "repeated"  // 類型字串: repeated(proto)

/* 模板相關 */

const PathTmpl = "template"                         // 輸出路徑: 模板檔案
const FileTmplJsonCsStruct = "json-cs-struct.txt"   // 輸出檔名: json-cs結構模板
const FileTmplJsonCsReader = "json-cs-reader.txt"   // 輸出檔名: json-cs讀取器模板
const FileTmplJsonGoStruct = "json-go-struct.txt"   // 輸出檔名: json-go結構模板
const FileTmplJsonGoReader = "json-go-reader.txt"   // 輸出檔名: json-go讀取器模板
const FileTmplProtoSchema = "proto-schema.txt"      // 輸出檔名: proto架構模板
const FileTmplProtoCsReader = "proto-cs-reader.txt" // 輸出檔名: proto-cs讀取器模板
const FileTmplProtoGoReader = "proto-go-reader.txt" // 輸出檔名: proto-go讀取器模板
const FileTmplProtoCsBat = "proto-cs-bat.txt"       // 輸出檔名: proto-cs-bat模板
const FileTmplProtoGoBat = "proto-go-bat.txt"       // 輸出檔名: proto-go-bat模板
const FileTmplProtoCsSh = "proto-cs-sh.txt"         // 輸出檔名: proto-cs-sh模板
const FileTmplProtoGoSh = "proto-go-sh.txt"         // 輸出檔名: proto-go-sh模板

/* 路徑, 檔名, 副檔名 */

const PathCs = "cs"              // 輸出路徑: cs
const PathGo = "go"              // 輸出路徑: go
const PathData = ".data"         // 輸出路徑: 資料
const PathJson = "json"          // 輸出路徑: json
const PathProto = "proto"        // 輸出路徑: proto
const ExtCs = "cs"               // 副檔名: cs
const ExtGo = "go"               // 副檔名: go
const ExtProtoSchema = "proto"   // 副檔名: proto架構
const ExtJsonData = "json"       // 副檔名: json資料
const ExtProtoData = "pbd"       // 副檔名: proto資料
const FileProtoBat = "build.bat" // 檔名: build.bat
const FileProtoSh = "build.sh"   // 檔名: build.sh
const Reader = "Reader"          // 讀取器名稱
