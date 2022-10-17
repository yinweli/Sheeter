package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestName(t *testing.T) {
	suite.Run(t, new(SuiteName))
}

type SuiteName struct {
	suite.Suite
	workDir string
}

func (this *SuiteName) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteName) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteName) TestNameCheck() {
	assert.True(this.T(), NameCheck("value"))
	assert.True(this.T(), NameCheck("Value"))
	assert.True(this.T(), NameCheck("value1"))
	assert.True(this.T(), NameCheck("Value1"))
	assert.True(this.T(), NameCheck("value_"))
	assert.True(this.T(), NameCheck("_value"))
	assert.False(this.T(), NameCheck(""))
	assert.False(this.T(), NameCheck("0value"))
	assert.False(this.T(), NameCheck("-value"))
	assert.False(this.T(), NameCheck("value-"))
	assert.False(this.T(), NameCheck("#value"))
	assert.False(this.T(), NameCheck("value#"))
	assert.False(this.T(), NameCheck("@value"))
	assert.False(this.T(), NameCheck("value@"))
	assert.False(this.T(), NameCheck("{value}"))
}

func (this *SuiteName) TestNameKeywords() {
	assert.True(this.T(), NameKeywords("value"))
	assert.False(this.T(), NameKeywords("depot"))
	assert.False(this.T(), NameKeywords("Depot"))
	assert.False(this.T(), NameKeywords("dePot"))
	assert.False(this.T(), NameKeywords("depoT"))
	assert.False(this.T(), NameKeywords("DEPOT"))
}
