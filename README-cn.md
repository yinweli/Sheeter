# Sheeter
以[go]做成的excel转换工具  
将指定格式的excel转换为json/proto资料档案, 以及产生用于读取资料档案所需的cs/go程式码  
前身是[sheet]  

# 系统需求
* [go]1.18以上
* [proto]3以上

# 安装说明
* 安装[go]
* 安装[protoc]
* 安装[protoc-go], 在终端执行以下命令
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```
* 安装[sheeter], 在终端执行以下命令
  ```sh
  go install github.com/yinweli/Sheeter/cmd/sheeter@latest
  ```

# 如何使用
* 建立[资料表单](#资料表单说明)或是[列举表单](#列举表单说明)
* 建立[设定档案](#设定说明)
* 在终端执行[建置命令](#命令说明)
  ```sh
  sheeter build --config 设定档案
  ```
* 执行后会产生以下档案
    * 用于操作表格的程式码档案(.cs或是.go)
    * 符合proto格式的二进位档案(.byte)
    * 符合json格式的文字档案(.json)
* 关于程式码的范例可以看[范例档案](#范例档案)
    * cs范例在 example/example.cs
    * go范例在 example/example.go

# 范例档案
[example]

## 范例目录/档案说明

| 目录/档案          | 说明                                                     |
|:-------------------|:---------------------------------------------------------|
| example.xlsx       | 范例的excel档案, 其中@Data是资料表格, $Enum是列举表格    |
| example.yaml       | 范例的设定档案                                           |
| rebuild.bat        | 执行转换的批次档                                         |
| rebuild-format.bat | 执行转换的批次档(会格式化产生的程式码, 需要安装额外工具) |
| example.cs         | cs版本的表格操作程式码范例                               |
| example.go         | go版本的表格操作程式码范例                               |
| json\codeCs        | 产生的cs版本的json资料档案操作程式码                     |
| json\codeGo        | 产生的go版本的json资料档案操作程式码                     |
| json\data          | 产生的json资料档案                                       |
| proto\codeCs       | 产生的cs版本的proto资料档案操作程式码                    |
| proto\codeGo       | 产生的go版本的proto资料档案操作程式码                    |
| proto\data         | 产生的proto资料档案                                      |
| proto\schema       | 产生的proto结构档案                                      |
| enum\codeCs        | 产生的cs版本的列举程式码                                 |
| enum\codeGo        | 产生的go版本的列举程式码                                 |
| enum\schema        | 产生的列举结构档案                                       |
| template           | 模板档案                                                 |

# 命令说明
以下描述了[sheeter]提供的命令与旗标

## help命令
用于显示命令说明  
```sh
sheeter help [command]
```

## version命令
用于显示版本资讯  
```sh
sheeter version
```

## build命令
用于建置资料档案与程式码  
```sh
sheeter build [flags]
```
例如  
```sh
sheeter build --config setting.yaml
sheeter build --config setting.yaml --json --namespace
sheeter build --config setting.yaml --lineOfField 1 --lineOfLayer 2
```
请注意目前[sheeter]最多只能开启999999个档案(不过在开启这么多档案之前, 记忆体就用光了吧)

| 旗标          | 参数                                    | 说明                     |
|:--------------|:----------------------------------------|:-------------------------|
| --config      | 路径与档名; 例如: path/seeting.yaml     | 设定档案路径             |
| --json        |                                         | 是否产生json档案         |
| --proto       |                                         | 是否产生proto档案        |
| --enum        |                                         | 是否产生enum档案         |
| --namespace   |                                         | 是否用简单的命名空间名称 |
| --lineOfName  | 行号(1为起始行)                         | 名称行号                 |
| --lineOfNote  | 行号(1为起始行)                         | 注解行号                 |
| --lineOfField | 行号(1为起始行)                         | 栏位行号                 |
| --lineOfLayer | 行号(1为起始行)                         | 阶层行号                 |
| --lineOfData  | 行号(1为起始行)                         | 资料行号                 |
| --lineOfEnum  | 行号(1为起始行)                         | 列举行号                 |
| --excludes    | 标签,标签,...                           | 输出时排除的标签列表     |
| --inputs      | 路径,档案名称,档案名称#表单名称,...     | 输入列表                 |

### --config
从设定档读取参数, 设定档中的参数都可以被其他的旗标值替代  
```sh
sheeter build --config setting.yaml --lineOfName 5
```
像是这种情况, 设定档中的`lineOfName`的值就会被`--lineOfName`的值替代  

### --json / --proto
用于控制是否要产生[json]与[proto]档案  
* `sheeter build`  
  输出[json]与[proto]档案  
* `sheeter build --json --proto`  
  输出[json]与[proto]档案  
* `sheeter build --json`  
  只输出[json]档案  
* `sheeter build --proto`  
  只输出[proto]档案  
  
### --namespace
用于控制产生的命名空间名称  
* `sheeter build`  
  命名空间名称: sheeterJson / SheeterJson / sheeterProto / SheeterProto / sheeterEnum / SheeterEnum  
* `sheeter build --namespace`  
  命名空间名称: sheeter / Sheeter  
  
### --inputs
输入列表, 可用以下几种格式组合, 每个项目以`,`分隔; 注意程式只会读取副档名为xlsx的档案  
请注意路径中需使用`/`而非`\`
* 路径名称  
  path, path/, path/path...  
* 档案名称  
  example.xlsx, path/example.xlsx...  
* 档案名称+表单名称  
  example.xlsx#sheet, path/example.xlsx#sheet...  
  这个格式中, 需用`#`把档案名称与表单名称隔开  

## tmpl命令
用于产生执行时使用的模板档案, 你可以通过修改模板来改变产生出来的程式码  
```sh
sheeter tmpl [flags]
```

| 旗标         | 参数 | 说明             |
|:-------------|:-----|:-----------------|
| --clean / -c |      | 重新产生模板档案 |

# 资料表单说明
![exceldata]

## 表单名称
需以`@`开头, [sheeter]会自动辨识以`@`开头的表单为资料表单

## 名称行
栏位名称, 只能是英文与数字与`_`的组合, 并且不能以数字开头, 也不允许空白

## 注解行
单行注解, 若为空格就输出空注解  

## 栏位行
栏位类型与标签设置, 格式为`类型`或是`类型#标签`, 空格之后的栏位不会输出  
一个栏位只能设置一个标签  

| 类型        | 说明                                 |
|:------------|:-------------------------------------|
| empty       | 不会输出的栏位                       |
| pkey        | 表格主要索引, 编号可跳号但是不可重复 |
| bool        | 布林值                               |
| boolArray   | 以逗号分隔的布林值阵列               |
| int         | 64位元整数                           |
| intArray    | 以逗号分隔的64位元整数阵列           |
| float       | 64位元浮点数                         |
| floatArray  | 以逗号分隔的64位元整数阵列           |
| string      | 字串                                 |
| stringArray | 以逗号分隔的字串阵列                 |

## 栏位行范例

| 范例     | 栏位类型   | 标签 |
|:---------|:-----------|:-----|
| pkey     | pkey       |      |
| string   | string     |      |
| string#A | string     | A    |

## 标签与排除机制
栏位行可用标签来控制栏位与其资料是否要输出到资料档案  
当设定档中的排除标签与栏位的标签符合时, 该栏位就不会输出到资料档案  
栏位若是没有设定标签, 则永远会输出到资料档案  
当栏位的类型是`empty`, 则永远不会输出到资料档案  
标签与排除设置不会影响产生的程式码  
一个栏位只能设置一个标签

## 阶层行
栏位结构布局, 格式有`{名称`, `{[]名称`, `/`, `}`, 之间以空格分隔  

| 格式        | 说明                                |
|:------------|:------------------------------------|
| {结构名称   | 结构的开始                          |
| {[]阵列名称 | 阵列的开始                          |
| /           | 分隔阵列                            |
| }           | 结构/阵列结束, 可以连续结束, 如`}}` |

## 阶层行范例

| 范例          | 说明                                                 |
|:--------------|:-----------------------------------------------------|
| {Item         | 建立Item结构                                         |
| {[]Item       | 建立以Item结构为元素的阵列                           |
| {Reward {Item | 建立Reward结构, Item结构; Item结构是Reward结构的成员 |

## 资料行
依照类型填写相应的内容即可, 其中`empty`, `string`, `stringArray`这三种类型允许空格, 其他类型的空格会造成错误  
空表格(也就是没有任何资料行)是允许的  
转换时, 只会转换到第一个空行为止  

## 其他的限制
* 档案名称与表单名称
    * 不能是规定的[关键字](#关键字)
* 表格设置
    * 表格必须有名称行, 注解行, 栏位行, 阶层行, 但是可以不需要有资料行
    * 名称行, 注解行, 栏位行, 阶层行必须在资料行之前
    * 设定档中必须设定好名称行, 注解行, 栏位行, 阶层行, 资料行的位置
    * 设定档中行数是从1开始的
* 主索引
    * 表格必须有`pkey`栏位
    * 表格只能有一个`pkey`栏位
    * `pkey`栏位中的内容不能重复
* 阶层
    * 不属于结构/阵列的栏位名称不能重复(包括`empty`栏位)
    * 结构/阵列名称可以重复, 重复的结构/阵列的栏位会合并
    * 结构/阵列的栏位可以不必填上所有的名称
        * 第一个表格设定了结构/阵列栏位: `data { field1, field2, field3 }`
        * 另一个表格同样使用了`data`结构/阵列, 而栏位只设定 `data { field1, field2 }`, 忽略了`field3`

# 列举表单说明
![excelenum]

## 表单名称
需以`$`开头, [sheeter]会自动辨识以`$`开头的表单为列举表单

## 名称行
实际上不写也没关系, 仅提供给使用者辨识用

## 资料行
必须是以下栏位格式  
* 第一栏: 列举名称
    * 只能是英文与数字与`_`的组合, 并且不能以数字开头, 也不允许空白
    * 列举名称不允许重复
    * 不能是规定的[关键字](#关键字)
* 第二栏: 列举编号
    * 只能是数字
    * 索引编号不允许重复
* 第三栏: 列举注解
    * 单行注解, 此栏空白也可以

# 设定说明
```yaml
global:
  exportJson:      true # 是否产生json档案
  exportProto:     true # 是否产生proto档案
  exportEnum:      true # 是否产生enum档案
  simpleNamespace: true # 是否用简单的命名空间名称
  lineOfName:      1    # 名称行号(1为起始行)
  lineOfNote:      2    # 注解行号(1为起始行)
  lineOfField:     3    # 栏位行号(1为起始行)
  lineOfLayer:     4    # 阶层行号(1为起始行)
  lineOfData:      5    # 资料行号(1为起始行)
  lineOfEnum:      2    # 列举行号(1为起始行)
  excludes:             # 排除标签列表
    - tag1
    - tag2

inputs:                   # 输入列表
  - path1                 # 转换path1目录底下符合规格的excel档案
  - path2                 # 转换path2目录底下符合规格的excel档案
  - path/excel.xlsx       # 转换指定的excel档案内符合规格的表单
  - path/excel.xlsx#@Data # 转换指定的excel档案内的@Data表单
  - path/excel.xlsx#$Enum # 转换指定的excel档案内的$Enum表单
```

# 关键字
档案名称与表单名称合并之后的名称不能是以下名称的组合(不分大小写)  
* depot
* loader
* reader
* readers

# 产生目录

| 名称         | 说明               |
|:-------------|:-------------------|
| json         | json目录           |
| json/codeCs  | 存放产生的cs程式码 |
| json/codeGo  | 存放产生的go程式码 |
| json/data    | 存放资料档案       |
| proto        | proto目录          |
| proto/codeCs | 存放产生的cs程式码 |
| proto/codeGo | 存放产生的go程式码 |
| proto/data   | 存放资料档案       |
| proto/schema | 存放.proto档案     |
| enum         | enum目录           |
| enum/codeCs  | 存放列举程式码     |
| enum/codeGo  | 存放列举程式码     |
| enum/schema  | 存放.proto档案     |
| template     | 存放模板档案       |

# 模板档案
[sheeter]转换时会把使用的程式码模板输出到`template`目录下  
使用者可以改变模板内容, 来产生自订的程式码  
模板档案使用[go]的[template]语法, 同时可以参考以下模板参数来做名称的替换  

## json结构, 读取器模板参数
影响的档案: json-struct-cs/go, json-reader-cs/go  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域设定](#全域设定参数)                                 |
|                    |             | [命名工具](#命名工具参数)                                 |
|                    |             | [栏位命名工具](#栏位命名工具参数)                         |
|                    |             | [json命名工具](#json命名工具参数)                         |
| .Reader            |             | 是否要产生读取器                                          |
| .Fields            |             | [栏位列表](#栏位类型参数)                                 |

## json仓库模板参数
影响的档案: json-depot-cs/go  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域设定](#全域设定参数)                                 |
|                    |             | [命名工具](#命名工具参数)                                 |
|                    |             | [json命名工具](#json命名工具参数)                         |
| .Struct            |             | 结构列表(注)                                              |

结构列表内容为  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [命名工具](#命名工具参数)                                 |
| Reader             |             | 是否要产生读取器                                          |

## proto结构, 读取器, proto架构模板参数
影响的档案: proto-struct-cs/go, proto-reader-cs/go, proto-schema  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域设定](#全域设定参数)                                 |
|                    |             | [命名工具](#命名工具参数)                                 |
|                    |             | [栏位命名工具](#栏位命名工具参数)                         |
|                    |             | [proto命名工具](#proto命名工具参数)                       |
| .Reader            |             | 是否要产生读取器                                          |
| .Fields            |             | [栏位列表](#栏位类型参数)                                 |
| .Depend            |             | 依赖列表                                                  |

## proto仓库模板参数
影响的档案: proto-depot-cs/go  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域设定](#全域设定参数)                                 |
|                    |             | [命名工具](#命名工具参数)                                 |
|                    |             | [proto命名工具](#proto命名工具参数)                       |
| .Struct            |             | 结构列表(注)                                              |

结构列表内容为  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [命名工具](#命名工具参数)                                 |
| Reader             |             | 是否要产生读取器                                          |

## enum架构模板参数
影响的档案: enum-schema  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
|                    |             | [全域设定](#全域设定参数)                                 |
|                    |             | [命名工具](#命名工具参数)                                 |
|                    |             | [enum命名工具](#enum命名工具参数)                         |
| .Enums             |             | 列举列表(注)                                              |

结构列表内容为  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| Name               |             | 列举名称                                                  |
| Index              |             | 列举编号                                                  |
| Comment            |             | 列举说明                                                  |

## 全域设定参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ExportJson        |             | 是否产生json档案                                          |
| .ExportProto       |             | 是否产生proto档案                                         |
| .ExportEnum        |             | 是否产生enum档案                                          |
| .SimpleNamespace   |             | 是否用简单的命名空间名称                                  |
| .LineOfName        |             | 名称行号(1为起始行)                                       |
| .LineOfNote        |             | 注解行号(1为起始行)                                       |
| .LineOfField       |             | 栏位行号(1为起始行)                                       |
| .LineOfLayer       |             | 阶层行号(1为起始行)                                       |
| .LineOfData        |             | 资料行号(1为起始行)                                       |
| .LineOfEnum        |             | 列举行号(1为起始行)                                       |
| .Excludes          |             | 排除标签列表                                              |

## 命名工具参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ExcelName         |             | excel名称                                                 |
| .SheetName         |             | sheet名称                                                 |
| .AppName           |             | 程式名称                                                  |
| .JsonNamespace     | bool        | json命名空间名称, 参数影响是否用简单的命名空间名称        |
| .ProtoNamespace    | bool        | proto命名空间名称, 参数影响是否用简单的命名空间名称       |
| .EnumNamespace     | bool        | enum命名空间名称, 参数影响是否用简单的命名空间名称        |
| .StructName        |             | 结构名称                                                  |
| .ReaderName        |             | 读取器名称                                                |
| .StorerName        |             | 储存器名称                                                |
| .StorerDatas       |             | 储存器资料名称                                            |
| .StorerMessage     | bool        | 储存器proto message名称, 参数影响是否用简单的命名空间名称 |
| .FirstUpper        | 字串        | 字串首字母大写                                            |
| .FirstLower        | 字串        | 字串首字母小写                                            |
| .Add               | 数值1 数值2 | 加法(数值1 + 数值2)                                       |
| .Sub               | 数值1 数值2 | 减法(数值1 - 数值2)                                       |
| .Mul               | 数值1 数值2 | 乘法(数值1 x 数值2)                                       |
| .Div               | 数值1 数值2 | 除法(数值1 / 数值2)                                       |

## 栏位命名工具参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .FieldName         | 栏位资料    | 栏位名称                                                  |
| .FieldNote         | 栏位资料    | 栏位注解                                                  |
| .FieldTypeCs       | 栏位资料    | cs栏位类型                                                |
| .FieldTypeGo       | 栏位资料    | go栏位类型                                                |
| .FieldTypeProto    | 栏位资料    | proto栏位类型                                             |
| .PkeyTypeCs        |             | pkey的cs类型                                              |
| .PkeyTypeGo        |             | pkey的go类型                                              |
| .PkeyTypeProto     |             | pkey的proto类型                                           |

## json命名工具参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .JsonDataName      |             | json资料名称                                              |
| .JsonDataExt       |             | json资料副档名                                            |
| .JsonDataFile      |             | json资料档名                                              |
| .JsonDataPath      |             | json资料路径                                              |
| .JsonStructCsPath  |             | json-cs结构程式码路径                                     |
| .JsonReaderCsPath  |             | json-cs读取器程式码路径                                   |
| .JsonDepotCsPath   |             | json-cs仓库程式码路径                                     |
| .JsonStructGoPath  |             | json-go结构程式码路径                                     |
| .JsonReaderGoPath  |             | json-go读取器程式码档名路径                               |
| .JsonDepotGoPath   |             | json-go仓库程式码路径                                     |

## proto命名工具参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .ProtoCsPath       |             | proto-cs路径                                              |
| .ProtoGoPath       |             | proto-go路径                                              |
| .ProtoSchemaPath   |             | proto-schema路径                                          |
| .ProtoName         |             | proto架构档名                                             |
| .ProtoPath         |             | proto架构路径                                             |
| .ProtoDataName     |             | proto资料名称                                             |
| .ProtoDataExt      |             | proto资料副档名                                           |
| .ProtoDataFile     |             | proto资料档名                                             |
| .ProtoDataPath     |             | proto资料路径                                             |
| .ProtoReaderCsPath |             | proto-cs读取器程式码路径                                  |
| .ProtoDepotCsPath  |             | proto-cs仓库程式码路径                                    |
| .ProtoReaderGoPath |             | proto-go读取器程式码路径                                  |
| .ProtoDepotGoPath  |             | proto-go仓库程式码路径                                    |
| .ProtoDepend       | 依赖名称    | proto依赖档案名称                                         |

## enum命名工具参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .EnumCsPath        |             | enum-cs路径                                               |
| .EnumGoPath        |             | enum-go路径                                               |
| .EnumSchemaPath    |             | enum-schema路径                                           |
| .EnumName          |             | enum架构档名                                              |
| .EnumPath          |             | enum架构路径                                              |

## 栏位类型参数

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .Name              |             | 栏位名称                                                  |
| .Note              |             | 栏位注解                                                  |
| .Field             |             | 栏位类型(注)                                              |
| .Alter             |             | 栏位类型别名                                              |
| .Array             |             | 阵列旗标                                                  |

栏位类型内容为  

| 名称               | 参数        | 说明                                                      |
|:-------------------|:------------|:----------------------------------------------------------|
| .Type              |             | excel栏位类型                                             |
| .IsShow            |             | 是否显示栏位                                              |
| .IsPkey            |             | 是否是主要索引                                            |
| .ToTypeCs          |             | cs类型字串                                                |
| .ToTypeGo          |             | go类型字串                                                |
| .ToTypeProto       |             | proto类型字串                                             |

# 格式化程式码
[sheeter]并不负责帮产生的档案排版, 如果需要排版, 就需要自己写.bat/.sh来执行  
以下介绍cs, go, proto的排版工具, 有需要可以自己去安装  

## csharpier
用于cs的排版工具  
* 安装
    * 安装[dotnet], 如果有安装.net sdk, 或有安装unity可能可以省略此步骤
    * 安装[csharpier], 在终端执行以下命令
      ```sh
      dotnet tool install csharpier -g
      ```
* 使用
    * 在终端执行以下命令
      ```sh
      dotnet csharpier .
      ```

## gofmt
用于go的排版工具  
* 安装
    * 安装[go]时会顺便安装
* 使用
    * 在终端执行以下命令
      ```sh
      gofmt -w .
      ```

## buf
用于proto的排版工具  
* 安装
    * 安装[buf], 在终端执行以下命令
      ```sh
      go install github.com/bufbuild/buf/cmd/buf@v1.8.0
      ```
* 使用
    * 在终端执行以下命令
      ```sh
      buf format -w .
      ```

# 专案目录说明

| 目录                    | 说明                             |
|:------------------------|:---------------------------------|
| cmd/sheeter             | sheeter命令程式                  |
| cmd/sheeter/build       | 建置表格命令                     |
| cmd/sheeter/tmpl        | 产生模板命令                     |
| cmd/sheeter/version     | 显示版本命令                     |
| sheeter                 | sheeter命令程式用到的各项组件    |
| sheeter/builds          | 表格转换(用于build命令)          |
| sheeter/excels          | 表格组件                         |
| sheeter/fields          | 栏位组件                         |
| sheeter/layers          | 阶层组件                         |
| sheeter/layouts         | 布局组件                         |
| sheeter/nameds          | 命名工具                         |
| sheeter/pipelines       | 管线组件                         |
| sheeter/tmpls           | 模板组件(用于templ及build命令)   |
| sheeter/utils           | 协助组件                         |
| support                 | 支援专案                         |
| support/benchmark_count | 档案数量效率测试资料             |
| support/benchmark_size  | 档案大小效率测试资料             |
| support/example         | 范例资料                         |
| support/handmade        | 手制模板, 用来检查模板是否有错误 |
| support/handmade/.json  | json手制模板                     |
| support/handmade/.proto | proto手制模板                    |
| support/verifycs        | cs程式码验证                     |
| support/verifygo        | go程式码验证                     |
| support/verifyunity     | unity程式码验证                  |
| testdata                | 测试资料                         |

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