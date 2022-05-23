package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFloat64Array(t *testing.T) {
	object := DoubleArray{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "doubleArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<double>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<double>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, testdata.Float64String()), "fill to metas failed")
	assert.Equal(t, testdata.Float64Array(), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}
