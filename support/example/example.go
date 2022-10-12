package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	sheeterJson "github.com/yinweli/Sheeter/support/example/json/codeGo"
	sheeterProto "github.com/yinweli/Sheeter/support/example/proto/codeGo"
)

func main() {
	exampleJson()
	exampleProto()
}

// exampleJson json範例
func exampleJson() {
	// 要使用sheeter, 首先建立繼承自sheeterJson.Loader介面的讀取器
	// 由於要設定資料檔案的路徑, 所以執行rootPath函式來取得當前路徑, 並且跟json資料的路徑組合來獲得最終路徑
	loader := &fileLoader{
		path: filepath.Join(rootPath(), "json", "data"),
	}
	// 接著建立sheeterJson.Depot物件, 這是存取表格資料最主要的物件
	// 要記得把剛剛建立的讀取器設定進去
	depot := sheeterJson.NewDepot(loader)

	// 然後執行FromData(或是MergeData)函式來讀取表格資料
	if depot.FromData() == false {
		fmt.Println("json failed: from data failed")
	}

	// 之後就可以用Depot底下的各個表格物件來取用資料內容
	if data, ok := depot.ExampleData.Get(1); ok {
		fmt.Printf("%v\n", data)
		fmt.Println("json success")
	} else {
		fmt.Println("json failed: get data failed")
	}
}

// exampleProto proto範例
func exampleProto() {
	// 要使用sheeter, 首先建立繼承自sheeterProto.Loader介面的讀取器
	// 由於要設定資料檔案的路徑, 所以執行rootPath函式來取得當前路徑, 並且跟proto資料的路徑組合來獲得最終路徑
	loader := &fileLoader{
		path: filepath.Join(rootPath(), "proto", "data"),
	}
	// 接著建立sheeterProto.Depot物件, 這是存取表格資料最主要的物件
	// 要記得把剛剛建立的讀取器設定進去
	depot := sheeterProto.NewDepot(loader)

	// 然後執行FromData(或是MergeData)函式來讀取表格資料
	if depot.FromData() == false {
		fmt.Println("proto failed: from data failed")
	}

	// 之後就可以用Depot底下的各個表格物件來取用資料內容
	if data, ok := depot.ExampleData.Get(1); ok {
		fmt.Printf("%v\n", data)
		fmt.Println("proto success")
	} else {
		fmt.Println("proto failed: get data failed")
	}
}

// fileLoader 檔案讀取器
type fileLoader struct {
	path string
}

// Error 用於處理讀取資料錯誤, 範例中只是單純印出錯誤訊息
func (this *fileLoader) Error(name string, err error) {
	fmt.Println(fmt.Errorf("%s: file load failed: %w", name, err))
}

// Load 用於讀取資料檔案, Depot會提供給你檔案名稱(name), 副檔名(ext), 完整名稱(fullname)
// 使用者需要依靠以上資訊來讀取資料檔案, 並回傳資料給Depot
func (this *fileLoader) Load(name, ext, fullname string) []byte {
	path := filepath.Join(this.path, fullname)
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(fmt.Errorf("%s: file load failed: %w", name, err))
		return nil
	}

	return data
}

// rootPath 取得根目錄路徑
func rootPath() string {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("root path failed"))
	} // if

	return filepath.Dir(root)
}
