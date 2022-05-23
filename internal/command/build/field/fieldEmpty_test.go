package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	object := Empty{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "object", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "", object.TypeCs(), "type cs failed")
	assert.Equal(t, "", object.TypeGo(), "type go failed")
	assert.Equal(t, true, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, "test"), "fill to metas failed")
	assert.Equal(t, 0, len(metas), "fill to metas failed")
}
