package testdata

import (
	"path/filepath"
	"runtime"
)

// Path 取得測試資料路徑
func Path(path string) string {
	return filepath.Join(testdataPath, path)
}

// RealYaml 取得real.yaml
func RealYaml() string {
	return Path("real.yaml")
}

// FakeYaml 取得fake.yaml
func FakeYaml() string {
	return Path("fake.yaml")
}

func init() {
	_, currentFile, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get testdataPath")
	}

	testdataPath = filepath.Dir(currentFile)
}

var testdataPath string // 測試資料路徑
