package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldInt(t *testing.T) {
	suite.Run(t, new(SuiteFieldInt))
}

type SuiteFieldInt struct {
	suite.Suite
}

func (this *SuiteFieldInt) target() *FieldInt {
	return &FieldInt{}
}

func (this *SuiteFieldInt) TestType() {
	assert.Equal(this.T(), "int", this.target().Type())
}

func (this *SuiteFieldInt) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldInt) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldInt) TestToJsonDefault() {
	assert.Equal(this.T(), int64(0), this.target().ToJsonDefault())
}

func (this *SuiteFieldInt) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("?????")
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldInt) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "123456789", result)

	_, err = target.ToLuaValue("?????")
	assert.NotNil(this.T(), err)
}
