package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	template := mockTemplate()
	assert.Equal(t, "", template.SetLine(2))
	assert.Equal(t, "\n", template.NewLine())
	assert.Equal(t, "\n", template.NewLine())
	assert.Equal(t, "", template.NewLine())
	assert.Equal(t, "", template.NewLine())
	assert.Equal(t, "TestColumn", template.FirstUpper("testColumn"))
	assert.Equal(t, "testColumn", template.FirstLower("TestColumn"))
}

func TestNewTemplate(t *testing.T) {
	template := NewTemplate("", "")
	assert.NotNil(t, template)
}

func mockTemplate() *Template {
	return &Template{}
}
