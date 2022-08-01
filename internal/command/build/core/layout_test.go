package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDuplField(t *testing.T) {
	suite.Run(t, new(SuiteDuplField))
}

type SuiteDuplField struct {
	suite.Suite
	item1 string
	item2 string
	item3 string
}

func (this *SuiteDuplField) SetupSuite() {
	this.item1 = "001"
	this.item2 = "001/002"
	this.item3 = "001/002/003"
}

func (this *SuiteDuplField) target() *duplField {
	return &duplField{
		datas: map[string]bool{},
	}
}

func (this *SuiteDuplField) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.Check(this.item1))
	assert.True(this.T(), target.Check(this.item2))
	assert.True(this.T(), target.Check(this.item3))
	assert.False(this.T(), target.Check(this.item1))
	assert.False(this.T(), target.Check(this.item2))
	assert.False(this.T(), target.Check(this.item3))
}

func TestDuplLayer(t *testing.T) {
	suite.Run(t, new(SuiteDuplLayer))
}

type SuiteDuplLayer struct {
	suite.Suite
	item1 Layer
	item2 Layer
	item3 Layer
	itemx Layer
}

func (this *SuiteDuplLayer) SetupSuite() {
	this.item1 = Layer{Name: "001", Type: LayerArray}
	this.item2 = Layer{Name: "002", Type: LayerStruct}
	this.item3 = Layer{Name: "003", Type: LayerStruct}
	this.itemx = Layer{Name: "001", Type: LayerStruct}
}

func (this *SuiteDuplLayer) target() *duplLayer {
	return &duplLayer{
		datas: map[string]int{},
	}
}

func (this *SuiteDuplLayer) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.Check(this.item1, this.item2, this.item3))
	assert.True(this.T(), target.Check(this.item1, this.item2, this.item3))
	assert.False(this.T(), target.Check(this.itemx))
}
