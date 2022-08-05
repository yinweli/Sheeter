package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

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
	return NewDuplLayer()
}

func (this *SuiteDuplLayer) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.Check(this.item1, this.item2, this.item3))
	assert.True(this.T(), target.Check(this.item1, this.item2, this.item3))
	assert.False(this.T(), target.Check(this.itemx))
}

func (this *SuiteDuplLayer) TestNewDuplLayer() {
	assert.NotNil(this.T(), NewDuplLayer())
}
