package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestVariable(t *testing.T) {
	suite.Run(t, new(SuiteVariable))
}

type SuiteVariable struct {
	suite.Suite
	workDir string
}

func (this *SuiteVariable) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteVariable) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteVariable) TestVariableCheck() {
	assert.True(this.T(), VariableCheck("value"))
	assert.True(this.T(), VariableCheck("Value"))
	assert.True(this.T(), VariableCheck("value1"))
	assert.True(this.T(), VariableCheck("Value1"))
	assert.True(this.T(), VariableCheck("value_"))
	assert.True(this.T(), VariableCheck("_value"))
	assert.False(this.T(), VariableCheck(""))
	assert.False(this.T(), VariableCheck("0value"))
	assert.False(this.T(), VariableCheck("-value"))
	assert.False(this.T(), VariableCheck("value-"))
	assert.False(this.T(), VariableCheck("#value"))
	assert.False(this.T(), VariableCheck("value#"))
	assert.False(this.T(), VariableCheck("@value"))
	assert.False(this.T(), VariableCheck("value@"))
	assert.False(this.T(), VariableCheck("{value}"))
}
