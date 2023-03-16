package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestIntArray(t *testing.T) {
	suite.Run(t, new(SuiteIntArray))
}

type SuiteIntArray struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteIntArray) SetupSuite() {
	this.TBegin("test-fields-intArray", "")
}

func (this *SuiteIntArray) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteIntArray) TestField() {
	target := &IntArray{}
	assert.Equal(this.T(), []string{"intArray", "[]int", "int[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeIntCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeIntGo, target.ToTypeGo())
}

func (this *SuiteIntArray) TestToJsonValue() {
	target := &IntArray{}

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int32{123, 456, 789}, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int32{}, result)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
