# Sheeter
用Go做成的表格轉換工具

# 執行方式
sheeter build xxxx.yaml

# 設定檔格式
參考testdata\real.yaml

# Excel格式
參考testdata\real.xlsx

# 支援的欄位格式
| 欄位格式        | 說明                  |
|:------------|:--------------------|
| empty       | 不會輸出的欄位             |
| pkey        | 表格主要索引, 編號可跳號但是不可重複 |
| bool        | 布林值                 |
| boolArray   | 以逗號分隔的布林值陣列         |
| int         | 32位元整數              |
| intArray    | 以逗號分隔的32位元整數陣列      |
| long        | 64位元整數              |
| longArray   | 以逗號分隔的64位元整數陣列      |
| float       | 32位元浮點數             |
| floatArray  | 以逗號分隔的32位元整數陣列      |
| double      | 64位元浮點數             |
| doubleArray | 以逗號分隔的64位元整數陣列      |
| text        | 字串                  |
| textArray   | 以逗號分隔的字串陣列          |

# TODO
* 產生proto message
* 產生proto bytes data
* 產生proto/c++ code
* 產生proto/cs code
* 產生proto/go code
* 新增json/c++驗證子專案
* 新增json/cs驗證子專案
* 新增json/go驗證子專案
* 新增proto/c++驗證子專案
* 新增proto/cs驗證子專案
* 新增proto/go驗證子專案