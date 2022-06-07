package testdata

import (
	"os"
	"path/filepath"
	"runtime"
)

/***** 測試路徑 *****/

// Path 取得測試路徑
func Path(path string) string {
	return filepath.Join(RootPath, path)
}

var RootPath string // 根路徑

/***** 變更工作目錄 *****/

// ChangeWorkDir 變更工作目錄到測試目錄
func ChangeWorkDir() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	} // if

	err = os.Chdir(RootPath)

	if err != nil {
		panic(err)
	} // if

	return dir
}

// RestoreWorkDir 復原工作目錄
func RestoreWorkDir(dir string) {
	err := os.Chdir(dir)

	if err != nil {
		panic(err)
	} // if
}

/***** 初始執行 *****/

func init() {
	_, file, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get root")
	}

	RootPath = filepath.Dir(file)
}
