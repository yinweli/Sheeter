package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteFieldBoolArray))
}

type SuiteFieldBoolArray struct {
	suite.Suite
}

func (this *SuiteFieldBoolArray) target() *FieldBoolArray {
	return &FieldBoolArray{}
}

func (this *SuiteFieldBoolArray) TestType() {
	assert.Equal(this.T(), "boolArray", this.target().Type())
}

func (this *SuiteFieldBoolArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldBoolArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldBoolArray) TestToJsonDefault() {
	assert.Equal(this.T(), []bool{}, this.target().ToJsonDefault())
}

func (this *SuiteFieldBoolArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	_, err = target.ToJsonValue("?????")
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldBoolArray) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "{true,false,true,false,true}", result)

	_, err = target.ToLuaValue("?????")
	assert.NotNil(this.T(), err)
}
