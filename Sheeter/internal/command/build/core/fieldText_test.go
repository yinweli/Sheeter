package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldText(t *testing.T) {
	field := mockFieldText()
	assert.Equal(t, "text", field.TypeExcel())
	assert.Equal(t, "std::string", field.TypeCpp())
	assert.Equal(t, "string", field.TypeCs())
	assert.Equal(t, "string", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("ball,book,pack")
	assert.Nil(t, err)
	assert.Equal(t, "ball,book,pack", result)
}

func mockFieldText() *FieldText {
	return &FieldText{}
}
