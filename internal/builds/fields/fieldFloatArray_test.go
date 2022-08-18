package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFloatArray))
}

type SuiteFloatArray struct {
	suite.Suite
}

func (this *SuiteFloatArray) target() *FloatArray {
	return &FloatArray{}
}

func (this *SuiteFloatArray) TestType() {
	assert.Equal(this.T(), "floatArray", this.target().Type())
}

func (this *SuiteFloatArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFloatArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFloatArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123,0.456,0.789", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{}, result)
	result, err = target.ToJsonValue("0.123,0.456,0.789", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, result)
	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
