![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以[go]做成的excel轉換工具, 前身是[sheet]  
用於將指定格式的excel轉換為json, cs程式碼, go程式碼, proto檔案等

# 目錄說明
| 目錄                | 說明           |
|:--------------------|:---------------|
| doc                 | 說明文件       |
| cmd/sheeter         | 主程式         |
| cmd/sheeter/build   | 建置表格命令   |
| cmd/sheeter/tmpl    | 產生模板命令   |
| cmd/sheeter/version | 顯示版本命令   |
| internal/builds     | 表格轉換       |
| internal/fields     | 欄位組件       |
| internal/layers     | 階層組件       |
| internal/layouts    | 布局組件       |
| internal/mixeds     | 綜合工具       |
| internal/tmpls      | 模板組件       |
| internal/utils      | 協助組件       |
| testdata            | 測試資料       |
| verify/example      | 範例資料       |
| verify/verifycs     | cs程式碼驗證   |
| verify/verifygo     | go程式碼驗證   |

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
* tmpl: 產生模板檔案  
  這會產生執行時使用的模板檔案, 你可以通過修改模板來改變產生出來的程式碼  
  執行建置表格命令時也會產生模板檔案  
  ```shell
  sheeter tmpl
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

## 其他的限制
* 表格設置
    * 表格必須有欄位行, 階層行, 註解行, 但是可以不需要有資料行
    * 欄位行, 階層行, 註解行必須在資料行之前
    * 設定檔中必須設定好欄位行, 階層行, 註解行, 資料行的位置
    * 設定檔中行數是從1開始的
* 主索引
    * 表格必須有`pkey`欄位
    * 表格只能有一個`pkey`欄位
    * `pkey`欄位中的內容不能重複
* 欄位
    * 欄位名稱不能重複(包括`empty`欄位)
* 階層
    * 結構/陣列名稱可以重複, 重複的結構/陣列的欄位會合併
    * 結構/陣列的欄位可以不必填上所有的名稱
        * 第一個表格設定了結構/陣列欄位: `data { field1, field2, field3 }`
        * 另一個表格同樣使用了`data`結構/陣列, 而欄位只設定 `data { field1, field2 }`, 忽略了`field3`

# 產生目錄
| 名稱        | 說明                                             |
|:------------|:-------------------------------------------------|
| json        | json根目錄                                       |
| json/.data  | 存放資料檔案                                     |
| json/cs     | 存放結構與讀取器程式碼                           |
| json/go     | 存放結構與讀取器程式碼                           |
| proto       | proto根目錄; 存放.proto檔案, build.bat, build.sh |
| proto/.data | 存放資料檔案                                     |
| proto/cs    | 存放結構與讀取器程式碼                           |
| proto/go    | 存放結構與讀取器程式碼                           |
| template    | 模板檔案                                         |

# 轉換範例
[example]

# 關於模板檔案
sheeter轉換時會把使用的程式碼模板輸出到`template`目錄下  
使用者可以改變模板內容, 來產生自訂的程式碼  
當需要重置模板時(例如[sheeter]更新版本時), 可以在終端執行以下命令重置模板  
```shell
sheeter tmpl -c
```
模板檔案使用golang的[template]語法, 同時可以參考以下變數來做結構名稱或是欄位名稱等的替換  

| 名稱                | 說明                                                                       |
|:--------------------|:---------------------------------------------------------------------------|
| $.AppName           | 程式名稱                                                                   |
| $.Namespace         | 命名空間名稱                                                               |
| $.StructName        | 結構名稱                                                                   |
| $.ReaderName        | 讀取器名稱                                                                 |
| $.FileJsonCsStruct  | json-cs結構程式碼檔名路徑                                                  |
| $.FileJsonCsReader  | json-cs讀取器程式碼檔名路徑                                                |
| $.FileJsonGoStruct  | json-go結構程式碼檔名路徑                                                  |
| $.FileJsonGoReader  | json-go讀取器程式碼檔名路徑                                                |
| $.FileJsonDataName  | json資料檔案名稱                                                           |
| $.FileJsonDataPath  | json資料檔名路徑                                                           |
| $.PathProtoSchema   | proto架構路徑                                                              |
| $.PathProtoCs       | proto-cs路徑                                                               |
| $.PathProtoGo       | proto-go路徑                                                               |
| $.FileProtoSchema   | proto架構檔名路徑                                                          |
| $.FileProtoCsReader | proto-cs讀取器程式碼檔名路徑                                               |
| $.FileProtoGoReader | proto-go讀取器程式碼檔名路徑                                               |
| $.FileProtoDataName | proto資料檔案名稱                                                          |
| $.FileProtoDataPath | proto資料檔名路徑                                                          |
| $.FileProtoBat      | proto-bat檔名路徑                                                          |
| $.FileProtoSh       | proto-sh檔名路徑                                                           |
| $.ProtoDepend       | proto依賴檔案名稱 [依賴名稱]                                               |
| $.FieldName         | 取得欄位名稱 [欄位資料]                                                    |
| $.FieldNote         | 取得欄位註解 [欄位資料]                                                    |
| $.FieldTypeCs       | 取得cs欄位類型 [欄位資料]                                                  |
| $.FieldTypeGo       | 取得go欄位類型 [欄位資料]                                                  |
| $.FieldTypeProto    | 取得go欄位類型 [欄位資料]                                                  |
| $.Add               | 加法(v1 + v2) [數值 數值]                                                  |
| $.Sub               | 減法(v1 - v2) [數值 數值]                                                  |
| $.Mul               | 乘法(v1 x v2) [數值 數值]                                                  |
| $.Div               | 除法(v1 / v2) [數值 數值]                                                  |
| $.Fields            | 欄位列表, 只有json-cs-struct.txt, json-go-struct.txt, proto-schema.txt可用 |
| $.Depend            | 依賴列表, 只有proto-schema.txt可用                                         |

# 關於proto轉換為cs程式碼
* 安裝[protoc]
* 執行產生出來的.bat/.sh

# 關於proto轉換為go程式碼
* 安裝[go]
* 安裝[protoc]
* 執行以下命令來安裝protobuf的[protoc-go]外掛
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
* 執行產生出來的.bat/.sh

# 關於格式化產出的proto檔案(為了美觀!)
* 安裝[buf]
* 執行以下命令來格式化proto檔案
```shell
buf format -w 存放proto檔案的路徑
```

# TODO
* 產生protobuffer bytes data
* 產生flatbuffer message
* 產生flatbuffer bytes data

# 暫時紀錄
* builds encoding [multi thread <每個config.element>]
    * proto
* .pbd     => encodingProto          : multi thread <runtimeSector>
* .bat/.sh => poststepCs, poststepGo : all runtimeStruct
* https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#Unmarshal
* https://github.com/jhump/protoreflect
* https://cloud.tencent.com/developer/article/1542624

[buf]: https://github.com/bufbuild/buf
[go]: https://go.dev/dl/
[protoc-go]: https://github.com/protocolbuffers/protobuf-go
[protoc]: https://github.com/protocolbuffers/protobuf
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter
[template]: https://pkg.go.dev/text/template

[example]: doc/example/example.7z
[excel]: doc/image/excel.jpg