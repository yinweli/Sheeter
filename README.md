![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以go做成的excel轉換工具, 前身是[sheet]  
將以指定格式做好的excel轉換為json, 再利用[quicktype]轉換出程式碼  

# 目錄說明
| 目錄            | 說明        |
|:----------------|:------------|
| doc             | 說明文件    |
| cmd/sheeter     | sheeter命令 |
| internal/builds | 表格轉換    |
| internal/util   | 協助組件    |
| testdata        | 測試資料    |

# 如何安裝
* 安裝[go]
* 安裝[node.js], 這會順便安裝npm
* 把npm的路徑加入系統環境變數的path中
* 安裝[quicktype], 在終端執行以下命令
  ```
  npm install -g quicktype
  ```
* 安裝[sheeter], 在終端執行以下命令
  ```
  go install github.com/yinweli/Sheeter/cmd/sheeter@latest
  ```

# 如何執行
配置好yaml格式的設定檔與excel檔案, 然後在終端執行
```
sheeter build --config 設定檔.yaml
```

# 如何寫設定檔
```
global:
  bom: true            # 輸出的檔案是否使用順序標記(BOM)
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
* json架構檔案: schema\exampleData.schema
* json資料檔案: json\exampleData.json
* json的cs程式碼: json-cs\exampleData\exampleData.cs
* json的cs讀取器程式碼: json-cs\exampleData\exampleDataReader.cs
* json的go程式碼: json-go\exampleData\exampleData.go
* json的go讀取器程式碼: json-go\exampleData\exampleDataReader.go
* 命名空間: exampledata
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
[node.js]: https://nodejs.org/en/
[quicktype]: https://github.com/quicktype/quicktype
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter

[excel]: doc/image/excel.jpg
[example]: doc/example/example.7z