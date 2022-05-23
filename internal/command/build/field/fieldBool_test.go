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
	jsons := Jsons{}
	realValue := true
	fakeValue := "abc"

	assert.Equal(t, "bool", pkey.TypeExcel(), "type excel failed")
	assert.Equal(t, "bool", pkey.TypeCpp(), "type cpp failed")
	assert.Equal(t, "bool", pkey.TypeCs(), "type cs failed")
	assert.Equal(t, "bool", pkey.TypeGo(), "type go failed")
	assert.Equal(t, false, pkey.Hide(), "hide failed")
	assert.Equal(t, false, pkey.PrimaryKey(), "primary key failed")
	assert.Nil(t, pkey.FillToJsons(jsons, strconv.FormatBool(realValue)), "fill to jsons failed")
	assert.Equal(t, realValue, jsons[pkey.Name], "fill to jsons failed")
	assert.NotNil(t, pkey.FillToJsons(jsons, fakeValue), "fill to jsons failed")
}
