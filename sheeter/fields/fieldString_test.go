package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteString) SetupSuite() {
	this.Change("test-field-string")
}

func (this *SuiteString) TearDownSuite() {
	this.Restore()
}

func (this *SuiteString) target() *String {
	return &String{}
}

func (this *SuiteString) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"string"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenStringCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenStringGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenStringProto, target.ToTypeProto())
}

func (this *SuiteString) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
