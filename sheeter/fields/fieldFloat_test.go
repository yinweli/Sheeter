package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFloat(t *testing.T) {
	suite.Run(t, new(SuiteFloat))
}

type SuiteFloat struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteFloat) SetupSuite() {
	this.Change("test-field-float")
}

func (this *SuiteFloat) TearDownSuite() {
	this.Restore()
}

func (this *SuiteFloat) target() *Float {
	return &Float{}
}

func (this *SuiteFloat) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"float"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenFloatCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenFloatGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenFloatProto, target.ToTypeProto())
}

func (this *SuiteFloat) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0.123456), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
