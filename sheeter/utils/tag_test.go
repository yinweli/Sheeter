package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestTag(t *testing.T) {
	suite.Run(t, new(SuiteTag))
}

type SuiteTag struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteTag) SetupSuite() {
	this.Change("test-tag")
}

func (this *SuiteTag) TearDownSuite() {
	this.Restore()
}

func (this *SuiteTag) TestTagMatch() {
	assert.True(this.T(), TagMatch("a", "abc"))
	assert.True(this.T(), TagMatch("b", "abc"))
	assert.True(this.T(), TagMatch("c", "abc"))
	assert.True(this.T(), TagMatch("ab", "abc"))
	assert.True(this.T(), TagMatch("bc", "abc"))
	assert.True(this.T(), TagMatch("ac", "abc"))
	assert.False(this.T(), TagMatch("x", "abc"))
}
