package testdata

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CompareFile 比對檔案內容
func CompareFile(t *testing.T, filepath string, expected []byte) {
	actual, err := os.ReadFile(filepath)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
