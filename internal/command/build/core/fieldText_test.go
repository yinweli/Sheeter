package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldText(t *testing.T) {
	field := FieldText{}

	assert.Equal(t, "text", field.TypeExcel())
	assert.Equal(t, "std::string", field.TypeCpp())
	assert.Equal(t, "string", field.TypeCs())
	assert.Equal(t, "string", field.TypeGo())
	assert.Equal(t, false, field.Hide())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform(testdata.Text)
	assert.Nil(t, err)
	assert.Equal(t, testdata.Text, result)
}
