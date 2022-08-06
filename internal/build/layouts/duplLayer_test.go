package layouts

import (
	"testing"

	"github.com/yinweli/Sheeter/internal/build/layers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDuplLayer(t *testing.T) {
	suite.Run(t, new(SuiteDuplLayer))
}

type SuiteDuplLayer struct {
	suite.Suite
	item1 layers.Layer
	item2 layers.Layer
	item3 layers.Layer
	itemx layers.Layer
}

func (this *SuiteDuplLayer) SetupSuite() {
	this.item1 = layers.Layer{Name: "001", Type: layers.LayerArray}
	this.item2 = layers.Layer{Name: "002", Type: layers.LayerStruct}
	this.item3 = layers.Layer{Name: "003", Type: layers.LayerStruct}
	this.itemx = layers.Layer{Name: "001", Type: layers.LayerStruct}
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
