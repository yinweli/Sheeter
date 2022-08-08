package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestEmpty(t *testing.T) {
	suite.Run(t, new(SuiteEmpty))
}

type SuiteEmpty struct {
	suite.Suite
}

func (this *SuiteEmpty) target() *Empty {
	return &Empty{}
}

func (this *SuiteEmpty) TestType() {
	assert.Equal(this.T(), "empty", this.target().Type())
}

func (this *SuiteEmpty) TestIsShow() {
	assert.Equal(this.T(), false, this.target().IsShow())
}

func (this *SuiteEmpty) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteEmpty) TestToJsonDefault() {
	assert.Equal(this.T(), nil, this.target().ToJsonDefault())
}

func (this *SuiteEmpty) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("test")
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), result)
}

func (this *SuiteEmpty) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("test")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "", result)
}
