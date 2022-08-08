package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInt(t *testing.T) {
	suite.Run(t, new(SuiteInt))
}

type SuiteInt struct {
	suite.Suite
}

func (this *SuiteInt) target() *Int {
	return &Int{}
}

func (this *SuiteInt) TestType() {
	assert.Equal(this.T(), "int", this.target().Type())
}

func (this *SuiteInt) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteInt) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteInt) TestToJsonDefault() {
	assert.Equal(this.T(), int64(0), this.target().ToJsonDefault())
}

func (this *SuiteInt) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteInt) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "123456789", result)

	_, err = target.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
