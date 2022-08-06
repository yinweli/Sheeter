package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldIntArray(t *testing.T) {
	suite.Run(t, new(SuiteFieldIntArray))
}

type SuiteFieldIntArray struct {
	suite.Suite
}

func (this *SuiteFieldIntArray) target() *FieldIntArray {
	return &FieldIntArray{}
}

func (this *SuiteFieldIntArray) TestType() {
	assert.Equal(this.T(), "intArray", this.target().Type())
}

func (this *SuiteFieldIntArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldIntArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldIntArray) TestToJsonDefault() {
	assert.Equal(this.T(), []int64{}, this.target().ToJsonDefault())
}

func (this *SuiteFieldIntArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldIntArray) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "{123,456,789}", result)

	_, err = target.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
