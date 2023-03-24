package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestStringArray(t *testing.T) {
	suite.Run(t, new(SuiteStringArray))
}

type SuiteStringArray struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteStringArray) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-fields-stringArray")
}

func (this *SuiteStringArray) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteStringArray) TestField() {
	target := &StringArray{}
	assert.Equal(this.T(), []string{"stringArray", "[]string", "string[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeStringCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeStringGo, target.ToTypeGo())
}

func (this *SuiteStringArray) TestToJsonValue() {
	target := &StringArray{}

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, result)
}
