package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldFloat(t *testing.T) {
	suite.Run(t, new(SuiteFieldFloat))
}

type SuiteFieldFloat struct {
	suite.Suite
}

func (this *SuiteFieldFloat) target() *FieldFloat {
	return &FieldFloat{}
}

func (this *SuiteFieldFloat) TestType() {
	assert.Equal(this.T(), "float", this.target().Type())
}

func (this *SuiteFieldFloat) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldFloat) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldFloat) TestToJsonDefault() {
	assert.Equal(this.T(), float64(0), this.target().ToJsonDefault())
}

func (this *SuiteFieldFloat) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.123456, result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldFloat) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "0.123456", result)

	_, err = target.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
