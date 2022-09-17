![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以go做成的excel轉換工具, 前身是[sheet]  
將以指定格式做好的excel轉換為json, cs程式碼, go程式碼

# 目錄說明
| 目錄                | 說明           |
|:--------------------|:---------------|
| doc                 | 說明文件       |
| cmd/sheeter         | sheeter        |
| cmd/sheeter/build   | 建置表格命令   |
| cmd/sheeter/code    | 建置模板命令   |
| cmd/sheeter/version | 顯示版本命令   |
| cmd/verifycs        | cs程式碼驗證器 |
| cmd/verifygo        | go程式碼驗證器 |
| internal/builds     | 表格轉換       |
| internal/codes      | 模板組件       |
| internal/fields     | 欄位組件       |
| internal/layers     | 階層組件       |
| internal/layouts    | 布局組件       |
| internal/names      | 命名組件       |
| internal/utils      | 協助組件       |
| testdata            | 測試資料       |

# 如何安裝
* 安裝[go]
* 安裝[sheeter], 在終端執行以下命令
  ```shell
  go install github.com/yinweli/Sheeter/cmd/sheeter@latest
  ```

# 命令說明
* build: 建置json檔案與結構檔案  
  配置好yaml格式的設定檔與excel檔案, 然後在終端執行  
  ```shell
  sheeter build --config 設定檔.yaml
  ```
* code: 產生模板檔案  
  這會產生建置時使用的模板檔案, 你可以通過修改模板來改變產生出來的程式碼  
  執行建置命令時也會產生模板檔案  
  ```shell
  sheeter code
  ```
* version: 顯示版本資訊  
  ```sheel
  sheeter version
  ```  

# 如何寫設定檔
```yaml
global:
  lineOfField: 1       # 欄位行號(1為起始行)
  lineOfLayer: 2       # 階層行號(1為起始行)
  lineOfNote:  3       # 註解行號(1為起始行)
  lineOfData:  4       # 資料行號(1為起始行)

elements:
  - excel: excel1.xlsx # excel檔案名稱
    sheet: Data        # excel表單名稱
  - excel: excel2.xlsx
    sheet: Data
  - excel: excel3.xlsx
    sheet: Data
```

# 如何寫excel檔案
![excel]

## 欄位行
欄位的格式為`名稱#格式`, 空格之後的欄位不會輸出  

| 格式        | 說明                                 |
|:------------|:-------------------------------------|
| empty       | 不會輸出的欄位                       |
| pkey        | 表格主要索引, 編號可跳號但是不可重複 |
| bool        | 布林值                               |
| boolArray   | 以逗號分隔的布林值陣列               |
| int         | 64位元整數                           |
| intArray    | 以逗號分隔的64位元整數陣列           |
| float       | 64位元浮點數                         |
| floatArray  | 以逗號分隔的64位元整數陣列           |
| text        | 字串                                 |
| textArray   | 以逗號分隔的字串陣列                 |

## 階層行
欄位結構布局, 格式有`{名稱`, `{[]名稱`, `/`, `}`, 之間以空格分隔  

| 格式        | 說明                                 |
|:------------|:-------------------------------------|
| {結構名稱   | 結構的開始                           |
| {[]陣列名稱 | 陣列的開始                           |
| /           | 分隔陣列                             |
| }           | 結構/陣列結束, 可以連續結束, 如`}}`  |

## 註解行
單行註解, 若為空格就輸出空註解  

## 資料行
依照格式填寫相應的內容即可, 其中`empty`, `text`, `textArray`這三種格式允許空格, 其他格式的空格會造成錯誤  
空表格(也就是沒有任何資料行)是允許的  
轉換時, 只會轉換到第一個空行為止  

## 轉出檔案路徑與檔案名稱
如果excel檔案名稱為`example.xlsx`, 表格名稱為`Data`  
* json資料檔案: json\exampleData.json
* json的cs程式碼: json-cs\exampleData.cs
* json的cs讀取器: json-cs\exampleDataReader.cs
* json的go程式碼: json-go\exampleData.go
* json的go讀取器: json-go\exampleDataReader.go
* json架構檔案: json-schema\exampleData.json
* 模板檔案: template\ ...
* 命名空間: sheeter
* 結構名稱: ExampleData
* 讀取器名稱: ExampleDataReader

## 其他的限制
* 表格必須有欄位行, 階層行, 註解行, 但是可以不需要有資料行
* 欄位行, 階層行, 註解行必須在資料行之前
* 設定檔中必須設定好欄位行, 階層行, 註解行, 資料行的位置
* 設定檔中行數是從1開始的
* 表格必須有`pkey`欄位
* 表格只能有一個`pkey`欄位
* `pkey`欄位中的內容不能重複
* 欄位名稱不能重複(包括`empty`欄位)
* cs程式碼使用`Newtonsoft.Json`來轉換json

## 關於模板檔案
sheeter轉換時會把使用的程式碼模板輸出到template目錄下  
使用者可以改變模板內容, 來產生自訂的程式碼  
當sheeter版本更新時, 需要在終端執行以下命令來重置模板  
```shell
sheeter code -c
```
模板檔案使用golang的[template]語法, 同時可以參考以下變數來做結構名稱或是欄位名稱等的替換  

| 名稱                     | 說明                                     |
|:-------------------------|:-----------------------------------------|
| $.Named.AppName          | 程式名稱                                 |
| $.Named.Namespace        | 命名空間名稱                             |
| $.Named.StructName       | 結構名稱                                 |
| $.Named.ReaderName       | 讀取器名稱                               |
| $.Named.FileJson         | json檔名路徑                             |
| $.Named.FileJsonCode     | json檔名路徑(程式碼可用)                 |
| $.Named.FileJsonSchema   | json架構檔名路徑                         |
| $.Named.FileJsonCsCode   | json-cs程式碼檔名路徑                    |
| $.Named.FileJsonCsReader | json-cs讀取器檔名路徑                    |
| $.Named.FileJsonGoCode   | json-go程式碼檔名路徑                    |
| $.Named.FileJsonGoReader | json-go讀取器檔名路徑                    |
| $.Field                  | 欄位列表                                 |
| $.FieldName              | 取得欄位名稱(需要輸入欄位資料作為參數)   |
| $.FieldNote              | 取得欄位註解(需要輸入欄位資料作為參數)   |
| $.FieldTypeCs            | 取得cs欄位類型(需要輸入欄位資料作為參數) |
| $.FieldTypeGo            | 取得go欄位類型(需要輸入欄位資料作為參數) |

# 轉換範例
[example]

# TODO
* 產生protobuffer message
* 產生protobuffer bytes data
* 產生protobuffer/cs code
* 產生protobuffer/go code
* 產生flatbuffer message
* 產生flatbuffer bytes data
* 產生flatbuffer/cs code
* 產生flatbuffer/go code
* quicktype >> c++  
  --src verifyData.json --src-lang json --top-level verifyDatas  
  --out verifyData.hpp --lang c++  
  --namespace sheeter  
  --code-format with-struct  
  --const-style west-const  
  --type-style pascal-case  
  --member-style camel-case  
  --enumerator-style camel-case  
  --source-style multi-source  
  --include-location global-include  
* quicktype >> java  
  --src verifyData.json --src-lang json --top-level verifyDatas  
  --out verifyData.java --lang java  
  --package sheeter  
  --just-types  

[go]: https://go.dev/dl/
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter
[template]: https://pkg.go.dev/text/template

[excel]: doc/image/excel.jpg
[example]: doc/example/example.7z