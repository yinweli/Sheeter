package sheeter

/* 應用程式相關 */

const AppName = "sheeter" // 程式名稱
const Version = "2.3.13"  // 版本字串, 遵循'大版本.小版本.修正版本'的規則
const BarWidth = 40       // 進度條寬度
const MaxExcel = 999999   // 最大開啟excel數量
const MaxSheet = 999999   // 最大開啟sheet數量
const IndexOutput = 0     // 輸出欄位置編號
const IndexAutoKey = 1    // 自動索引位置編號

/* 路徑名 */

const CsPath = "codeCs" // 輸出路徑: cs
const GoPath = "codeGo" // 輸出路徑: go
const JsonPath = "json" // 輸出路徑: json

/* 副檔名 */

const ExcelExt = ".xlsx" // 副檔名: excel
const CsExt = ".cs"      // 副檔名: cs
const GoExt = ".go"      // 副檔名: go
const JsonExt = ".json"  // 副檔名: json

/* 程式名稱 */

const Reader = "Reader"   // 讀取器名稱
const Sheeter = AppName   // 表格器名稱
const Namespace = AppName // 命名空間名稱

/* 其他名稱 */

const TokenIgnore = "ignore" // 忽略符號, 當輸出或是標籤為此名稱時不輸出
const TokenArray = ","       // 陣列分割符號, 陣列以','符號分割元素
const TokenName = "$"        // 名稱分割符號, 以'$'來分隔名稱與後續項目
const TokenTerm = "&"        // 項目分割符號, 以'&'來分隔項目
const TokenExcel = "#"       // 表格分割符號, 以'#'來分隔項目中的excel與sheet名稱
const JsonPrefix = ""        // json前綴字串
const JsonIdent = "    "     // json縮排字串

/* 類型名稱 */

const TypePkeyCs = "System.Int32"  // 類型字串: pkey(cs)
const TypePkeyGo = "int32"         // 類型字串: pkey(go)
const TypeLkeyCs = "System.Int64"  // 類型字串: lkey(cs)
const TypeLkeyGo = "int64"         // 類型字串: lkey(go)
const TypeSkeyCs = "System.String" // 類型字串: skey(cs)
const TypeSkeyGo = "string"        // 類型字串: skey(go)
const TypeBoolCs = "bool"          // 類型字串: 布林值(cs)
const TypeBoolGo = "bool"          // 類型字串: 布林值(go)
const TypeIntCs = "int"            // 類型字串: 32位元整數(cs)
const TypeIntGo = "int32"          // 類型字串: 32位元整數(go)
const TypeLongCs = "long"          // 類型字串: 64位元整數(cs)
const TypeLongGo = "int64"         // 類型字串: 64位元整數(go)
const TypeFloatCs = "float"        // 類型字串: 32位元浮點數(cs)
const TypeFloatGo = "float32"      // 類型字串: 32位元浮點數(go)
const TypeDoubleCs = "double"      // 類型字串: 64位元浮點數(cs)
const TypeDoubleGo = "float64"     // 類型字串: 64位元浮點數(go)
const TypeStringCs = "string"      // 類型字串: 字串(cs)
const TypeStringGo = "string"      // 類型字串: 字串(go)
const TypeArray = "[]"             // 類型字串: 陣列
