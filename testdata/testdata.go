package testdata

import (
	"path/filepath"
	"runtime"
)

// BoolString 取得布林值字串
func BoolString() string {
	return "true,false,false,true,true,false,true"
}

// BoolArray 取得布林值陣列
func BoolArray() []bool {
	return []bool{true, false, false, true, true, false, true}
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

// Path 取得測試路徑
func Path(path string) string {
	return filepath.Join(root, path)
}

func init() {
	_, currentFile, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get root")
	}

	root = filepath.Dir(currentFile)
}

var root string // 根路徑
