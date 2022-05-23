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
	jsons := Jsons{}
	realValue := 999
	fakeValue := "abc"

	assert.Equal(t, "pkey", pkey.TypeExcel(), "type excel failed")
	assert.Equal(t, "Sheet::pkey", pkey.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", pkey.TypeCs(), "type cs failed")
	assert.Equal(t, "int", pkey.TypeGo(), "type go failed")
	assert.Equal(t, false, pkey.Hide(), "hide failed")
	assert.Equal(t, true, pkey.PrimaryKey(), "primary key failed")
	assert.Nil(t, pkey.FillToJsons(jsons, strconv.Itoa(realValue)), "fill to jsons failed")
	assert.Equal(t, realValue, jsons[pkey.Name], "fill to jsons failed")
	assert.NotNil(t, pkey.FillToJsons(jsons, fakeValue), "fill to jsons failed")
}
