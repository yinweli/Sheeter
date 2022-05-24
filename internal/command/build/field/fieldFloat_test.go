package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	object := Float{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "float", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "float", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "float", object.TypeCs(), "type cs failed")
	assert.Equal(t, "float32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Equal(t, "note", object.GetNote(), "get note failed")
	assert.Equal(t, "name", object.GetName(), "get name failed")
	assert.Equal(t, "field", object.GetField(), "get field failed")
	assert.Nil(t, object.FillToMetas(metas, "0.999"), "fill to metas failed")
	assert.InDelta(t, float32(0.999), metas[object.Name], 0.0001, "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
