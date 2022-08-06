package thirdParty

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// Println 輸出字串函式類型, 並且輸出完要換行
type Println func(i ...interface{})

// Check 檢查是否安裝第三方軟體
func Check(println Println) bool {
	result := true
	result = result && check(println, "go")
	result = result && check(println, "quicktype")
	return result
}

// check 檢查是否安裝第三方軟體
func check(println Println, name string) bool {
	if util.ShellExist(name) == false {
		println(fmt.Sprintf("%s not installed", name))
		return false
	} // if

	return true
}
