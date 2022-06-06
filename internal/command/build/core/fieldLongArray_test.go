package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldLongArray(t *testing.T) {
	field := FieldLongArray{}
	assert.Equal(t, "longArray", field.TypeExcel())
	assert.Equal(t, "std::vector<int64_t>", field.TypeCpp())
	assert.Equal(t, "List<long>", field.TypeCs())
	assert.Equal(t, "[]int64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform(testdata.Int64String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Int64Array(), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
