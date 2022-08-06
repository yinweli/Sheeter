package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteBoolArray))
}

type SuiteBoolArray struct {
	suite.Suite
}

func (this *SuiteBoolArray) target() *BoolArray {
	return &BoolArray{}
}

func (this *SuiteBoolArray) TestType() {
	assert.Equal(this.T(), "boolArray", this.target().Type())
}

func (this *SuiteBoolArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteBoolArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteBoolArray) TestToJsonDefault() {
	assert.Equal(this.T(), []bool{}, this.target().ToJsonDefault())
}

func (this *SuiteBoolArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteBoolArray) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "{true,false,true,false,true}", result)

	_, err = target.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
