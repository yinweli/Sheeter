package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestLongArray(t *testing.T) {
	suite.Run(t, new(SuiteLongArray))
}

type SuiteLongArray struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteLongArray) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-longArray")
}

func (this *SuiteLongArray) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteLongArray) TestField() {
	target := &LongArray{}
	assert.Equal(this.T(), []string{"longArray", "[]long", "long[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeLongCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeLongGo, target.ToTypeGo())
}

func (this *SuiteLongArray) TestToJsonValue() {
	target := &LongArray{}

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{}, result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
