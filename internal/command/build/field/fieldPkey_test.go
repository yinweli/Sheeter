package field

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkey(t *testing.T) {
	object := Pkey{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "pkey", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "Sheet::pkey", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, true, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, strconv.Itoa(999)), "fill to metas failed")
	assert.Equal(t, 999, metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
