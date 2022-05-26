package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloatArray(t *testing.T) {
	field := FieldFloatArray{}

	assert.Equal(t, "floatArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<float>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<float>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float32", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.Float32String())
	assert.Equal(t, testdata.Float32Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")

}
