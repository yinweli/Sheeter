package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldIntArray(t *testing.T) {
	field := FieldIntArray{}

	assert.Equal(t, "intArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int32_t>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<int>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int32", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.Int32String())
	assert.Nil(t, err, "transform failed")
	assert.Equal(t, testdata.Int32Array(), result, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
