package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestStringArray(t *testing.T) {
	suite.Run(t, new(SuiteStringArray))
}

type SuiteStringArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteStringArray) SetupSuite() {
	this.Change("test-field-stringArray")
}

func (this *SuiteStringArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteStringArray) target() *StringArray {
	return &StringArray{}
}

func (this *SuiteStringArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "stringArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenStringCs+sheeter.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenArray+sheeter.TokenStringGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenRepeated+" "+sheeter.TokenStringProto, target.ToTypeProto())
}

func (this *SuiteStringArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, result)
}
