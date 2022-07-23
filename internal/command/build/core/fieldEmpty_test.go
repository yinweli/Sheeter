package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldEmpty(t *testing.T) {
	suite.Run(t, new(SuiteFieldEmpty))
}

type SuiteFieldEmpty struct {
	suite.Suite
}

func (this *SuiteFieldEmpty) target() *FieldEmpty {
	return &FieldEmpty{}
}

func (this *SuiteFieldEmpty) TestType() {
	assert.Equal(this.T(), "empty", this.target().Type())
}

func (this *SuiteFieldEmpty) TestIsShow() {
	assert.Equal(this.T(), false, this.target().IsShow())
}

func (this *SuiteFieldEmpty) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldEmpty) TestToJsonDefault() {
	assert.Equal(this.T(), nil, this.target().ToJsonDefault())
}

func (this *SuiteFieldEmpty) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("test")
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), result)
}

func (this *SuiteFieldEmpty) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("test")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "", result)
}
