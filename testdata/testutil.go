package testdata

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// IsWindows 取得是否在windows下執行
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// CompareFile 比對檔案內容, 預期資料來自位元陣列
func CompareFile(t *testing.T, path string, expected []byte) {
	actual, err := os.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(actual))
}
