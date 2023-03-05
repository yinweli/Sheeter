package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLongArray(t *testing.T) {
	suite.Run(t, new(SuiteLongArray))
}

type SuiteLongArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteLongArray) SetupSuite() {
	this.Change("test-field-longArray")
}

func (this *SuiteLongArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteLongArray) target() *LongArray {
	return &LongArray{}
}

func (this *SuiteLongArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"longArray", "[]long", "long[]"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenLongCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenLongGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenLongProto, target.ToTypeProto())
}

func (this *SuiteLongArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
