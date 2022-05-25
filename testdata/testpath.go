package testdata

import (
	"os"
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
	_, file, _, ok := runtime.Caller(0)

	if ok == false {
		panic("can't get root")
	}

	rootPath = filepath.Dir(file)
}

// Path 取得測試資料路徑
func Path(path string) string {
	return filepath.Join(rootPath, path)
}

// ChangeWorkDir 變更工作目錄到測試目錄
func ChangeWorkDir() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	} // if

	err = os.Chdir(rootPath)

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

var rootPath string // 根路徑
