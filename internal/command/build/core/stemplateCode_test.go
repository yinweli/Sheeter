package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSTemplateCode(t *testing.T) {
	stemplateCode := STemplateCode{
		STemplate: STemplate{
			OriginalName: "test1",
			StructName:   "test2",
		},
		JsonFileName: "test3",
		Columns: []*Column{
			{Field: &FieldInt{}},
			{Field: &FieldInt{}},
			{Field: &FieldEmpty{}},
		},
	}
	assert.Equal(t, CsNamespace, stemplateCode.CsNamespace())
	assert.Equal(t, GoPackage, stemplateCode.GoPackage())
	assert.Equal(t, "", stemplateCode.SetLine())
	assert.Equal(t, "\n", stemplateCode.NewLine())
	assert.Equal(t, "", stemplateCode.NewLine())
	assert.Equal(t, "", stemplateCode.NewLine())

	bytes, err := stemplateCode.Generate("{{.OriginalName}}#{{.StructName}}#{{.JsonFileName}}#{{.CsNamespace}}#{{.GoPackage}}")
	assert.Nil(t, err)
	assert.Equal(t, "test1#test2#test3#Sheeter#sheeter", string(bytes[:]))
	bytes, err = stemplateCode.Generate("{{{}}")
	assert.NotNil(t, err)
	bytes, err = stemplateCode.Generate("{{.Unknown}}")
	assert.NotNil(t, err)
}
