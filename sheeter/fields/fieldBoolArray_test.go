package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteBoolArray))
}

type SuiteBoolArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteBoolArray) SetupSuite() {
	this.Change("test-field-boolArray")
}

func (this *SuiteBoolArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteBoolArray) target() *BoolArray {
	return &BoolArray{}
}

func (this *SuiteBoolArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "boolArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenBoolCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenBoolGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenBoolProto, target.ToTypeProto())
}

func (this *SuiteBoolArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
