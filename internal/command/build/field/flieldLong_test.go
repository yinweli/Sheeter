package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLong(t *testing.T) {
	var result interface{}
	var err error
	var object Long

	assert.Equal(t, "long", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "int64_t", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "long", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("999999999999")
	assert.Equal(t, int64(999999999999), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
