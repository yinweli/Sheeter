package sheeter

/* 應用程式相關 */

const AppName = "sheeter" // 程式名稱
const Version = "1.11.0"  // 版本字串, 遵循'大版本.小版本.修正版本'的規則
const BarWidth = 40       // 進度條寬度
const MaxExcel = 999999   // 最大開啟excel數量

/* 類型相關 */

type PkeyType = int64 // pkey類型, 修改時要記得跟著修改TokenPkey...系列的類型字串

const TokenPkeyCs = "System.Int64" // 類型字串: pkey(cs)
const TokenPkeyGo = "int64"        // 類型字串: pkey(go)
const TokenPkeyProto = "int64"     // 類型字串: pkey(proto)
const TokenBoolCs = "bool"         // 類型字串: 布林值(cs)
const TokenBoolGo = "bool"         // 類型字串: 布林值(go)
const TokenBoolProto = "bool"      // 類型字串: 布林值(proto)
const TokenIntCs = "long"          // 類型字串: 整數(cs)
const TokenIntGo = "int64"         // 類型字串: 整數(go)
const TokenIntProto = "int64"      // 類型字串: 整數(proto)
const TokenFloatCs = "double"      // 類型字串: 浮點數(cs)
const TokenFloatGo = "float64"     // 類型字串: 浮點數(go)
const TokenFloatProto = "double"   // 類型字串: 浮點數(proto)
const TokenStringCs = "string"     // 類型字串: 字串(cs)
const TokenStringGo = "string"     // 類型字串: 字串(go)
const TokenStringProto = "string"  // 類型字串: 字串(proto)
const TokenArray = "[]"            // 類型字串: 陣列
const TokenOptional = "optional"   // 類型字串: optional(proto)
const TokenRepeated = "repeated"   // 類型字串: repeated(proto)

/* 模板相關 */

const TmplPath = "template"                         // 輸出路徑: 模板檔案
const TmplJsonStructCsFile = "json-struct-cs.txt"   // 輸出檔名: json結構cs模板
const TmplJsonReaderCsFile = "json-reader-cs.txt"   // 輸出檔名: json讀取器cs模板
const TmplJsonDepotCsFile = "json-depot-cs.txt"     // 輸出檔名: json倉庫cs模板
const TmplJsonStructGoFile = "json-struct-go.txt"   // 輸出檔名: json結構go模板
const TmplJsonReaderGoFile = "json-reader-go.txt"   // 輸出檔名: json讀取器go模板
const TmplJsonDepotGoFile = "json-depot-go.txt"     // 輸出檔名: json倉庫go模板
const TmplProtoSchemaFile = "proto-schema.txt"      // 輸出檔名: proto架構模板
const TmplProtoReaderCsFile = "proto-reader-cs.txt" // 輸出檔名: proto讀取器cs模板
const TmplProtoDepotCsFile = "proto-depot-cs.txt"   // 輸出檔名: proto倉庫cs模板
const TmplProtoReaderGoFile = "proto-reader-go.txt" // 輸出檔名: proto讀取器go模板
const TmplProtoDepotGoFile = "proto-depot-go.txt"   // 輸出檔名: proto倉庫go模板
const TmplEnumSchemaFile = "enum-schema.txt"        // 輸出檔名: enum架構模板

/* 路徑名 */

const SchemaPath = "schema" // 輸出路徑: 架構
const DataPath = "data"     // 輸出路徑: 資料
const CsPath = "codeCs"     // 輸出路徑: cs
const GoPath = "codeGo"     // 輸出路徑: go
const JsonPath = "json"     // 輸出路徑: json
const ProtoPath = "proto"   // 輸出路徑: proto
const EnumPath = "enum"     // 輸出路徑: enum

/* 副檔名 */

const ExcelExt = ".xlsx"        // 副檔名: excel
const CsExt = ".cs"             // 副檔名: cs
const GoExt = ".go"             // 副檔名: go
const JsonDataExt = ".json"     // 副檔名: json資料
const ProtoSchemaExt = ".proto" // 副檔名: proto架構
const ProtoDataExt = ".bytes"   // 副檔名: proto資料
const EnumSchemaExt = ".proto"  // 副檔名: enum架構

/* 程式名稱 */

const Reader = "Reader"                  // 讀取器名稱
const Storer = "Storer"                  // 儲存器名稱
const StorerDatas = "Datas"              // 儲存器資料名稱
const Depot = "depot"                    // 倉庫名稱
const JsonNamespace = AppName + "Json"   // 命名空間名稱: json
const ProtoNamespace = AppName + "Proto" // 命名空間名稱: proto
const EnumNamespace = AppName + "Enum"   // 命名空間名稱: enum

/* 其他名稱 */

const SignData = "@"      // 資料表單開頭字元
const SignEnum = "$"      // 列舉表單開頭字元
const SeparateSheet = "#" // 表單字串以'#'符號分割為檔案名稱與表單名稱
const SeparateArray = "," // 陣列字串以','符號分割
const JsonPrefix = ""     // json前綴字串
const JsonIdent = "    "  // json縮排字串

// Keywords 關鍵字列表
var Keywords = []string{
	Depot,
	"loader",
	"reader",
	"readers",
}
