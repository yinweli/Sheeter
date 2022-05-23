package testdata

import (
	"path/filepath"
	"runtime"
)

// Path 取得測試資料路徑
func Path(path string) string {
	return filepath.Join(root, path)
}

// RealYaml 取得real.yaml路徑
func RealYaml() string {
	return Path("real.yaml")
}

// FakeYaml 取得fake.yaml路徑
func FakeYaml() string {
	return Path("fake.yaml")
}

// UnknownYaml 取得????.yaml路徑
func UnknownYaml() string {
	return Path("????.yaml")
}

func init() {
	_, currentFile, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get root")
	}

	root = filepath.Dir(currentFile)
}

var root string // 測試資料根路徑
