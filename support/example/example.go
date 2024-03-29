package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/yinweli/Sheeter/v2/support/example/codeGo"
)

func main() {
	usecase()
}

// 在多執行緒環境下執行時, 各個表格物件的取用資料內容操作都是執行緒安全的
// 但是FromData與MergeData操作則是非執行緒安全的, 請注意此點

// usecase 使用範例
func usecase() {
	// 要使用sheeter, 首先建立繼承自Sheeter.Loader介面的裝載器
	// 裝載器負責從磁碟(或是其他的資料來源)讀取資料, 這部分由使用者自行處理
	// 範例中的裝載器只是簡單的從磁碟讀取檔案而已
	loader := &fileLoader{
		path: rootPath(),
	}
	// 接著建立Sheeter.Sheeter物件, 這是存取表格資料最主要的物件
	// 要記得把剛剛建立的裝載器設定進去
	sheet := sheeter.NewSheeter(loader)

	// 然後執行FromData(或是MergeData)函式來讀取表格資料
	if sheet.FromData() == false {
		fmt.Println("load failed: from data failed")
	} // if

	// 之後就可以用Sheeter.Sheeter底下的各個表格物件來取用資料內容
	if data := sheet.VerifyData.Get(1); data != nil {
		fmt.Printf("%v\n", data)
		fmt.Println("get data success")
	} else {
		fmt.Println("get data failed")
	} // if

	if data := sheet.VerifyData.Get(2); data != nil {
		fmt.Printf("%v\n", data)
		fmt.Println("get data success")
	} else {
		fmt.Println("get data failed")
	} // if
}

// fileLoader 裝載器
type fileLoader struct {
	path string
}

// Load 用於讀取資料檔案, sheeter提供給你FileName物件, 使用者依靠FileName的功能取得檔名來讀取資料, 並回傳檔案內容給sheeter
func (this *fileLoader) Load(filename sheeter.FileName) []byte {
	path := filepath.Join(this.path, "json", filename.File()) // 資料檔案放在json目錄下
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(fmt.Errorf("file load failed: %v: %w", filename.Name(), err))
		return nil
	} // if

	return data
}

// Error 用於處理錯誤, 範例中只是單純印出錯誤訊息
func (this *fileLoader) Error(name string, err error) {
	fmt.Println(fmt.Errorf("file load failed: %v: %w", name, err))
}

// rootPath 取得根目錄路徑
func rootPath() string {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("root path failed"))
	} // if

	return filepath.Dir(root)
}
