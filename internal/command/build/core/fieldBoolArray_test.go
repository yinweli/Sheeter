package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldBoolArray(t *testing.T) {
	field := mockFieldBoolArray()
	assert.Equal(t, "boolArray", field.TypeExcel())
	assert.Equal(t, "std::vector<bool>", field.TypeCpp())
	assert.Equal(t, "List<bool>", field.TypeCs())
	assert.Equal(t, "[]bool", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform(testdata.BoolString())
	assert.Nil(t, err)
	assert.Equal(t, testdata.BoolArray(), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}

func mockFieldBoolArray() *FieldBoolArray {
	return &FieldBoolArray{}
}
