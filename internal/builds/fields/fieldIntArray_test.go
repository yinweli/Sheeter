package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestIntArray(t *testing.T) {
	suite.Run(t, new(SuiteIntArray))
}

type SuiteIntArray struct {
	suite.Suite
}

func (this *SuiteIntArray) target() *IntArray {
	return &IntArray{}
}

func (this *SuiteIntArray) TestType() {
	assert.Equal(this.T(), "intArray", this.target().Type())
}

func (this *SuiteIntArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteIntArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteIntArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{}, result)
	result, err = target.ToJsonValue("123,456,789", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)
	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
