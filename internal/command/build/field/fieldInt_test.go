package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	object := Int{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "int", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "int", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, "999"), "fill to metas failed")
	assert.Equal(t, 999, metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}