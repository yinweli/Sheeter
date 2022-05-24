package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	var result interface{}
	var err error
	object := Empty{}

	assert.Equal(t, "empty", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "", object.TypeCs(), "type cs failed")
	assert.Equal(t, "", object.TypeGo(), "type go failed")
	assert.Equal(t, true, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("test")
	assert.Nil(t, result, "transform failed")
	assert.Nil(t, err, "transform failed")
}
