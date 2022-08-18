package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestFloat(t *testing.T) {
	suite.Run(t, new(SuiteFloat))
}

type SuiteFloat struct {
	suite.Suite
}

func (this *SuiteFloat) target() *Float {
	return &Float{}
}

func (this *SuiteFloat) TestType() {
	assert.Equal(this.T(), "float", this.target().Type())
}

func (this *SuiteFloat) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFloat) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFloat) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float64(0), result)

	result, err = target.ToJsonValue("0.123456", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.123456, result)

	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
