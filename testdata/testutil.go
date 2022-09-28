package testdata

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CompareFileByte 比對檔案內容, 預期資料來自位元陣列
func CompareFileByte(t *testing.T, path string, expected []byte) {
	actual, err := os.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(actual))
}

// CompareFilePath 比對檔案內容, 預期資料來自路徑
func CompareFilePath(t *testing.T, pathActual, pathExpected string) {
	actual, err := os.ReadFile(pathActual)
	assert.Nil(t, err)
	expected, err := os.ReadFile(pathExpected)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(actual))
}
