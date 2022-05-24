package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestLongArray(t *testing.T) {
	var result interface{}
	var err error
	object := LongArray{}

	assert.Equal(t, "longArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int64_t>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<long>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.Int64String())
	assert.Equal(t, result, testdata.Int64Array(), "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
