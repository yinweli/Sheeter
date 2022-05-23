package field

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	pkey := Bool{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "bool", pkey.TypeExcel(), "type excel failed")
	assert.Equal(t, "bool", pkey.TypeCpp(), "type cpp failed")
	assert.Equal(t, "bool", pkey.TypeCs(), "type cs failed")
	assert.Equal(t, "bool", pkey.TypeGo(), "type go failed")
	assert.Equal(t, false, pkey.Hide(), "hide failed")
	assert.Equal(t, false, pkey.PrimaryKey(), "primary key failed")
	assert.Nil(t, pkey.FillToMetas(metas, strconv.FormatBool(true)), "fill to metas failed")
	assert.Equal(t, true, metas[pkey.Name], "fill to metas failed")
	assert.NotNil(t, pkey.FillToMetas(metas, "abc"), "fill to metas failed")
}
