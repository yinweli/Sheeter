package field

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkey(t *testing.T) {
	pkey := Pkey{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "pkey", pkey.TypeExcel(), "type excel failed")
	assert.Equal(t, "Sheet::pkey", pkey.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", pkey.TypeCs(), "type cs failed")
	assert.Equal(t, "int", pkey.TypeGo(), "type go failed")
	assert.Equal(t, false, pkey.Hide(), "hide failed")
	assert.Equal(t, true, pkey.PrimaryKey(), "primary key failed")
	assert.Nil(t, pkey.FillToMetas(metas, strconv.Itoa(999)), "fill to metas failed")
	assert.Equal(t, 999, metas[pkey.Name], "fill to metas failed")
	assert.NotNil(t, pkey.FillToMetas(metas, "abc"), "fill to metas failed")
}
