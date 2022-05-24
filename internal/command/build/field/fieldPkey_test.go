package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkey(t *testing.T) {
	var result interface{}
	var err error
	object := Pkey{}

	assert.Equal(t, "pkey", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "Sheet::pkey", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, true, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("999")
	assert.Equal(t, result, 999, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
