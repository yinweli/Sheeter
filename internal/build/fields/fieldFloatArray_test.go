package fields

import (
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFieldFloatArray))
}

type SuiteFieldFloatArray struct {
	suite.Suite
}

func (this *SuiteFieldFloatArray) target() *FieldFloatArray {
	return &FieldFloatArray{}
}

func (this *SuiteFieldFloatArray) TestType() {
	assert.Equal(this.T(), "floatArray", this.target().Type())
}

func (this *SuiteFieldFloatArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldFloatArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldFloatArray) TestToJsonDefault() {
	assert.Equal(this.T(), []float64{}, this.target().ToJsonDefault())
}

func (this *SuiteFieldFloatArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteFieldFloatArray) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "{0.123,0.456,0.789}", result)
	_, err = target.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
