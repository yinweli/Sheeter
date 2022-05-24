package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestBoolArray(t *testing.T) {
	object := BoolArray{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "boolArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<bool>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<bool>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]bool", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, testdata.BoolString()), "fill to metas failed")
	assert.Equal(t, testdata.BoolArray(), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
