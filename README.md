![license](https://img.shields.io/github/license/yinweli/Sheeter)
![lint](https://github.com/yinweli/Sheeter/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/yinweli/Sheeter/actions/workflows/test.yml/badge.svg)
![codecov](https://codecov.io/gh/yinweli/Sheeter/branch/main/graph/badge.svg?token=LK5HL58LSN)

# Sheeter
以[go]做成的excel轉換工具  
用於將指定格式的excel轉換為[json]資料檔案, [proto]資料檔案, 讀取資料的程式碼, 列舉程式碼  
程式碼目前支援的語言為cs, go  
前身是[sheet]  

# 系統需求
* [go]1.18以上
* [proto]3以上

# 安裝說明
* 安裝[go]
* 安裝[protoc]
* 安裝[protoc-go], 在終端執行以下命令
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```
* 安裝[sheeter], 在終端執行以下命令
  ```sh
  go install github.com/yinweli/Sheeter/cmd/sheeter@latest
  ```

# 如何使用
* 建立[資料表單](#資料表單說明)或是[列舉表單](#列舉表單說明)
* 建立[設定檔案](#設定說明)
* 在終端執行[建置命令](#命令說明)
  ```sh
  sheeter build --config setting.yaml
  ```
* 最後會產生結構程式碼, 讀取器程式碼, 倉庫程式碼讓程式可以控制表格
* 關於程式碼的範例可以看[範例檔案](#範例檔案)
    * cs範例在 example/example.cs
    * go範例在 example/example.go

# 範例檔案
[example]

# 命令說明
以下描述了[sheeter]提供的命令與旗標

## help命令
用於顯示命令說明  
```sh
sheeter help [command]
```

## version命令
用於顯示版本資訊  
```sh
sheeter version
```

## build命令
用於建置資料檔案與程式碼  
```sh
sheeter build [flags]
```
例如  
```sh
sheeter build --config setting.yaml
sheeter build --config setting.yaml --json --namespace
sheeter build --config setting.yaml --lineOfField 1 --lineOfLayer 2
```

| 旗標          | 參數                                    | 說明                     |
|:--------------|:----------------------------------------|:-------------------------|
| --config      | 路徑與檔名; 例如: path/seeting.yaml     | 設定檔案路徑             |
| --json        |                                         | 是否產生json檔案         |
| --proto       |                                         | 是否產生proto檔案        |
| --enum        |                                         | 是否產生enum檔案         |
| --namespace   |                                         | 是否用簡單的命名空間名稱 |
| --lineOfName  | 行號(1為起始行)                         | 名稱行號                 |
| --lineOfNote  | 行號(1為起始行)                         | 註解行號                 |
| --lineOfField | 行號(1為起始行)                         | 欄位行號                 |
| --lineOfLayer | 行號(1為起始行)                         | 階層行號                 |
| --lineOfData  | 行號(1為起始行)                         | 資料行號                 |
| --lineOfEnum  | 行號(1為起始行)                         | 列舉行號                 |
| --excludes    | 標籤,標籤,...                           | 輸出時排除的標籤列表     |
| --inputs      | 路徑,檔案名稱,檔案名稱#表單名稱,...     | 輸入列表                 |

### --config
從設定檔讀取參數, 設定檔中的參數都可以被其他的旗標值替代  
```sh
sheeter build --config setting.yaml --lineOfName 5
```
像是這種情況, 設定檔中的`lineOfName`的值就會被`--lineOfName`的值替代  

### --json / --proto
用於控制是否要產生[json]與[proto]檔案  
* sheeter build  
  輸出[json]與[proto]檔案  
* sheeter build --json  
  只輸出[json]檔案  
* sheeter build --proto  
  只輸出[proto]檔案  
  
### --namespace
用於控制產生的命名空間名稱  
* sheeter build  
  命名空間名稱: sheeterJson / SheeterJson / sheeterProto / SheeterProto / sheeterEnum / SheeterEnum  
* sheeter build --namespace  
  命名空間名稱: sheeter / Sheeter  
  
### --inputs
輸入列表, 可用以下幾種格式組合, 每個項目以`,`分隔; 注意程式只會讀取副檔名為xlsx的檔案  
* 路徑名稱  
  path, path/, path/path...  
* 檔案名稱  
  example.xlsx, path/example.xlsx...  
* 檔案名稱+表單名稱  
  example.xlsx#sheet, path/example.xlsx#sheet...  
  這個格式中, 需用`#`把檔案名稱與表單名稱隔開  

## tmpl命令
用於產生執行時使用的模板檔案, 你可以通過修改模板來改變產生出來的程式碼  
```sh
sheeter tmpl [flags]
```

| 旗標         | 參數 | 說明             |
|:-------------|:-----|:-----------------|
| --clean / -c |      | 重新產生模板檔案 |

# 資料表單說明
![exceldata]

## 表單名稱
需以`@`開頭, [sheeter]會自動辨識以`@`開頭的表單為資料表單

## 名稱行
欄位名稱, 只能是英文與數字與`_`的組合, 並且不能以數字開頭, 也不允許空白

## 註解行
單行註解, 若為空格就輸出空註解  

## 欄位行
欄位類型與標籤設置, 格式為`類型`或是`類型#標籤`, 空格之後的欄位不會輸出  
一個欄位只能設置一個標籤  

| 類型        | 說明                                 |
|:------------|:-------------------------------------|
| empty       | 不會輸出的欄位                       |
| pkey        | 表格主要索引, 編號可跳號但是不可重複 |
| bool        | 布林值                               |
| boolArray   | 以逗號分隔的布林值陣列               |
| int         | 64位元整數                           |
| intArray    | 以逗號分隔的64位元整數陣列           |
| float       | 64位元浮點數                         |
| floatArray  | 以逗號分隔的64位元整數陣列           |
| string      | 字串                                 |
| stringArray | 以逗號分隔的字串陣列                 |

## 欄位行範例

| 範例     | 欄位類型   | 標籤 |
|:---------|:-----------|:-----|
| pkey     | pkey       |      |
| string   | string     |      |
| string#A | string     | A    |

## 標籤與排除機制
欄位行可用標籤來控制欄位與其資料是否要輸出到資料檔案  
當設定檔中的排除標籤與欄位的標籤符合時, 該欄位就不會輸出到資料檔案  
欄位若是沒有設定標籤, 則永遠會輸出到資料檔案  
當欄位的類型是`empty`, 則永遠不會輸出到資料檔案  
標籤與排除設置不會影響產生的程式碼  
一個欄位只能設置一個標籤  

## 階層行
欄位結構布局, 格式有`{名稱`, `{[]名稱`, `/`, `}`, 之間以空格分隔  

| 格式        | 說明                                |
|:------------|:------------------------------------|
| {結構名稱   | 結構的開始                          |
| {[]陣列名稱 | 陣列的開始                          |
| /           | 分隔陣列                            |
| }           | 結構/陣列結束, 可以連續結束, 如`}}` |

## 階層行範例

| 範例          | 說明                                                 |
|:--------------|:-----------------------------------------------------|
| {Item         | 建立Item結構                                         |
| {[]Item       | 建立以Item結構為元素的陣列                           |
| {Reward {Item | 建立Reward結構, Item結構; Item結構是Reward結構的成員 |

## 資料行
依照類型填寫相應的內容即可, 其中`empty`, `string`, `stringArray`這三種類型允許空格, 其他類型的空格會造成錯誤  
空表格(也就是沒有任何資料行)是允許的  
轉換時, 只會轉換到第一個空行為止  

## 其他的限制
* 檔案名稱與表單名稱
    * 不能是規定的[關鍵字](#關鍵字)
* 表格設置
    * 表格必須有名稱行, 註解行, 欄位行, 階層行, 但是可以不需要有資料行
    * 名稱行, 註解行, 欄位行, 階層行必須在資料行之前
    * 設定檔中必須設定好名稱行, 註解行, 欄位行, 階層行, 資料行的位置
    * 設定檔中行數是從1開始的
* 主索引
    * 表格必須有`pkey`欄位
    * 表格只能有一個`pkey`欄位
    * `pkey`欄位中的內容不能重複
* 階層
    * 不屬於結構/陣列的欄位名稱不能重複(包括`empty`欄位)
    * 結構/陣列名稱可以重複, 重複的結構/陣列的欄位會合併
    * 結構/陣列的欄位可以不必填上所有的名稱
        * 第一個表格設定了結構/陣列欄位: `data { field1, field2, field3 }`
        * 另一個表格同樣使用了`data`結構/陣列, 而欄位只設定 `data { field1, field2 }`, 忽略了`field3`

# 列舉表單說明
![excelenum]

## 表單名稱
需以`$`開頭, [sheeter]會自動辨識以`$`開頭的表單為列舉表單

## 名稱行
實際上不寫也沒關係, 僅提供給使用者辨識用

## 資料行
必須是以下欄位格式  
* 第一欄: 列舉名稱
    * 只能是英文與數字與`_`的組合, 並且不能以數字開頭, 也不允許空白
    * 列舉名稱不允許重複
    * 不能是規定的[關鍵字](#關鍵字)
* 第二欄: 列舉編號
    * 只能是數字
    * 索引編號不允許重複
* 第三欄: 列舉註解
    * 單行註解, 此欄空白也可以

# 設定說明
```yaml
global:
  exportJson:      true # 是否產生json檔案
  exportProto:     true # 是否產生proto檔案
  exportEnum:      true # 是否產生enum檔案
  simpleNamespace: true # 是否用簡單的命名空間名稱
  lineOfName:      1    # 名稱行號(1為起始行)
  lineOfNote:      2    # 註解行號(1為起始行)
  lineOfField:     3    # 欄位行號(1為起始行)
  lineOfLayer:     4    # 階層行號(1為起始行)
  lineOfData:      5    # 資料行號(1為起始行)
  lineOfEnum:      2    # 列舉行號(1為起始行)
  excludes:             # 排除標籤列表
    - tag1
    - tag2

inputs:                   # 輸入列表
  - path1                 # 轉換path1目錄底下符合規格的excel檔案
  - path2                 # 轉換path2目錄底下符合規格的excel檔案
  - path/excel.xlsx       # 轉換指定的excel檔案內符合規格的表單
  - path/excel.xlsx#@Data # 轉換指定的excel檔案內的@Data表單
  - path/excel.xlsx#$Enum # 轉換指定的excel檔案內的$Enum表單
```

# 關鍵字
檔案名稱與表單名稱合併之後的名稱不能是以下名稱的組合(不分大小寫)  
* depot
* loader
* reader
* readers

# 產生目錄

| 名稱         | 說明               |
|:-------------|:-------------------|
| json         | json目錄           |
| json/codeCs  | 存放產生的cs程式碼 |
| json/codeGo  | 存放產生的go程式碼 |
| json/data    | 存放資料檔案       |
| proto        | proto目錄          |
| proto/codeCs | 存放產生的cs程式碼 |
| proto/codeGo | 存放產生的go程式碼 |
| proto/data   | 存放資料檔案       |
| proto/schema | 存放.proto檔案     |
| enum         | enum目錄           |
| enum/codeCs  | 存放列舉程式碼     |
| enum/codeGo  | 存放列舉程式碼     |
| enum/schema  | 存放.proto檔案     |
| template     | 存放模板檔案       |

# 模板檔案
[sheeter]轉換時會把使用的程式碼模板輸出到`template`目錄下  
使用者可以改變模板內容, 來產生自訂的程式碼  
模板檔案使用[go]的[template]語法, 同時可以參考以下模板參數來做名稱的替換  

## json結構, 讀取器模板參數
影響的檔案: json-struct-cs/go, json-reader-cs/go  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域設定](#全域設定參數)                                 |
|                    |             | [命名工具](#命名工具參數)                                 |
|                    |             | [欄位命名工具](#欄位命名工具參數)                         |
|                    |             | [json命名工具](#json命名工具參數)                         |
| .Reader            |             | 是否要產生讀取器                                          |
| .Fields            |             | [欄位列表](#欄位類型參數)                                 |

## json倉庫模板參數
影響的檔案: json-depot-cs/go  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域設定](#全域設定參數)                                 |
|                    |             | [命名工具](#命名工具參數)                                 |
|                    |             | [json命名工具](#json命名工具參數)                         |
| .Struct            |             | 結構列表(註)                                              |

結構列表內容為  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [命名工具](#命名工具參數)                                 |
| Reader             |             | 是否要產生讀取器                                          |

## proto結構, 讀取器, proto架構模板參數
影響的檔案: proto-struct-cs/go, proto-reader-cs/go, proto-schema  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域設定](#全域設定參數)                                 |
|                    |             | [命名工具](#命名工具參數)                                 |
|                    |             | [欄位命名工具](#欄位命名工具參數)                         |
|                    |             | [proto命名工具](#proto命名工具參數)                       |
| .Reader            |             | 是否要產生讀取器                                          |
| .Fields            |             | [欄位列表](#欄位類型參數)                                 |
| .Depend            |             | 依賴列表                                                  |

## proto倉庫模板參數
影響的檔案: proto-depot-cs/go  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域設定](#全域設定參數)                                 |
|                    |             | [命名工具](#命名工具參數)                                 |
|                    |             | [proto命名工具](#proto命名工具參數)                       |
| .Struct            |             | 結構列表(註)                                              |

結構列表內容為  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [命名工具](#命名工具參數)                                 |
| Reader             |             | 是否要產生讀取器                                          |

## enum架構模板參數
影響的檔案: enum-schema  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域設定](#全域設定參數)                                 |
|                    |             | [命名工具](#命名工具參數)                                 |
|                    |             | [enum命名工具](#enum命名工具參數)                         |
| .Enums             |             | 列舉列表(註)                                              |

結構列表內容為  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| Name               |             | 列舉名稱                                                  |
| Index              |             | 列舉編號                                                  |
| Comment            |             | 列舉說明                                                  |

## 全域設定參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ExportJson        |             | 是否產生json檔案                                          |
| .ExportProto       |             | 是否產生proto檔案                                         |
| .ExportEnum        |             | 是否產生enum檔案                                          |
| .SimpleNamespace   |             | 是否用簡單的命名空間名稱                                  |
| .LineOfName        |             | 名稱行號(1為起始行)                                       |
| .LineOfNote        |             | 註解行號(1為起始行)                                       |
| .LineOfField       |             | 欄位行號(1為起始行)                                       |
| .LineOfLayer       |             | 階層行號(1為起始行)                                       |
| .LineOfData        |             | 資料行號(1為起始行)                                       |
| .LineOfEnum        |             | 列舉行號(1為起始行)                                       |
| .Excludes          |             | 排除標籤列表                                              |

## 命名工具參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ExcelName         |             | excel名稱                                                 |
| .SheetName         |             | sheet名稱                                                 |
| .AppName           |             | 程式名稱                                                  |
| .JsonNamespace     | bool        | json命名空間名稱, 參數影響是否用簡單的命名空間名稱        |
| .ProtoNamespace    | bool        | proto命名空間名稱, 參數影響是否用簡單的命名空間名稱       |
| .EnumNamespace     | bool        | enum命名空間名稱, 參數影響是否用簡單的命名空間名稱        |
| .StructName        |             | 結構名稱                                                  |
| .ReaderName        |             | 讀取器名稱                                                |
| .StorerName        |             | 儲存器名稱                                                |
| .StorerDatas       |             | 儲存器資料名稱                                            |
| .StorerMessage     | bool        | 儲存器proto message名稱, 參數影響是否用簡單的命名空間名稱 |
| .FirstUpper        | 字串        | 字串首字母大寫                                            |
| .FirstLower        | 字串        | 字串首字母小寫                                            |
| .Add               | 數值1 數值2 | 加法(數值1 + 數值2)                                       |
| .Sub               | 數值1 數值2 | 減法(數值1 - 數值2)                                       |
| .Mul               | 數值1 數值2 | 乘法(數值1 x 數值2)                                       |
| .Div               | 數值1 數值2 | 除法(數值1 / 數值2)                                       |

## 欄位命名工具參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .FieldName         | 欄位資料    | 欄位名稱                                                  |
| .FieldNote         | 欄位資料    | 欄位註解                                                  |
| .FieldTypeCs       | 欄位資料    | cs欄位類型                                                |
| .FieldTypeGo       | 欄位資料    | go欄位類型                                                |
| .FieldTypeProto    | 欄位資料    | proto欄位類型                                             |
| .PkeyTypeCs        |             | pkey的cs類型                                              |
| .PkeyTypeGo        |             | pkey的go類型                                              |
| .PkeyTypeProto     |             | pkey的proto類型                                           |

## json命名工具參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .JsonDataName      |             | json資料名稱                                              |
| .JsonDataExt       |             | json資料副檔名                                            |
| .JsonDataFile      |             | json資料檔名                                              |
| .JsonDataPath      |             | json資料路徑                                              |
| .JsonStructCsPath  |             | json-cs結構程式碼路徑                                     |
| .JsonReaderCsPath  |             | json-cs讀取器程式碼路徑                                   |
| .JsonDepotCsPath   |             | json-cs倉庫程式碼路徑                                     |
| .JsonStructGoPath  |             | json-go結構程式碼路徑                                     |
| .JsonReaderGoPath  |             | json-go讀取器程式碼檔名路徑                               |
| .JsonDepotGoPath   |             | json-go倉庫程式碼路徑                                     |

## proto命名工具參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ProtoCsPath       |             | proto-cs路徑                                              |
| .ProtoGoPath       |             | proto-go路徑                                              |
| .ProtoSchemaPath   |             | proto-schema路徑                                          |
| .ProtoName         |             | proto架構檔名                                             |
| .ProtoPath         |             | proto架構路徑                                             |
| .ProtoDataName     |             | proto資料名稱                                             |
| .ProtoDataExt      |             | proto資料副檔名                                           |
| .ProtoDataFile     |             | proto資料檔名                                             |
| .ProtoDataPath     |             | proto資料路徑                                             |
| .ProtoReaderCsPath |             | proto-cs讀取器程式碼路徑                                  |
| .ProtoDepotCsPath  |             | proto-cs倉庫程式碼路徑                                    |
| .ProtoReaderGoPath |             | proto-go讀取器程式碼路徑                                  |
| .ProtoDepotGoPath  |             | proto-go倉庫程式碼路徑                                    |
| .ProtoDepend       | 依賴名稱    | proto依賴檔案名稱                                         |

## enum命名工具參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .EnumCsPath        |             | enum-cs路徑                                               |
| .EnumGoPath        |             | enum-go路徑                                               |
| .EnumSchemaPath    |             | enum-schema路徑                                           |
| .EnumName          |             | enum架構檔名                                              |
| .EnumPath          |             | enum架構路徑                                              |

## 欄位類型參數

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .Name              |             | 欄位名稱                                                  |
| .Note              |             | 欄位註解                                                  |
| .Field             |             | 欄位類型(註)                                              |
| .Alter             |             | 欄位類型別名                                              |
| .Array             |             | 陣列旗標                                                  |

欄位類型內容為  

| 名稱               | 參數        | 說明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .Type              |             | excel欄位類型                                             |
| .IsShow            |             | 是否顯示欄位                                              |
| .IsPkey            |             | 是否是主要索引                                            |
| .ToTypeCs          |             | cs類型字串                                                |
| .ToTypeGo          |             | go類型字串                                                |
| .ToTypeProto       |             | proto類型字串                                             |

# 格式化程式碼
[sheeter]並不負責幫產生的檔案排版, 如果需要排版, 就需要自己寫.bat/.sh來執行  
以下介紹cs, go, proto的排版工具, 有需要可以自己去安裝  

## csharpier
用於cs的排版工具  
* 安裝
    * 安裝[dotnet], 如果有安裝.net sdk, 或有安裝unity可能可以省略此步驟
    * 安裝[csharpier], 在終端執行以下命令
      ```sh
      dotnet tool install csharpier -g
      ```
* 使用
    * 在終端執行以下命令
      ```sh
      dotnet csharpier .
      ```

## gofmt
用於go的排版工具  
* 安裝
    * 安裝[go]時會順便安裝
* 使用
    * 在終端執行以下命令
      ```sh
      gofmt -w .
      ```

## buf
用於proto的排版工具  
* 安裝
    * 安裝[buf], 在終端執行以下命令
      ```sh
      go install github.com/bufbuild/buf/cmd/buf@v1.8.0
      ```
* 使用
    * 在終端執行以下命令
      ```sh
      buf format -w .
      ```

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
| internal/nameds         | 命名工具                         |
| internal/pipelines      | 管線組件                         |
| internal/tmpls          | 模板組件                         |
| internal/utils          | 協助組件                         |
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

[buf]: https://github.com/bufbuild/buf
[csharpier]: https://github.com/belav/csharpier
[dotnet]: https://learn.microsoft.com/zh-tw/dotnet/core/sdk
[go]: https://go.dev/dl/
[json]: https://www.json.org/json-en.html
[proto]: https://github.com/protocolbuffers/protobuf
[protoc-go]: https://github.com/protocolbuffers/protobuf-go
[protoc]: https://github.com/protocolbuffers/protobuf
[sheet]: https://github.com/yinweli/Sheet
[sheeter]: https://github.com/yinweli/sheeter
[template]: https://pkg.go.dev/text/template

[example]: doc/example/example.7z
[exceldata]: doc/image/exceldata.jpg
[excelenum]: doc/image/excelenum.jpg