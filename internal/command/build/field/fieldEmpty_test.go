package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	empty := Empty{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "empty", empty.TypeExcel(), "type excel failed")
	assert.Equal(t, "", empty.TypeCpp(), "type cpp failed")
	assert.Equal(t, "", empty.TypeCs(), "type cs failed")
	assert.Equal(t, "", empty.TypeGo(), "type go failed")
	assert.Equal(t, true, empty.Hide(), "hide failed")
	assert.Equal(t, false, empty.PrimaryKey(), "primary key failed")
	assert.Nil(t, empty.FillToMetas(metas, "test"), "fill to metas failed")
	assert.Equal(t, 0, len(metas), "fill to metas failed")
}
