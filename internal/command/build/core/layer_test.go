package core

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
	assert.Equal(this.T(), layer[0].Type, LayerArray)
	assert.Equal(this.T(), layer[0].Name, "name1")
	assert.Equal(this.T(), 0, back)

	layer, back, err = ParseLayer("{name1")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 1)
	assert.Equal(this.T(), layer[0].Type, LayerStruct)
	assert.Equal(this.T(), layer[0].Name, "name1")
	assert.Equal(this.T(), 0, back)

	layer, back, err = ParseLayer("}")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 0)
	assert.Equal(this.T(), 1, back)

	layer, back, err = ParseLayer("{[]name1 {name2 {[]name3 {name4 {name5 }}}}}")
	assert.Nil(this.T(), err)
	assert.Len(this.T(), layer, 5)
	assert.Equal(this.T(), layer[0].Type, LayerArray)
	assert.Equal(this.T(), layer[0].Name, "name1")
	assert.Equal(this.T(), layer[1].Type, LayerStruct)
	assert.Equal(this.T(), layer[1].Name, "name2")
	assert.Equal(this.T(), layer[2].Type, LayerArray)
	assert.Equal(this.T(), layer[2].Name, "name3")
	assert.Equal(this.T(), layer[3].Type, LayerStruct)
	assert.Equal(this.T(), layer[3].Name, "name4")
	assert.Equal(this.T(), layer[4].Type, LayerStruct)
	assert.Equal(this.T(), layer[4].Name, "name5")
	assert.Equal(this.T(), 5, back)

	layer, back, err = ParseLayer("{ name1 ")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{[name1")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{]name1")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{name1+")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{name1 } {name2")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{name1 }{name2")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{name1 }name2")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("{name1 name2")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer("} {name1 ")
	assert.NotNil(this.T(), err)

	layer, back, err = ParseLayer(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
