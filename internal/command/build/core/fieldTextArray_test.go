package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldTextArray(t *testing.T) {
	field := mockFieldTextArray()
	assert.Equal(t, "textArray", field.TypeExcel())
	assert.Equal(t, "List<string>", field.TypeCs())
	assert.Equal(t, "[]string", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("ball,book,pack")
	assert.Nil(t, err)
	assert.Equal(t, []string{"ball", "book", "pack"}, result)
}

func mockFieldTextArray() *FieldTextArray {
	return &FieldTextArray{}
}
