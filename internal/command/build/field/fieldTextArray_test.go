package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTextArray(t *testing.T) {
	var result interface{}
	var err error
	var object TextArray

	assert.Equal(t, "textArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<std::string>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<string>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]string", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.StringString())
	assert.Equal(t, testdata.StringArray(), result, "transform failed")
	assert.Nil(t, err, "transform failed")
}
