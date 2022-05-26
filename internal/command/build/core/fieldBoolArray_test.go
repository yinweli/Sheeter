package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldBoolArray(t *testing.T) {
	field := FieldBoolArray{}

	assert.Equal(t, "boolArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<bool>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<bool>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]bool", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.BoolString())
	assert.Equal(t, testdata.BoolArray(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
