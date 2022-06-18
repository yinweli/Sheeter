package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateCode(t *testing.T) {
	templateCode := mockTemplateCode()
	assert.Equal(t, "", templateCode.SetLine())
	assert.Equal(t, "\n", templateCode.NewLine())
	assert.Equal(t, "", templateCode.NewLine())
	assert.Equal(t, "", templateCode.NewLine())

	bytes, err := templateCode.Generate("{{.Namespace}}#{{.JsonFileName}}")
	assert.Nil(t, err)
	assert.Equal(t, "test1#test2", string(bytes[:]))
	bytes, err = templateCode.Generate("{{{}}")
	assert.NotNil(t, err)
	bytes, err = templateCode.Generate("{{.Unknown}}")
	assert.NotNil(t, err)
}

func TestNewTemplateCode(t *testing.T) {
	templateCode := NewTemplateCode("", "", "")
	assert.NotNil(t, templateCode)
}

func mockTemplateCode() *TemplateCode {
	return &TemplateCode{
		Namespace:    "test1",
		JsonFileName: "test2",
		Columns: []*Column{
			{Field: &FieldInt{}},
			{Field: &FieldInt{}},
			{Field: &FieldEmpty{}},
		},
	}
}
