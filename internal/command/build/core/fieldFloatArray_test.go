package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFloatArray(t *testing.T) {
	var result interface{}
	var err error
	var object FieldFloatArray

	assert.Equal(t, "floatArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<float>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<float>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.Float32String())
	assert.Equal(t, testdata.Float32Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")

}
