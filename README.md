![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以[go]做成的excel轉換工具, 前身是[sheet]  
用於將指定格式的excel轉換為[json]資料檔案, [proto]資料檔案, 讀取資料的程式碼; 程式碼目前支援的語言為cs, go  
在windows以及mac通過測試, 但是沒有在linux上測試過  

# 系統需求
* [go]1.18以上
* [proto]3以上

# 安裝說明
* 安裝[go]
* 安裝[sheeter], 在終端執行以下命令
  ```shell
  go install github.com/yinweli/Sheeter/cmd/sheeter@latest
  ```

# 如何使用
* 建立[excel檔案](#excel說明)
* 建立[設定檔案](#設定說明)
* 在終端執行[建置命令](#命令說明)
  ```shell
  sheeter build --config setting.yaml
  ```
* 如果要產生[proto]程式碼, 可以執行產生出來的protoCs.bat/.sh或是protoGo.bat/.sh
* 最後會產生結構程式碼, 讀取器程式碼, 倉庫程式碼讓程式可以控制表格
* 關於程式碼的範例可以看[範例檔案](#範例檔案)

# 範例檔案
[example]

# 產生目錄

| 名稱           | 說明                          |
|:---------------|:------------------------------|
| ./             | 存放建置proto的批次檔/腳本    |
| ./json         | json目錄                      |
| ./json/codeCs  | 存放結構與讀取器程式碼        |
| ./json/codeGo  | 存放結構與讀取器程式碼        |
| ./json/data    | 存放資料檔案                  |
| ./proto        | proto目錄                     |
| ./proto/codeCs | 存放結構與讀取器程式碼        |
| ./proto/codeGo | 存放結構與讀取器程式碼        |
| ./proto/data   | 存放資料檔案                  |
| ./proto/schema | 存放.proto檔案                |
| ./template     | 存放模板檔案                  |

# 命令說明
以下描述了[sheeter]提供的命令與旗標

## help命令
用於顯示命令說明  
```shell
sheeter help [command]
```

## version命令
用於顯示版本資訊  
```shell
sheeter version
```

## build命令
用於建置資料檔案與程式碼  
```shell
sheeter build [flags]
```
例如  
```shell
sheeter build --config setting.yaml
sheeter build --config setting.yaml --json --namespace
sheeter build --config setting.yaml --lineOfField 1 --lineOfLayer 2
```

| 旗標          | 參數                                    | 說明                     |
|:--------------|:----------------------------------------|:-------------------------|
| --config      | 路徑與檔名; 例如: path/seeting.yaml     | 設定檔案路徑             |
| --json        |                                         | 是否產生json檔案         |
| --proto       |                                         | 是否產生proto檔案        |
| --namespace   |                                         | 是否用簡單的命名空間名稱 |
| --lineOfField | 行號(1為起始行)                         | 欄位行號                 |
| --lineOfLayer | 行號(1為起始行)                         | 階層行號                 |
| --lineOfNote  | 行號(1為起始行)                         | 註解行號                 |
| --lineOfData  | 行號(1為起始行)                         | 資料行號                 |
| --excludes    | 標籤,標籤,...                           | 輸出時排除的標籤列表     |
| --elements    | 檔案名稱#表單名稱,檔案名稱#表單名稱,... | 項目列表                 |

* --json / --proto: 用於控制是否要產生[json]與[proto]檔案
    * sheeter build         => 輸出[json]與[proto]檔案
    * sheeter build --json  => 只輸出[json]檔案
    * sheeter build --proto => 只輸出[proto]檔案
* --namespace: 用於控制產生的命名空間名稱
    * sheeter build             => 命名空間名稱: sheeterJson / SheeterJson / sheeterProto / SheeterProto
    * sheeter build --namespace => 命名空間名稱: sheeter / Sheeter

## tmpl命令
用於產生執行時使用的模板檔案, 你可以通過修改模板來改變產生出來的程式碼  
```shell
sheeter tmpl [flags]
```

| 旗標          | 參數 | 說明             |
|:--------------|:-----|:-----------------|
| --clean / -c  |      | 重新產生模板檔案 |

# excel說明
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
* 表格名稱
    * excel與表格的名稱合併後不能以下名稱的組合(不分大小寫)
        * depot
        * loader
        * reader
        * readers
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

# 設定說明
```yaml
global:
  exportJson: true      # 是否產生json檔案
  exportProto: true     # 是否產生proto檔案
  simpleNamespace: true # 是否用簡單的命名空間名稱
  lineOfField: 1        # 欄位行號(1為起始行)
  lineOfLayer: 2        # 階層行號(1為起始行)
  lineOfNote:  3        # 註解行號(1為起始行)
  lineOfData:  4        # 資料行號(1為起始行)
  excludes:             # 排除標籤列表
    - tag1
    - tag2

elements:
  - excel: excel1.xlsx  # excel檔案名稱
    sheet: Data         # excel表單名稱
  - excel: excel2.xlsx
    sheet: Data
  - excel: excel3.xlsx
    sheet: Data
```

# 模板檔案
[sheeter]轉換時會把使用的程式碼模板輸出到`template`目錄下  
使用者可以改變模板內容, 來產生自訂的程式碼  
模板檔案使用[go]的[template]語法, 同時可以參考以下模板參數來做名稱的替換  

## 產生程式碼模板參數
包括struct-cs/go, reader-cs/go, proto-schema等檔案  

| 名稱    | 參數 | 說明                      |
|:--------|:-----|:--------------------------|
|         |      | [全域設定](#全域設定參數) |
|         |      | [綜合工具](#綜合工具參數) |
|         |      | [類型資料](#類型資料參數) |
| .Depend |      | 依賴列表                  |

## 產生後製檔案模板參數
包括depot-cs/go, proto-bat/sh等檔案  

| 名稱    | 參數 | 說明                      |
|:--------|:-----|:--------------------------|
|         |      | [全域設定](#全域設定參數) |
|         |      | [綜合工具](#綜合工具參數) |
| .Struct |      | [結構列表](#結構資料參數) |

## 全域設定參數

| 名稱             | 參數 | 說明                     |
|:-----------------|:-----|:-------------------------|
| .ExportJson      |      | 是否產生json檔案         |
| .ExportProto     |      | 是否產生proto檔案        |
| .SimpleNamespace |      | 是否用簡單的命名空間名稱 |
| .LineOfField     |      | 欄位行號(1為起始行)      |
| .LineOfLayer     |      | 階層行號(1為起始行)      |
| .LineOfNote      |      | 註解行號(1為起始行)      |
| .LineOfData      |      | 資料行號(1為起始行)      |
| .Excludes        |      | 排除標籤列表             |

## 綜合工具參數

| 名稱               | 參數                     | 說明                        |
|:-------------------|:-------------------------|:----------------------------|
| .AppName           |                          | 程式名稱                    |
| .JsonNamespace     | 是否用簡單的命名空間名稱 | json命名空間名稱            |
| .ProtoNamespace    | 是否用簡單的命名空間名稱 | proto命名空間名稱           |
| .StructName        |                          | 結構名稱                    |
| .ReaderName        |                          | 讀取器名稱                  |
| .StorerName        |                          | 儲存器名稱                  |
| .StorerDatas       |                          | 儲存器資料名稱              |
| .StorerMessage     | 是否用簡單的命名空間名稱 | 儲存器proto message名稱     |
| .FieldName         | 欄位資料                 | 欄位名稱                    |
| .FieldNote         | 欄位資料                 | 欄位註解                    |
| .FieldTypeCs       | 欄位資料                 | cs欄位類型                  |
| .FieldTypeGo       | 欄位資料                 | go欄位類型                  |
| .FieldTypeProto    | 欄位資料                 | go欄位類型                  |
| .PkeyTypeCs        |                          | pkey的cs類型                |
| .PkeyTypeGo        |                          | pkey的go類型                |
| .PkeyTypeProto     |                          | pkey的proto類型             |
| .JsonDataName      |                          | json資料名稱                |
| .JsonDataExt       |                          | json資料副檔名              |
| .JsonDataFile      |                          | json資料檔名                |
| .JsonDataPath      |                          | json資料路徑                |
| .JsonCsStructPath  |                          | json-cs結構程式碼路徑       |
| .JsonCsReaderPath  |                          | json-cs讀取器程式碼路徑     |
| .JsonCsDepotPath   |                          | json-cs倉庫程式碼路徑       |
| .JsonGoStructPath  |                          | json-go結構程式碼路徑       |
| .JsonGoReaderPath  |                          | json-go讀取器程式碼檔名路徑 |
| .JsonGoDepotPath   |                          | json-go倉庫程式碼路徑       |
| .ProtoCsPath       |                          | proto-cs路徑                |
| .ProtoGoPath       |                          | proto-go路徑                |
| .ProtoSchemaPath   |                          | proto-schema路徑            |
| .ProtoName         |                          | proto架構檔名               |
| .ProtoPath         |                          | proto架構路徑               |
| .ProtoDataName     |                          | proto資料名稱               |
| .ProtoDataExt      |                          | proto資料副檔名             |
| .ProtoDataFile     |                          | proto資料檔名               |
| .ProtoDataPath     |                          | proto資料路徑               |
| .ProtoCsReaderPath |                          | proto-cs讀取器程式碼路徑    |
| .ProtoCsDepotPath  |                          | proto-cs倉庫程式碼路徑      |
| .ProtoGoReaderPath |                          | proto-go讀取器程式碼路徑    |
| .ProtoGoDepotPath  |                          | proto-go倉庫程式碼路徑      |
| .ProtoCsBatFile    |                          | proto-cs-bat檔名            |
| .ProtoCsShFile     |                          | proto-cs-sh檔名             |
| .ProtoGoBatFile    |                          | proto-go-bat檔名            |
| .ProtoGoShFile     |                          | proto-go-sh檔名             |
| .ProtoDepend       | 依賴名稱                 | proto依賴檔案名稱           |
| .FirstUpper        | 字串                     | 字串首字母大寫              |
| .FirstLower        | 字串                     | 字串首字母小寫              |
| .Add               | 數值1 數值2              | 加法(數值1 + 數值2)         |
| .Sub               | 數值1 數值2              | 減法(數值1 - 數值2)         |
| .Mul               | 數值1 數值2              | 乘法(數值1 x 數值2)         |
| .Div               | 數值1 數值2              | 除法(數值1 / 數值2)         |

## 類型資料參數

| 名稱    | 參數 | 說明                      |
|:--------|:-----|:--------------------------|
| .Excel  |      | excel檔案名稱             |
| .Sheet  |      | excel表格名稱             |
| .Reader |      | 是否要產生讀取器          |
| .Fields |      | [欄位列表](#欄位資料參數) |

## 欄位資料參數

| 名稱   | 參數 | 說明                      |
|:-------|:-----|:--------------------------|
| .Name  |      | 欄位名稱                  |
| .Note  |      | 欄位註解                  |
| .Field |      | [欄位類型](#欄位類型參數) |
| .Alter |      | 欄位類型別名              |
| .Array |      | 陣列旗標                  |

## 欄位類型參數

| 名稱         | 參數 | 說明           |
|:-------------|:-----|:---------------|
| .Type        |      | excel欄位類型  |
| .IsShow      |      | 是否顯示欄位   |
| .IsPkey      |      | 是否是主要索引 |
| .ToTypeCs    |      | cs類型字串     |
| .ToTypeGo    |      | go類型字串     |
| .ToTypeProto |      | proto類型字串  |

## 結構資料參數

| 名稱   | 參數 | 說明                      |
|:-------|:-----|:--------------------------|
|        |      | [綜合工具](#綜合工具參數) |
|        |      | [類型資料](#類型資料參數) |

# proto說明
以下描述了如果要使用[proto]時的資訊  

## proto轉換為cs程式碼
* 安裝[protoc]
* 執行產生出來的.bat/.sh

## proto轉換為go程式碼
* 安裝[go]
* 安裝[protoc]
* 執行以下命令來安裝[protoc-go]外掛
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
* 執行產生出來的.bat/.sh

## 格式化產出的proto檔案(非必要)
* 安裝[go]
* 安裝[buf]
* 執行以下命令來格式化[proto]檔案
```shell
buf format -w 存放proto檔案的路徑
```

## mac執行.sh
可能要先執行`chmod 755 ****.sh`來把變更產生出來的腳本檔案的權限

# 目錄說明

| 目錄                    | 說明                             |
|:------------------------|:---------------------------------|
| doc                     | 說明文件                         |
| cmd/sheeter             | 主程式                           |
| cmd/sheeter/build       | 建置表格命令                     |
| cmd/sheeter/tmpl        | 產生模板命令                     |
| cmd/sheeter/version     | 顯示版本命令                     |
| internal/builds         | 表格轉換                         |
| internal/excels         | 表格組件                         |
| internal/fields         | 欄位組件                         |
| internal/layers         | 階層組件                         |
| internal/layouts        | 布局組件                         |
| internal/mixeds         | 綜合工具                         |
| internal/tmpls          | 模板組件                         |
| internal/utils          | 協助組件                         |
| testdata                | 測試資料                         |
| support                 | 支援專案                         |
| support/benchmark_count | 檔案數量效率測試資料             |
| support/benchmark_size  | 檔案大小效率測試資料             |
| support/example         | 範例資料                         |
| support/handmade        | 手製模板, 用來檢查模板是否有錯誤 |
| support/handmade/.json  | json手製模板                     |
| support/handmade/.proto | proto手製模板                    |
| support/verifycs        | cs程式碼驗證                     |
| support/verifygo        | go程式碼驗證                     |
| support/verifyunity     | unity程式碼驗證                  |
| testdata                | 測試資料                         |

# TODO
* 新增clear功能
* 考慮看看: 把欄位名稱與欄位類型跟標籤分開為不同行
    * 例如: 欄位名稱行, 欄位設定行(欄位類型與標籤)
* 目前的表格讀取方式會把全部需要的表格都讀取進來, 然後分析跟輸出; 但是在大量表格時會使用到大量記憶體, 可能需要想辦法減少記憶體使用量
* 嘗試 https://github.com/tealeg/xlsx
* 嘗試 https://github.com/TheDataShed/xlsxreader
* 產生flatbuffer

[buf]: https://github.com/bufbuild/buf
[go]: https://go.dev/dl/
[json]: https://www.json.org/json-en.html
[proto]: https://github.com/protocolbuffers/protobuf
[protoc-go]: https://github.com/protocolbuffers/protobuf-go
[protoc]: https://github.com/protocolbuffers/protobuf
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter
[template]: https://pkg.go.dev/text/template

[example]: doc/example/example.7z
[excel]: doc/image/excel.jpg