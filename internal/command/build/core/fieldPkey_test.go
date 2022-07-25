package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldPkey(t *testing.T) {
	suite.Run(t, new(SuiteFieldPkey))
}

type SuiteFieldPkey struct {
	suite.Suite
}

func (this *SuiteFieldPkey) target() *FieldPkey {
	return &FieldPkey{}
}

func (this *SuiteFieldPkey) TestType() {
	assert.Equal(this.T(), "pkey", this.target().Type())
}

func (this *SuiteFieldPkey) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldPkey) TestIsPkey() {
	assert.Equal(this.T(), true, this.target().IsPkey())
}

func (this *SuiteFieldPkey) TestToJsonDefault() {
	assert.Equal(this.T(), int64(0), this.target().ToJsonDefault())
}

func (this *SuiteFieldPkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("?????")
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldPkey) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "123456789", result)

	_, err = target.ToLuaValue("?????")
	assert.NotNil(this.T(), err)
}
