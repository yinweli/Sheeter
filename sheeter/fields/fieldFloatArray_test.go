package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFloatArray))
}

type SuiteFloatArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteFloatArray) SetupSuite() {
	this.Change("test-field-floatArray")
}

func (this *SuiteFloatArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteFloatArray) target() *FloatArray {
	return &FloatArray{}
}

func (this *SuiteFloatArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "floatArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenFloatCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenFloatGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenFloatProto, target.ToTypeProto())
}

func (this *SuiteFloatArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
