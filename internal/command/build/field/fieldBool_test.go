package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	object := Bool{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "bool", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "bool", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "bool", object.TypeCs(), "type cs failed")
	assert.Equal(t, "bool", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, "true"), "fill to metas failed")
	assert.Equal(t, true, metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}