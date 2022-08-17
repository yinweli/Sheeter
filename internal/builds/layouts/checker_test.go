package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/internal/builds/layers"
)

func TestChecker(t *testing.T) {
	suite.Run(t, new(SuiteChecker))
}

type SuiteChecker struct {
	suite.Suite
	item1 layers.Layer
	item2 layers.Layer
	item3 layers.Layer
	item4 layers.Layer
}

func (this *SuiteChecker) SetupSuite() {
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

func (this *SuiteChecker) target() *checker {
	return &checker{}
}

func (this *SuiteChecker) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.check(this.item1))
	assert.True(this.T(), target.check(this.item2))
	assert.True(this.T(), target.check(this.item3))
	assert.False(this.T(), target.check(this.item4))
}
