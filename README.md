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

# 已知問題
浮點數數值在輸出到json格式時, 會因為浮點數精度的問題造成錯誤. 例如表格中原本是0.1, 輸出到json變成0.10000000149011612

# TODO
* 多執行緒版本
* 讓BOM機制有效
* 事前檢查機制(可能要加到Jobs中), 要檢查是否有安裝Go/Protoc
* * 浮點數轉換由strconv.ParseFloat改用strconv.FormatFloat
* 浮點數精度改由設定檔指定或是在欄位上指定
* 加上writeProto格式