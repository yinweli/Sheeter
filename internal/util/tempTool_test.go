package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTempTool(t *testing.T) {
	tempTool := mockTempTool()
	assert.Equal(t, "", tempTool.SetLine(2))
	assert.Equal(t, "\n", tempTool.NewLine())
	assert.Equal(t, "\n", tempTool.NewLine())
	assert.Equal(t, "", tempTool.NewLine())
	assert.Equal(t, "", tempTool.NewLine())
	assert.Equal(t, "TestColumn", tempTool.FirstUpper("testColumn"))
	assert.Equal(t, "testColumn", tempTool.FirstLower("TestColumn"))
}

func mockTempTool() *TempTool {
	return &TempTool{}
}
