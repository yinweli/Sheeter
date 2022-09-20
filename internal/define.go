package internal

/* 應用程式相關 */

const AppName = "sheeter"   // 程式名稱
const Version = "0.3.8"     // 版本字串
const BarWidth = 40         // 進度條寬度
const SeparateField = "#"   // 欄位字串以'#'符號分割為名稱與欄位
const SeparateElement = "#" // 項目字串以'#'符號分割為檔案名稱與表單名稱

/* 模板相關 */

const PathCode = "template"                       // 輸出路徑: 程式碼模板
const FileCodeJsonCsStruct = "json-cs-struct.txt" // 輸出檔名: json-cs結構模板
const FileCodeJsonCsReader = "json-cs-reader.txt" // 輸出檔名: json-cs讀取器模板
const FileCodeJsonGoStruct = "json-go-struct.txt" // 輸出檔名: json-go結構模板
const FileCodeJsonGoReader = "json-go-reader.txt" // 輸出檔名: json-go讀取器模板
const TokenArray = "[]"                           // 模板字串: 陣列
const TokenBool = "bool"                          // 模板字串: 布林值
const TokenFloatCs = "double"                     // 模板字串: 浮點數(cs)
const TokenFloatGo = "float64"                    // 模板字串: 浮點數(go)
const TokenIntCs = "long"                         // 模板字串: 整數(cs)
const TokenIntGo = "int64"                        // 模板字串: 整數(go)
const TokenString = "string"                      // 模板字串: 字串

/* json相關 */

const PathJson = "json"              // 輸出路徑: json資料
const PathJsonSchema = "json-schema" // 輸出路徑: json架構
const PathJsonCs = "json-cs"         // 輸出路徑: json-cs
const PathJsonGo = "json-go"         // 輸出路徑: json-go
const ExtJson = "json"               // 副檔名: json
const ExtCs = "cs"                   // 副檔名: cs
const ExtGo = "go"                   // 副檔名: go
const Reader = "Reader"              // 讀取器名稱

/* protobuf相關 */

// TODO: protobuf

/* flatbuf相關 */

// TODO: flatbuf
