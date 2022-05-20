package testdata

import (
	"path/filepath"
	"runtime"
)

// Path 取得測試資料路徑
func Path(path string) (result string) {
	return filepath.Join(testdataPath, path)
}

func init() {
	_, currentFile, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get testdata path")
	}

	testdataPath = filepath.Dir(currentFile)
}

var testdataPath string // 測試資料路徑
