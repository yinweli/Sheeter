package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSTemplate(t *testing.T) {
	stemplate := STemplate{}
	assert.Equal(t, "", stemplate.SetLine(2))
	assert.Equal(t, "\n", stemplate.NewLine())
	assert.Equal(t, "\n", stemplate.NewLine())
	assert.Equal(t, "", stemplate.NewLine())
	assert.Equal(t, "", stemplate.NewLine())
	assert.Equal(t, "TestColumn", stemplate.FirstUpper("testColumn"))
	assert.Equal(t, "testColumn", stemplate.FirstLower("TestColumn"))
}
