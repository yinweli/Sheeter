# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Planning]
- 嘗試看看用excel, sheet名稱以及資料夾搜尋(包括子資料夾)來建立表格資料; 取代寫一堆的設定檔案
```
build --config xxxx
build --element xxxx\xxxx#xxxx
build --enum xxxx\xxxx#xxxx

build --config xxxx
build --input xxxx,xxxx\xxxx.xlsx,xxxx.xlsx#sheet

如何分辨
    非excel檔案
    跳過檔案/表單
    資料表單 @
    列舉表單 $

執行流程
    從設定檔或是命令行獲取path list, excel list, sheet list
    initializeInput {
        in:  result chan any
             path string
        out: path&file list
    }
    initializeExcel {
        in:  result chan any
             path&file
        out: sheetInfo
    }
    initializeSheet { 
        struct sheetInfo {
            excel path&file
            sheet name
            excel object
        }
        in:  result chan any
             sheetInfo
        out: data or enum object
    }
    initializeData => 填滿data結構中其餘項目
    initializeEnum => 填滿enum結構中其餘項目
    initializePick => pick!

單元測試流程
    flag_test
    config_test

技術資訊
    遍歷目錄與檔案使用filepath.Walk
    err := filepath.Walk(path, 處理函式)
    ////////////////////////////////////////////
    path := "./path/to/fileOrDir"
    fileInfo, err := os.Stat(path)
    if err != nil {
        // error handling
    }

    if fileInfo.IsDir() {
        // is a directory
    } else {
        // is not a directory
    }
```
- 產生flatbuffer

## [1.10.0]
### Changed
- 產生出來的讀取器程式碼中的DataExt函式, 回傳的副檔名會在最前面添加"."符號
- 設定檔中的element與enum區塊被inputs區塊代替
### Added
- build命令可從指定目錄中搜尋excel檔案, 並辨別帶有"@"符號或是"$"符號的表單來建置成資料檔案或是列舉檔案

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
### Changed
- 內部流程改用管線機制重構
- 變更模板檔案名稱
- 後製時會執行protoc產生程式碼
- 變更範例專案
- 產生的cs程式碼一律為大寫開頭
### Added
- 新增從excel產生列舉
### Removed
- 不再產生proto**.bat/.sh

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
### Changed
- 改造成結構化表格轉換
### Removed
- 移除lua支援