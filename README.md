# Sheeter
用Go做成的excel轉換工具  
用於將按照格式做好的excel轉換為json格式  
轉換時會自動產生c++/c#/go的結構程式碼, 就不必再手寫了  

# 如何執行
配置好yaml格式的設定檔與excel檔案  
執行 sheeter build 設定檔.yaml

# 如何寫設定檔
```
global:
  excelPath: .\                 # excel檔案的路徑
  cppLibraryPath: nlohmann\json # c++使用的json函式庫路徑
  bom: true                     # 輸出的檔案是否含BOM
  lineOfField: 1                # excel表格中欄位行位置, 從1起算
  lineOfNote: 2                 # excel表格中註解行位置, 從1起算
  lineOfData: 3                 # excel表格中資料從哪行開始, 從1起算

elements:
  - excel: excel1.xlsx          # 要轉換的excel檔名
    sheet: Data                 # 要轉換的表格名稱
  - excel: excel2.xlsx
    sheet: Data
  - excel: excel3.xlsx
    sheet: Data
```

# 如何寫excel檔案
![excel_example](docs/excel_example.jpg)

## 欄位行
欄位的格式為`名稱#格式`, 空格之後的欄位不會輸出
目前支援的格式列於下表
| 格式    | 說明                  |
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

## 註解行
單行註解, 若為空格就輸出空註解

## 資料行
依照格式填寫相應的內容即可, 其中`empty, text, textArray`這三種格式允許空格, 其他格式的空格會造成錯誤  
另外空表格(也就是沒有任何資料行)是允許的

## 其他的限制
* 表格必須有欄位行與註解行, 但是可以不需要有資料行
* 欄位行與註解行必須在資料行之前
* 設定檔中必須設定好欄位行, 註解行, 資料行的位置; 設定時要注意行數是從1開始的
* 表格必須有`pkey`欄位
* 表格只能有一個`pkey`欄位
* `pkey`欄位中的內容不能重複
* 欄位名稱不能重複(包括`empty`欄位)

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
