package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldText(t *testing.T) {
	suite.Run(t, new(SuiteFieldText))
}

type SuiteFieldText struct {
	suite.Suite
}

func (this *SuiteFieldText) target() *FieldText {
	return &FieldText{}
}

func (this *SuiteFieldText) TestType() {
	assert.Equal(this.T(), "text", this.target().Type())
}

func (this *SuiteFieldText) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldText) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldText) TestToJsonDefault() {
	assert.Equal(this.T(), "", this.target().ToJsonDefault())
}

func (this *SuiteFieldText) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}

func (this *SuiteFieldText) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "\"ball,book,pack\"", result)
}
