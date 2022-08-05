package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLayer(t *testing.T) {
	suite.Run(t, new(SuiteLayer))
}

type SuiteLayer struct {
	suite.Suite
}

func (this *SuiteLayer) TestParseLayer() {
	layer, back, err := ParseLayer("{[]name1")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 1)
	assert.Equal(this.T(), LayerArray, layer[0].Type)
	assert.Equal(this.T(), "name1", layer[0].Name)
	assert.Equal(this.T(), 0, back)

	layer, back, err = ParseLayer("{name1")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 1)
	assert.Equal(this.T(), LayerStruct, layer[0].Type)
	assert.Equal(this.T(), "name1", layer[0].Name)
	assert.Equal(this.T(), 0, back)

	layer, back, err = ParseLayer("/")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 1)
	assert.Equal(this.T(), LayerDivider, layer[0].Type)
	assert.Equal(this.T(), "", layer[0].Name)
	assert.Equal(this.T(), 0, back)

	layer, back, err = ParseLayer("}")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 0)
	assert.Equal(this.T(), 1, back)

	layer, back, err = ParseLayer("/ {[]name1 {name2 {[]name3 {name4 {name5 }}}}}")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 6)
	assert.Equal(this.T(), LayerDivider, layer[0].Type)
	assert.Equal(this.T(), "", layer[0].Name)
	assert.Equal(this.T(), LayerArray, layer[1].Type)
	assert.Equal(this.T(), "name1", layer[1].Name)
	assert.Equal(this.T(), LayerStruct, layer[2].Type)
	assert.Equal(this.T(), "name2", layer[2].Name)
	assert.Equal(this.T(), LayerArray, layer[3].Type)
	assert.Equal(this.T(), "name3", layer[3].Name)
	assert.Equal(this.T(), LayerStruct, layer[4].Type)
	assert.Equal(this.T(), "name4", layer[4].Name)
	assert.Equal(this.T(), LayerStruct, layer[5].Type)
	assert.Equal(this.T(), "name5", layer[5].Name)
	assert.Equal(this.T(), 5, back)

	_, _, err = ParseLayer("{ name1")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{[] name1")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{[name1")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{]name1")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1+")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1 } {name2")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1 }{name2")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1 }name2")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1 name2")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("} {name1 ")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("{name1 /")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("/ /")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer("/name1")
	assert.NotNil(this.T(), err)

	_, _, err = ParseLayer(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
