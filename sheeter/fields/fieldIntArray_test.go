package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestIntArray(t *testing.T) {
	suite.Run(t, new(SuiteIntArray))
}

type SuiteIntArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteIntArray) SetupSuite() {
	this.Change("test-field-intArray")
}

func (this *SuiteIntArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteIntArray) target() *IntArray {
	return &IntArray{}
}

func (this *SuiteIntArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "intArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenIntCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenIntGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenIntProto, target.ToTypeProto())
}

func (this *SuiteIntArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
