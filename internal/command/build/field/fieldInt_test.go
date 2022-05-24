package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	var result interface{}
	var err error
	object := Int{}

	assert.Equal(t, "int", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "int32_t", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("999")
	assert.Equal(t, result, int64(999), "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
