package testdata

import (
	"path/filepath"
	"runtime"
)

const RealConfig = "real.yaml"
const FakeConfig = "fake.yaml"
const DefectConfig = "defect.yaml"
const UnknownConfig = "????.yaml"
const TestExcel = "test.xlsx"
const TestSheet = "Data"

func init() {
	_, currentFile, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get root")
	}

	rootPath = filepath.Dir(currentFile)
}

// Path 取得測試資料路徑
func Path(path string) string {
	return filepath.Join(rootPath, path)
}

var rootPath string // 根路徑
