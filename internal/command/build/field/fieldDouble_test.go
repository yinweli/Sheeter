package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	var result interface{}
	var err error
	object := Double{}

	assert.Equal(t, "double", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "double", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "double", object.TypeCs(), "type cs failed")
	assert.Equal(t, "float64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("0.99999999")
	assert.InDelta(t, 0.99999999, result, 0.00000001, "fill to metas failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
