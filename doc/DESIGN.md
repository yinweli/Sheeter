# 執行起點
`Sheeter`的執行起點(也就是main函式所在地)在`cmd/sheeter/sheeter.go`  

# 命令流程
各項命令的執行流程如下  

## build
建置命令, 將excel轉換成程式碼與資料檔案  
`build`會執行以下程序  
* 模板初始化
* 讀取設定
* 建置初始化
* 產生程式碼
* 產生資料檔案
* 後製處理
> 程式碼位於  
> `cmd/sheeter/build/...`  

## tmpl
模板命令, 進行模板檔案的清理或初始化  
`tmpl`會執行以下程序  
* 模板初始化
> 程式碼位於  
> `cmd/sheeter/tmpl/...`  

## version
輸出版本字串  
> 程式碼位於  
> `cmd/sheeter/version/...`  

# 流程細節
以下說明流程的細節  

## 模板初始化
模板是產生程式碼時使用的模板字串  
模板初始化會執行以下程序  
* 執行清除模板檔案(如果有設定清理旗標的話)
* 讀取模板檔案(如果沒有模板檔案的話, 就會使用預設模板)
* 儲存模板檔案
> 程式碼位於  
> `sheeter/tmpls/...`  

## 讀取設定
設定有兩個來源`設定檔案`與`命令列旗標`, 優先順序是`命令列旗標`大於`設定檔案`  
讀取設定會執行以下程序  
* 讀取設定
* 檢查設定是否正確
> 程式碼位於  
> `sheeter/builds/config.go`  
> `sheeter/builds/flag.go`  

## 建置初始化
進行建置前的準備工作, 以及從excel檔案中讀取資料等  
建置初始化會執行以下程序  
* 獲取可供轉換的excel檔案列表
* 開啟excel檔案, 獲取表格列表
* 讀取資料表格 / 列舉表格
* 建立產生程式碼 / 產生資料檔案 / 後製處理的指令列表
> 程式碼位於  
> `sheeter/builds/context.go`  
> `sheeter/builds/initialize.go`  
> `sheeter/builds/initializeFile.go`  
> `sheeter/builds/initializeExcel.go`  
> `sheeter/builds/initializeSheetData.go`  
> `sheeter/builds/initializeSheetEnum.go`  
> `sheeter/builds/InitializePick.go`  

## 產生程式碼
使用模板來產生產生程式碼  
產生程式碼會執行以下程序  
* 產生json的結構程式碼(cs)
* 產生json的讀取器程式碼(cs)
* 產生json的結構程式碼(go)
* 產生json的讀取器程式碼(go)
* 產生proto的架構檔案
* 產生proto的讀取器程式碼(cs)
* 產生proto的讀取器程式碼(go)
* 產生列舉的架構檔案
> 程式碼位於  
> `sheeter/builds/context.go`  
> `sheeter/builds/generate.go`  
> `sheeter/builds/generateJson.go`  
> `sheeter/builds/generateProto.go`  
> `sheeter/builds/generateEnum.go`  

## 產生資料檔案
產生`.json`以及`.byte`的資料檔案  
產生資料檔案會執行以下程序  
* 產生json資料檔案
* 產生proto資料檔案
> 程式碼位於  
> `sheeter/builds/context.go`  
> `sheeter/builds/encoding.go`  
> `sheeter/builds/encodingJson.go`  
> `sheeter/builds/encodingProto.go`  

## 後製處理
進行無法在前面的步驟中執行的指令  
後製處理會執行以下程序  
* 產生json的倉庫程式碼(cs)
* 產生json的倉庫程式碼(go)
* 產生proto的倉庫程式碼(cs)
* 產生proto的倉庫程式碼(go)
* 從proto的架構檔案產生程式碼(cs)
* 從proto的架構檔案產生程式碼(go)
> 程式碼位於  
> `sheeter/builds/context.go`  
> `sheeter/builds/poststep.go`  
> `sheeter/builds/poststepJsonDepot.go`  
> `sheeter/builds/poststepProtoDepot.go`  
> `sheeter/builds/poststepConvert.go`  

# 協力組件
流程執行時會用到下列組件  

## excels
從excel檔案中讀取資料的組件  
> 程式碼位於  
> `sheeter/excels/...`  

## fields
定義`Sheeter`支援的資料類型以及其相關設定, 若想要新增資料類型, 就得到此修改程式碼  
負責從表單的欄位行解析欄位資料, 提供給後續步驟使用  
> 程式碼位於  
> `sheeter/fields/...`  

## layers
負責從表單的階層行解析階層資料, 提供給後續步驟使用  
> 程式碼位於  
> `sheeter/layers/...`  

## layouts
負責儲存表單的各項資料, 以及相關工具, 詳列如下  
* layoutType  
  負責儲存表單的類型資料  
  > 程式碼位於  
  > `sheeter/layouts/layoutType.go`  
* layoutData  
  負責儲存表單的內容資料  
  > 程式碼位於  
  > `sheeter/layouts/layoutData.go`  
* layoutDepend  
  負責儲存依賴關係, 產生proto的架構檔案時會用到  
  > 程式碼位於  
  > `sheeter/layouts/layoutDepend.go`  
* layoutEnum  
  負責儲存表單的列舉資料  
  > 程式碼位於  
  > `sheeter/layouts/layoutEnum.go`  
* jsonPack  
  打包json資料; 將會把excel中的資料, 依據資料布局與排除標籤, 轉換為json格式的位元陣列  
  > 程式碼位於  
  > `sheeter/layouts/jsonPack.go`  
* structor  
  結構化儲存容器, 提供給layoutData使用  
  > 程式碼位於  
  > `sheeter/layouts/structor.go`  

## nameds
負責在產生程式碼時, 替代模板中的變數; 或是提供組合檔案名稱 / 檔案路徑等功能  
> 程式碼位於  
> `sheeter/nameds/...`  

## pipelines
負責有多項工作需要執行時, 安排多執行緒事項與進度條顯示的功能  
> 程式碼位於  
> `sheeter/pipelines/...`  

## tmpls
負責模板相關功能  
> 程式碼位於  
> `sheeter/tmpls/...`  

## utils
提供雜項工具  
> 程式碼位於  
> `sheeter/utils/...`  