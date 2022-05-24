package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	object := Double{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "double", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "double", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "double", object.TypeCs(), "type cs failed")
	assert.Equal(t, "float64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Equal(t, "note", object.GetNote(), "get note failed")
	assert.Equal(t, "name", object.GetName(), "get name failed")
	assert.Equal(t, "field", object.GetField(), "get field failed")
	assert.Nil(t, object.FillToMetas(metas, "0.99999999"), "fill to metas failed")
	assert.InDelta(t, 0.99999999, metas[object.Name], 0.000000001, "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
