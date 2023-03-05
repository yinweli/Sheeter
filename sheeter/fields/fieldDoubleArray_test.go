package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestDoubleArray(t *testing.T) {
	suite.Run(t, new(SuiteDoubleArray))
}

type SuiteDoubleArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteDoubleArray) SetupSuite() {
	this.Change("test-field-doubleArray")
}

func (this *SuiteDoubleArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteDoubleArray) target() *DoubleArray {
	return &DoubleArray{}
}

func (this *SuiteDoubleArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"doubleArray", "[]double", "double[]"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenDoubleCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenDoubleGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenDoubleProto, target.ToTypeProto())
}

func (this *SuiteDoubleArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
