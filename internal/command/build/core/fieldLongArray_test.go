package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldLongArray(t *testing.T) {
	field := FieldLongArray{}

	assert.Equal(t, "longArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int64_t>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<long>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int64", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.Int64String())
	assert.Equal(t, testdata.Int64Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
