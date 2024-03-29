package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteString) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-string")
}

func (this *SuiteString) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteString) TestField() {
	target := &String{}
	assert.Equal(this.T(), []string{"string"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.IsType(this.T(), &Skey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeStringCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeStringGo, target.ToTypeGo())
}

func (this *SuiteString) TestToJsonValue() {
	target := &String{}

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
