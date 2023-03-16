package testdata

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
)

// TestData 測試資料
type TestData struct {
	// 通用測試

	Unknown string // 未知字串

	// 測試環境

	original string // 原始路徑
	workpath string // 工作路徑
}

// TBegin 開始測試
func (this *TestData) TBegin(work, data string) {
	// 初始化通用測試

	this.Unknown = "?????"

	// 初始化測試環境

	original, err := os.Getwd()

	if err != nil {
		panic(err)
	} // if

	workpath := filepath.Join(rootpath, work)

	if err = os.MkdirAll(workpath, os.ModePerm); err != nil {
		panic(err)
	} // if

	datapath := filepath.Join(envpath, data)

	if err = copy.Copy(datapath, workpath); err != nil {
		panic(err)
	} // if

	if err = os.Chdir(workpath); err != nil {
		panic(err)
	} // if

	this.original = original
	this.workpath = workpath
}

// TFinal 結束測試
func (this *TestData) TFinal() {
	if err := os.Chdir(this.original); err != nil {
		panic(err)
	} // if

	if err := os.RemoveAll(this.workpath); err != nil {
		panic(err)
	} // if
}

// AssertCompareFile 比對檔案內容, 預期資料來自位元陣列
func (this *TestData) AssertCompareFile(t *testing.T, path string, expected []byte) {
	actual, err := os.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(actual))
}

func init() {
	_, file, _, ok := runtime.Caller(0)

	if ok == false {
		panic("get rootpath failed")
	} // if

	rootpath = filepath.Dir(file)
	envpath = filepath.Join(rootpath, "env")

	// 如果env資料夾不存在, 就建立一個, 免得測試時拋出錯誤
	if _, err := os.Stat(envpath); os.IsNotExist(err) {
		if err = os.MkdirAll(envpath, os.ModePerm); err != nil {
			panic(err)
		} // if
	} // if
}

var rootpath string // 根路徑
var envpath string  // 環境路徑
