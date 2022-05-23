package field

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFloatArray(t *testing.T) {
	object := FloatArray{
		Data{
			Raw:  "raw",
			Note: "note",
			Name: "name",
		},
	}
	metas := Metas{}

	assert.Equal(t, "floatArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<float>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<float>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")
	assert.Nil(t, object.FillToMetas(metas, testdata.Float32String()), "fill to metas failed")
	assert.Equal(t, testdata.Float32Array(), metas[object.Name], "fill to metas failed")
	assert.NotNil(t, object.FillToMetas(metas, "fake"), "fill to metas failed")
}