package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldTextArray(t *testing.T) {
	field := mockFieldTextArray()
	assert.Equal(t, "textArray", field.TypeExcel())
	assert.Equal(t, "std::vector<std::string>", field.TypeCpp())
	assert.Equal(t, "List<string>", field.TypeCs())
	assert.Equal(t, "[]string", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform(testdata.StringString())
	assert.Nil(t, err)
	assert.Equal(t, testdata.StringArray(), result)
}

func mockFieldTextArray() *FieldTextArray {
	return &FieldTextArray{}
}
