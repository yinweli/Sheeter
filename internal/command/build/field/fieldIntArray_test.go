package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestIntArray(t *testing.T) {
	object := IntArray{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "intArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int32_t>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<int>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Equal(t, "note", object.GetNote(), "get note failed")
	assert.Equal(t, "name", object.GetName(), "get name failed")
	assert.Equal(t, "field", object.GetField(), "get field failed")
	assert.Nil(t, object.FillToMetas(metas, testdata.Int32String()), "fill to metas failed")
	assert.Equal(t, testdata.Int32Array(), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
