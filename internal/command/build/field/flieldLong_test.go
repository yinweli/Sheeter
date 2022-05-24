package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLong(t *testing.T) {
	object := Long{
		Data{
			Note:  "note",
			Name:  "name",
			Field: "field",
		},
	}
	metas := Metas{}

	assert.Equal(t, "long", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "int64_t", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "long", object.TypeCs(), "type cs failed")
	assert.Equal(t, "int64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, "999999999999"), "fill to metas failed")
	assert.Equal(t, int64(999999999999), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
