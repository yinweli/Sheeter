package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestCheck(t *testing.T) {
	suite.Run(t, new(SuiteCheck))
}

type SuiteCheck struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteCheck) SetupSuite() {
	this.Env = testdata.EnvSetup("test-utils-check")
}

func (this *SuiteCheck) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteCheck) TestCheckIgnore() {
	assert.True(this.T(), CheckIgnore(sheeter.TokenIgnore+"data"))
	assert.True(this.T(), CheckIgnore("data"+sheeter.TokenIgnore))
	assert.True(this.T(), CheckIgnore("da"+sheeter.TokenIgnore+"ta"))
	assert.False(this.T(), CheckIgnore(testdata.Unknown))
}

func (this *SuiteCheck) TestCheckExcel() {
	assert.True(this.T(), CheckExcel("value"))
	assert.True(this.T(), CheckExcel("Value"))
	assert.True(this.T(), CheckExcel("value1"))
	assert.True(this.T(), CheckExcel("Value1"))
	assert.True(this.T(), CheckExcel("value_"))
	assert.True(this.T(), CheckExcel("_value"))
	assert.False(this.T(), CheckExcel(""))
	assert.False(this.T(), CheckExcel("0value"))
	assert.False(this.T(), CheckExcel("-value"))
	assert.False(this.T(), CheckExcel("value-"))
	assert.False(this.T(), CheckExcel("#value"))
	assert.False(this.T(), CheckExcel("value#"))
	assert.False(this.T(), CheckExcel("@value"))
	assert.False(this.T(), CheckExcel("value@"))
	assert.False(this.T(), CheckExcel("{value}"))
}

func (this *SuiteCheck) TestCheckSheet() {
	assert.True(this.T(), CheckSheet("value"))
	assert.True(this.T(), CheckSheet("Value"))
	assert.True(this.T(), CheckSheet("value1"))
	assert.True(this.T(), CheckSheet("Value1"))
	assert.True(this.T(), CheckSheet("value_"))
	assert.True(this.T(), CheckSheet("_value"))
	assert.True(this.T(), CheckSheet("0value"))
	assert.False(this.T(), CheckSheet(""))
	assert.False(this.T(), CheckSheet("-value"))
	assert.False(this.T(), CheckSheet("value-"))
	assert.False(this.T(), CheckSheet("#value"))
	assert.False(this.T(), CheckSheet("value#"))
	assert.False(this.T(), CheckSheet("@value"))
	assert.False(this.T(), CheckSheet("value@"))
	assert.False(this.T(), CheckSheet("{value}"))
}

func (this *SuiteCheck) TestCheckField() {
	assert.True(this.T(), CheckField("value"))
	assert.True(this.T(), CheckField("Value"))
	assert.True(this.T(), CheckField("value1"))
	assert.True(this.T(), CheckField("Value1"))
	assert.True(this.T(), CheckField("value_"))
	assert.True(this.T(), CheckField("_value"))
	assert.False(this.T(), CheckField(""))
	assert.False(this.T(), CheckField("0value"))
	assert.False(this.T(), CheckField("-value"))
	assert.False(this.T(), CheckField("value-"))
	assert.False(this.T(), CheckField("#value"))
	assert.False(this.T(), CheckField("value#"))
	assert.False(this.T(), CheckField("@value"))
	assert.False(this.T(), CheckField("value@"))
	assert.False(this.T(), CheckField("{value}"))
}

func (this *SuiteCheck) TestCheckTag() {
	assert.True(this.T(), CheckTag("a", "abc"))
	assert.True(this.T(), CheckTag("b", "abc"))
	assert.True(this.T(), CheckTag("c", "abc"))
	assert.True(this.T(), CheckTag("ab", "abc"))
	assert.True(this.T(), CheckTag("bc", "abc"))
	assert.True(this.T(), CheckTag("ac", "abc"))
	assert.False(this.T(), CheckTag("x", "abc"))
	assert.False(this.T(), CheckTag("i", sheeter.TokenIgnore))
}
