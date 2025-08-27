package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestFloat(t *testing.T) {
	suite.Run(t, new(SuiteFloat))
}

type SuiteFloat struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteFloat) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-float")
}

func (this *SuiteFloat) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteFloat) TestField() {
	target := &Float{}
	assert.Equal(this.T(), []string{"float"}, target.Field())
	assert.Equal(this.T(), "Single", target.ToTypeCs())
	assert.Equal(this.T(), "float32", target.ToTypeGo())
}

func (this *SuiteFloat) TestToJsonValue() {
	target := &Float{}

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0.123456), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0), result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
