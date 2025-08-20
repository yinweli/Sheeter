package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestDoubleArray(t *testing.T) {
	suite.Run(t, new(SuiteDoubleArray))
}

type SuiteDoubleArray struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteDoubleArray) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-doubleArray")
}

func (this *SuiteDoubleArray) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteDoubleArray) TestField() {
	target := &DoubleArray{}
	assert.Equal(this.T(), []string{"doubleArray", "[]double", "double[]"}, target.Field())
	assert.Equal(this.T(), "double[]", target.ToTypeCs())
	assert.Equal(this.T(), "[]float64", target.ToTypeGo())
}

func (this *SuiteDoubleArray) TestToJsonValue() {
	target := &DoubleArray{}

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{}, result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
