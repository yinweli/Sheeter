package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldTextArray(t *testing.T) {
	field := FieldTextArray{}

	assert.Equal(t, "textArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<std::string>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<string>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]string", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.StringString())
	assert.Nil(t, err, "transform failed")
	assert.Equal(t, testdata.StringArray(), result, "transform failed")
}
