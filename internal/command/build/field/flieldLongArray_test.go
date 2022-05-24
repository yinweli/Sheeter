package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestLongArray(t *testing.T) {
	object := LongArray{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "longArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int64_t>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<long>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Equal(t, "note", object.GetNote(), "get note failed")
	assert.Equal(t, "name", object.GetName(), "get name failed")
	assert.Equal(t, "field", object.GetField(), "get field failed")
	assert.Nil(t, object.FillToMetas(metas, testdata.Int64String()), "fill to metas failed")
	assert.Equal(t, testdata.Int64Array(), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
