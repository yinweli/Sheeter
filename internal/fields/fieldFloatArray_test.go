package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFloatArray))
}

type SuiteFloatArray struct {
	suite.Suite
	workDir string
}

func (this *SuiteFloatArray) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteFloatArray) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFloatArray) target() *FloatArray {
	return &FloatArray{}
}

func (this *SuiteFloatArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "floatArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenFloatCs+internal.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenArray+internal.TokenFloatGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenRepeated+" "+internal.TokenFloatProto, target.ToTypeProto())
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
