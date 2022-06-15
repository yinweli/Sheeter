package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldBoolArray(t *testing.T) {
	field := mockFieldBoolArray()
	assert.Equal(t, "boolArray", field.TypeExcel())
	assert.Equal(t, "List<bool>", field.TypeCs())
	assert.Equal(t, "[]bool", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("1,0,1,0,1")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldBoolArray() *FieldBoolArray {
	return &FieldBoolArray{}
}
