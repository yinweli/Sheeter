package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteBoolArray))
}

type SuiteBoolArray struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteBoolArray) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-fields-boolArray")
}

func (this *SuiteBoolArray) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteBoolArray) TestField() {
	target := &BoolArray{}
	assert.Equal(this.T(), []string{"boolArray", "[]bool", "bool[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeBoolCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeBoolGo, target.ToTypeGo())
}

func (this *SuiteBoolArray) TestToJsonValue() {
	target := &BoolArray{}

	result, err := target.ToJsonValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{}, result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
