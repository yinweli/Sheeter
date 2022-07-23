package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldBool(t *testing.T) {
	suite.Run(t, new(SuiteFieldBool))
}

type SuiteFieldBool struct {
	suite.Suite
}

func (this *SuiteFieldBool) target() *FieldBool {
	return &FieldBool{}
}

func (this *SuiteFieldBool) TestType() {
	assert.Equal(this.T(), "bool", this.target().Type())
}

func (this *SuiteFieldBool) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldBool) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldBool) TestToJsonDefault() {
	assert.Equal(this.T(), false, this.target().ToJsonDefault())
}

func (this *SuiteFieldBool) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, result)
	result, err = target.ToJsonValue("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)
	_, err = target.ToJsonValue("?????")
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldBool) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true", result)
	result, err = target.ToLuaValue("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "false", result)
	_, err = target.ToLuaValue("?????")
	assert.NotNil(this.T(), err)
}
