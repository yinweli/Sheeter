package testdata

import (
	"path/filepath"
	"runtime"
)

// BoolString 取得bool字串
func BoolString() string {
	return "true,false,false,true,true,false,true"
}

// BoolArray 取得bool陣列
func BoolArray() []bool {
	return []bool{true, false, false, true, true, false, true}
}

// Float32String 取得float32字串
func Float32String() string {
	return "0.101,0.202,0.303,0.404,0.505"
}

// Float32Array 取得float32陣列
func Float32Array() []float32 {
	return []float32{0.101, 0.202, 0.303, 0.404, 0.505}
}

// Float64String 取得float64字串
func Float64String() string {
	return "0.000101,0.000202,0.000303,0.000404,0.000505"
}

// Float64Array 取得float64陣列
func Float64Array() []float64 {
	return []float64{0.000101, 0.000202, 0.000303, 0.000404, 0.000505}
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
