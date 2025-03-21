# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Planning]
- 針對單元測試新增goleak來測試是否有gorourtine洩漏

## [Unrelease]

## [2.3.13] - 2025-03-10
### Fixed
- 修正排序

## [2.3.12] - 2025-03-10
### Changed
- 改善 sheeter.cs / sheeter.go 程式碼中的表單物件排序方式為字典順序

## [2.3.11] - 2024-08-14
### Fixed
- 修正錯誤輸出

## [2.3.10] - 2024-08-14
### Changed
- build action改為等待lint, test action完成後才執行
### Deleted
- 移除test-race action

## [2.3.9] - 2024-08-14
### Fixed
- 修正版本問題

## [2.3.8] - 2024-08-14
### Changed
- 標準輸出物件增加顏色功能

## [2.3.7] - 2024-08-14
### Changed
- 錯誤訊息改顯示到錯誤輸出

## [2.3.6] - 2024-05-24
### Fixed
- 優化多表合一語法
### Changed
- 更新格式化方式

## [2.3.5] - 2024-05-24
### Fixed
- 優化多表合一語法

## [2.3.4] - 2024-05-07
### Fixed
- 修正多表合一語法錯誤

## [2.3.3] - 2024-05-07
### Fixed
- 修正cs檔案語法錯誤

## [2.3.2] - 2024-04-15
### Added
- 現在會自動產生安裝包

## [2.3.1] - 2024-01-10
### Added
- 新增重複索引檢查

## [2.3.0] - 2023-12-11
### Added
- 新增多表合一功能

## [2.2.0] - 2023-08-10
### Added
- go版本的讀取程式碼新增取得全部內容的any版本
### Changed
- 變更結構說明字串

## [2.1.5] - 2023-05-09
### Fixed
- 修正版本編號錯誤

## [2.1.3] - 2023-05-08
### Changed
- 添加錯誤訊息

## [2.1.2] - 2023-03-24
### Changed
- 整理測試機制

## [2.1.1] - 2023-03-17
### Changed
- 排序輸出程式碼中的表格列表與欄位列表

## [2.1.0] - 2023-03-16
### Added
- 新增排除列表

## [2.0.1] - 2023-03-16
### Changed
- 新增輸出路徑功能
- 整理輸出檔名規則

## [2.0.0] - 2023-03-16
### Removed
- 移除proto支援
- 移除enum功能
- 移除階層功能
### Changed
- 變更sheeter支援的excel格式, 因此v2之後的sheeter不再能讀取v1的excel檔案

## [1.11.0] - 2022-12-02
### Changed
- 模板中新增FileName類別, 用來取得檔名, 副檔名等
- 更改Reader介面, 把DataName, DataExt, DataFile合併成FileName函式, 並統一回傳FileName物件
- 更改Depot.Loader介面, 改用FileName物件來做處理
- 把標籤從欄位行獨立出來, 標籤改成單字元, 並且改成正相關方式判斷是否要輸出
### Fixed
- 修正測試環境機制在env資料夾未建立的狀況下會拋出錯誤

## [1.10.2] - 2022-11-04
### Added
- 新增測試環境機制
- 新增build指令的單元測試
- 新增tmpl指令的單元測試
- 新增version指令的單元測試
### Changed
- 改變單元測試的工作路徑與測試資料機制, 改用測試環境機制
- 設置最大可以開啟的excel數量為999999個

## [1.10.1] - 2022-11-01
### Added
- 新增多重輸入驗證專案
### Fixed
- 修正重複的路徑/excel/sheet會被重複執行的問題

## [1.10.0] - 2022-10-31
### Added
- build命令可從指定目錄中搜尋excel檔案, 並辨別帶有"@"符號或是"$"符號的表單來建置成資料檔案或是列舉檔案
### Changed
- 設定檔中的element與enum區塊被inputs區塊代替
- 產生出來的讀取器程式碼中的DataExt函式, 回傳的副檔名字串會在最前面添加"."符號

## [1.9.4] - 2022-10-28
### Changed
- 更新範例檔案, 加入關於多執行緒相關的說明
- 驗證程式加入多執行緒驗證
### Fixed
- 修正sheet組件的next會造成無法判斷是否到達最後一行的問題

## [1.9.3] - 2022-10-25
### Changed
- text, testArray類型更名為string, stringArray
- 更新範例檔案

## [1.9.2] - 2022-10-24
### Changed
- 精簡程式碼, 並增加組件說明
- 更新範例檔案

## [1.9.1] - 2022-10-21
### Changed
- 簡化layoutType機制
- 變更進度條顯示文字

## [1.9.0] - 2022-10-20
### Added
- 新增從excel產生列舉
### Removed
- 不再產生proto**.bat/.sh
### Changed
- 內部流程改用管線機制重構
- 變更模板檔案名稱
- 後製時會執行protoc產生程式碼
- 變更範例專案
- 產生的cs程式碼一律為大寫開頭

## [1.8.1] - 2022-10-14
### Changed
- 欄位名稱從欄位行中獨立成名稱行
### Fixed
- 修正初始化遭遇錯誤不會正確顯示錯誤訊息的問題

## [1.7.2] - 2022-10-13
### Changed
- excel函式庫改用xlsxreader, 對於轉換大量excel檔案或是大型excel檔案時效能增加

## [1.7.0] - 2022-10-12
### Added
- 新增clear功能, 可以在執行期清空表格資料

## [1.6.0] - 2022-10-11
### Added
- 新增產生倉庫程式碼
### Changed
- 改進讀取器介面
- 改進json儲存器
- 改進模板參數機制

## [1.5.2] - 2022-10-07
### Changed
- 簡化讀取器
- 調整支援專案目錄結構

## [1.5.1] - 2022-10-06
### Changed
- 修改proto資料檔案副檔名為bytes

## [1.5.0] - 2022-10-06
### Added
- 新增設定, 可選擇是否是用簡單命名空間名稱

## [1.4.0] - 2022-10-05
### Added
- 新增設定, 可選擇要輸出json或是proto, 或是兩者皆有

## [1.3.2] - 2022-10-05
### Fixed
- 把experimental_allow_proto3_optional旗標還原回來

## [1.3.0] - 2022-10-05
### Added
- 新增merge系列函式

## [1.2.1] - 2022-10-03
### Fixed
- 移除quicktype安裝檢查

## [1.2.0] - 2022-09-30
### Added
- 新增排除標籤功能

## [1.1.1] - 2022-09-30
### Fixed
- 修正json-cs, json-go的結構模板加入是否要產生storer的判斷

## [1.1.0] - 2022-09-29
### Removed
- 刪除讀取器的FromPathFull函式
### Changed
- 讀取器的FromPathHalf改名為FromPath

## [1.0.1] - 2022-09-29
### Fixed
- 修正產生的proto-cs讀取器檔名, 使其與proto產生的結構檔名格式一致

## [1.0.0] - 2022-09-29
### Added
- 增加輸出proto檔案, proto資料功能

## [0.3.8] - 2022-09-20
### Fixed
- 修正go的讀取器模板程式碼名稱衝突錯誤

## [0.3.7] - 2022-09-17
### Added
- 新增模板機制, 使用者可以按自己需要寫模板檔案, 變更產生的程式碼內容
### Changed
- 不再使用quicktype來產生程式碼, 改用模板來產生

## [0.3.6] - 2022-09-12
### Added
- 完成cs程式碼驗證
- 完成go程式碼驗證
### Removed
- 移除bom支援

## [0.3.5] - 2022-09-12
### Added
- code指令新增清理旗標

## [0.3.4] - 2022-09-11
### Added
- 新增產生程式碼的驗證程序
### Fixed
- 修正go模板的語法錯誤
- 修正json檔案路徑錯誤

## [0.3.3] - 2022-09-11
### Added
- 新增code命令
### Fixed
- 修正命令旗標錯誤

## [0.3.2] - 2022-09-07
### Changed
- 變更產出的程式碼為單一檔案, 讓重複的結構可以被程式碼重複使用

## [0.3.1] - 2022-08-27
### Changed
- 重新安排輸出路徑與命名空間, 避免相同結構名稱問題

## [0.3.0] - 2022-08-26
### Removed
- 移除lua支援
### Changed
- 改造成結構化表格轉換