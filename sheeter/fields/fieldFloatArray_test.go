package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFloatArray))
}

type SuiteFloatArray struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteFloatArray) SetupSuite() {
	this.TBegin("test-fields-floatArray", "")
}

func (this *SuiteFloatArray) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteFloatArray) TestField() {
	target := &FloatArray{}
	assert.Equal(this.T(), []string{"floatArray", "[]float", "float[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeFloatCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeFloatGo, target.ToTypeGo())
}

func (this *SuiteFloatArray) TestToJsonValue() {
	target := &FloatArray{}

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float32{0.123, 0.456, 0.789}, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float32{}, result)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
