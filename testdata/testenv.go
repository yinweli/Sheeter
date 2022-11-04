package testdata

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func init() {
	_, file, _, ok := runtime.Caller(0)

	if ok == false {
		panic("get rootpath failed")
	} // if

	rootpath = filepath.Dir(file)
	envpath = filepath.Join(rootpath, "env")
}

// TestEnv 測試環境
type TestEnv struct {
	original string // 原始路徑
	workpath string // 工作路徑
}

// Change 變更工作目錄
func (this *TestEnv) Change(dir string) {
	original, err := os.Getwd()

	if err != nil {
		panic(err)
	} // if

	workpath := filepath.Join(rootpath, dir)

	if err = os.MkdirAll(workpath, os.ModePerm); err != nil {
		panic(err)
	} // if

	copyCmd := exec.Command("cp", "-rT", envpath, workpath)

	if err = copyCmd.Run(); err != nil {
		panic(err)
	} // if

	if err = os.Chdir(workpath); err != nil {
		panic(err)
	} // if

	this.original = original
	this.workpath = workpath
}

// Restore 復原工作目錄
func (this *TestEnv) Restore() {
	if err := os.Chdir(this.original); err != nil {
		panic(err)
	} // if

	if err := os.RemoveAll(this.workpath); err != nil {
		panic(err)
	} // if
}

var rootpath string // 測試資料路徑
var envpath string  // 環境資料路徑
