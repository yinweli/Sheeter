package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/internal/build/layers"
)

func TestCheckLayer(t *testing.T) {
	suite.Run(t, new(SuiteCheckLayer))
}

type SuiteCheckLayer struct {
	suite.Suite
	item1 layers.Layer
	item2 layers.Layer
	item3 layers.Layer
	item4 layers.Layer
}

func (this *SuiteCheckLayer) SetupSuite() {
	this.item1 = layers.Layer{
		Name: "name1",
		Type: 1,
	}
	this.item2 = layers.Layer{
		Name: "name2",
		Type: 2,
	}
	this.item3 = layers.Layer{
		Name: "name3",
		Type: 3,
	}
	this.item4 = layers.Layer{
		Name: "name3",
		Type: 4,
	}
}

func (this *SuiteCheckLayer) target() *checkLayer {
	return &checkLayer{}
}

func (this *SuiteCheckLayer) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.check(this.item1))
	assert.True(this.T(), target.check(this.item2))
	assert.True(this.T(), target.check(this.item3))
	assert.False(this.T(), target.check(this.item4))
}
