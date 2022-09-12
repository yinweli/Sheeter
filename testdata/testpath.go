package testdata

import (
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, file, _, ok := runtime.Caller(0)

	if ok == false {
		panic("get testdata rootPath failed")
	} // if

	RootPath = filepath.Dir(file)
}

// Path 取得測試路徑
func Path(path string) string {
	return filepath.Join(RootPath, path)
}

// ChangeWorkDir 變更工作目錄到測試目錄
func ChangeWorkDir() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	} // if

	if err = os.Chdir(RootPath); err != nil {
		panic(err)
	} // if

	return dir
}

// RestoreWorkDir 復原工作目錄
func RestoreWorkDir(dir string) {
	if err := os.Chdir(dir); err != nil {
		panic(err)
	} // if
}

var RootPath string // 根路徑
