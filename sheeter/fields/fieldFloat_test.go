package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestFloat(t *testing.T) {
	suite.Run(t, new(SuiteFloat))
}

type SuiteFloat struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteFloat) SetupSuite() {
	this.TBegin("test-fields-float", "")
}

func (this *SuiteFloat) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteFloat) TestField() {
	target := &Float{}
	assert.Equal(this.T(), []string{"float"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeFloatCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeFloatGo, target.ToTypeGo())
}

func (this *SuiteFloat) TestToJsonValue() {
	target := &Float{}

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0.123456), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0), result)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
