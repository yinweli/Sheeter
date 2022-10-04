![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以[go]做成的excel轉換工具, 前身是[sheet]  
用於將指定格式的excel轉換為json, cs程式碼, go程式碼, proto檔案等

# 目錄說明
| 目錄                | 說明         |
|:--------------------|:-------------|
| doc                 | 說明文件     |
| cmd/sheeter         | 主程式       |
| cmd/sheeter/build   | 建置表格命令 |
| cmd/sheeter/tmpl    | 產生模板命令 |
| cmd/sheeter/version | 顯示版本命令 |
| internal/builds     | 表格轉換     |
| internal/fields     | 欄位組件     |
| internal/layers     | 階層組件     |
| internal/layouts    | 布局組件     |
| internal/mixeds     | 綜合工具     |
| internal/tmpls      | 模板組件     |
| internal/utils      | 協助組件     |
| testdata            | 測試資料     |
| verify/benchmark    | 效率測試資料 |
| verify/example      | 範例資料     |
| verify/verifycs     | cs程式碼驗證 |
| verify/verifygo     | go程式碼驗證 |

# 系統需求
* [go]1.18以上
* [proto]3以上

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
  build命令有以下的參數可用  
  | 參數          | 說明                                                          |
  |:--------------|:--------------------------------------------------------------|
  | --config      | 設定檔案路徑                                                  |
  | --lineOfField | 欄位行號                                                      |
  | --lineOfLayer | 階層行號                                                      |
  | --lineOfNote  | 註解行號                                                      |
  | --lineOfData  | 資料行號                                                      |
  | --excludes    | 輸出時排除的標籤列表, 以標籤名稱,標籤名稱,...的方式填寫       |
  | --elements    | 項目列表, 以檔案名稱#表單名稱,檔案名稱#表單名稱,...的方式填寫 |
  
* tmpl: 產生模板檔案  
  這會產生執行時使用的模板檔案, 你可以通過修改模板來改變產生出來的程式碼  
  執行建置表格命令時也會產生模板檔案  
  ```shell
  sheeter tmpl
  ```
  tmpl命令有以下的參數可用  
  | 參數          | 說明                         |
  |:--------------|:-----------------------------|
  | --clean / -c  | 在產生模板前, 把模板檔案刪除 |
* version: 顯示版本資訊  
  ```shell
  sheeter version
  ```

# 如何寫設定檔
```yaml
global:
  lineOfField: 1       # 欄位行號(1為起始行)
  lineOfLayer: 2       # 階層行號(1為起始行)
  lineOfNote:  3       # 註解行號(1為起始行)
  lineOfData:  4       # 資料行號(1為起始行)
  excludes:            # 排除標籤列表
    - tag1
    - tag2

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
欄位的格式為`名稱#格式`或是`名稱#格式#標籤`, 空格之後的欄位不會輸出  
可在欄位行中用標籤來控制欄位與其資料是否要輸出到資料檔案  
當欄位的標籤符合設定檔中的排除標籤列表的時候, 該欄位就不會輸出到資料檔案  
欄位若是沒有設定標籤, 則永遠不會被排除; 亦即一定會輸出到資料檔案  
標籤與排除設置不會影響產生的程式碼, 程式碼中永遠會有除了`empty`類型以外的欄位  

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

## 欄位行範例
| 範例        | 欄位名稱 | 欄位類型 |
|:------------|:---------|:---------|
| itemID#pkey | itemID   | pkey     |
| enable#bool | enable   | bool     |
| gold#int    | gold     | int      |
| note#text   | note     | text     |
| hide#text#A | note     | text     |

## 階層行
欄位結構布局, 格式有`{名稱`, `{[]名稱`, `/`, `}`, 之間以空格分隔  

| 格式        | 說明                                 |
|:------------|:-------------------------------------|
| {結構名稱   | 結構的開始                           |
| {[]陣列名稱 | 陣列的開始                           |
| /           | 分隔陣列                             |
| }           | 結構/陣列結束, 可以連續結束, 如`}}`  |

## 階層行範例
| 範例          | 說明                                                 |
|:--------------|:-----------------------------------------------------|
| {Item         | 建立Item結構                                         |
| {[]Item       | 建立以Item結構為元素的陣列                           |
| {Reward {Item | 建立Reward結構, Item結構; Item結構是Reward結構的成員 |


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
根目錄會存放建置程式碼的批次檔與腳本檔案  

| 名稱         | 說明                   |
|:-------------|:-----------------------|
| json         | json根目錄             |
| json/codeCs  | 存放結構與讀取器程式碼 |
| json/codeGo  | 存放結構與讀取器程式碼 |
| json/data    | 存放資料檔案           |
| proto        | proto根目錄            |
| proto/codeCs | 存放結構與讀取器程式碼 |
| proto/codeGo | 存放結構與讀取器程式碼 |
| proto/data   | 存放資料檔案           |
| proto/schema | 存放.proto檔案         |
| template     | 存放模板檔案           |

# 範例檔案
[example]

# 模板檔案
sheeter轉換時會把使用的程式碼模板輸出到`template`目錄下  
使用者可以改變模板內容, 來產生自訂的程式碼  
當需要重置模板時(例如[sheeter]更新版本時), 可以在終端執行以下命令重置模板  
```shell
sheeter tmpl -c
```
模板檔案使用golang的[template]語法, 同時可以參考以下變數來做結構名稱或是欄位名稱等的替換  

| 名稱                | 參數        | 說明                           |
|:--------------------|:------------|:-------------------------------|
| $.AppName           |             | 程式名稱                       |
| $.NamespaceJson     |             | json命名空間名稱               |
| $.NamespaceProto    |             | proto命名空間名稱              |
| $.StructName        |             | 結構名稱                       |
| $.ReaderName        |             | 讀取器名稱                     |
| $.StorerName        |             | 儲存器名稱                     |
| $.StorerDatas       |             | 儲存器資料名稱                 |
| $.StorerMessage     |             | 儲存器proto message名稱        |
| $.FieldName         | 欄位資料    | 欄位名稱                       |
| $.FieldNote         | 欄位資料    | 欄位註解                       |
| $.FieldTypeCs       | 欄位資料    | cs欄位類型                     |
| $.FieldTypeGo       | 欄位資料    | go欄位類型                     |
| $.FieldTypeProto    | 欄位資料    | go欄位類型                     |
| $.PkeyTypeCs        |             | pkey的cs類型                   |
| $.PkeyTypeGo        |             | pkey的go類型                   |
| $.PkeyTypeProto     |             | pkey的proto類型                |
| $.FileJsonData      |             | json資料檔名                   |
| $.PathJsonData      |             | json資料路徑                   |
| $.PathJsonCsStruct  |             | json-cs結構程式碼路徑          |
| $.PathJsonCsReader  |             | json-cs讀取器程式碼路徑        |
| $.PathJsonGoStruct  |             | json-go結構程式碼路徑          |
| $.PathJsonGoReader  |             | json-go讀取器程式碼檔名路徑    |
| $.PathProtoCs       |             | proto-cs路徑                   |
| $.PathProtoGo       |             | proto-go路徑                   |
| $.PathProtoSchema   |             | proto-schema路徑               |
| $.FileProtoName     |             | proto架構檔名                  |
| $.PathProtoName     |             | proto架構路徑                  |
| $.FileProtoData     |             | proto資料檔名                  |
| $.PathProtoData     |             | proto資料路徑                  |
| $.PathProtoCsReader |             | proto-cs讀取器程式碼路徑       |
| $.PathProtoGoReader |             | proto-go讀取器程式碼路徑       |
| $.ProtoDepend       | 依賴名稱    | proto依賴檔案名稱              |
| $.FirstUpper        | 字串        | 字串首字母大寫                 |
| $.FirstLower        | 字串        | 字串首字母小寫                 |
| $.Add               | 數值1 數值2 | 加法(數值1 + 數值2)            |
| $.Sub               | 數值1 數值2 | 減法(數值1 - 數值2)            |
| $.Mul               | 數值1 數值2 | 乘法(數值1 x 數值2)            |
| $.Div               | 數值1 數值2 | 除法(數值1 / 數值2)            |
| $.Reader            |             | 是否要產生讀取器(部分模板可用) |
| $.Fields            |             | 欄位列表(部分模板可用)         |
| $.Depend            |             | 依賴列表(部分模板可用)         |

# protobuf支援
以下描述了如果要使用protobuf時的資訊

## proto轉換為cs程式碼
* 安裝[protoc]
* 執行產生出來的.bat/.sh

## proto轉換為go程式碼
* 安裝[go]
* 安裝[protoc]
* 執行以下命令來安裝protobuf的[protoc-go]外掛
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
* 執行產生出來的.bat/.sh

## 格式化產出的proto檔案(非必要)
* 安裝[go]
* 安裝[buf]
* 執行以下命令來格式化proto檔案
```shell
buf format -w 存放proto檔案的路徑
```

# TODO
* reader
    * reader的merge系列功能(為了mod)
    * 當merge時有重複索引, 要報告重複索引列表
* 全域讀取器
    * 可以從多個來源路徑讀取資料
    * 當有複數個來源都是相同檔案時, 採用合併方式
    * 當合併時有重複索引, 仍然會持續執行剩餘表格的讀取, 但是最後會報告結果(包含檔名與重複索引列表)
* 寫個mac的轉檔腳本(類似rebuild.bat)
* .bat跟.sh都需要加權限+chmod
* 把excel層抽象出來, 方便以後換excel組件
* 目前的表格讀取方式會把全部需要的表格都讀取進來, 然後分析跟輸出; 但是在大量表格時會使用到大量記憶體, 可能需要想辦法減少記憶體使用量
* 嘗試 https://github.com/tealeg/xlsx
* 嘗試 https://github.com/TheDataShed/xlsxreader
* 從NewtownJson遷移到Unity官方Json, 或是採行2種皆可的方式
* 考慮如果用google sheet當輸入資料的話呢?
* 產生flatbuffer message
* 產生flatbuffer bytes data

[buf]: https://github.com/bufbuild/buf
[go]: https://go.dev/dl/
[proto]: https://github.com/protocolbuffers/protobuf
[protoc-go]: https://github.com/protocolbuffers/protobuf-go
[protoc]: https://github.com/protocolbuffers/protobuf
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter
[template]: https://pkg.go.dev/text/template

[example]: doc/example/example.7z
[excel]: doc/image/excel.jpg