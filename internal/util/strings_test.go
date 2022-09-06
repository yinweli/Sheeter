package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestStrings(t *testing.T) {
	suite.Run(t, new(SuiteStrings))
}

type SuiteStrings struct {
	suite.Suite
	workDir string
	items   []string
}

func (this *SuiteStrings) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.items = []string{"a", "b", "c"}
}

func (this *SuiteStrings) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteStrings) TestFirstUpper() {
	assert.Equal(this.T(), "", FirstUpper(""))
	assert.Equal(this.T(), "TestString", FirstUpper("testString"))
}

func (this *SuiteStrings) TestFirstLower() {
	assert.Equal(this.T(), "", FirstLower(""))
	assert.Equal(this.T(), "testString", FirstLower("TestString"))
}

func (this *SuiteStrings) TestAllSame() {
	assert.Equal(this.T(), true, AllSame(""))
	assert.Equal(this.T(), true, AllSame("aaaaa"))
	assert.Equal(this.T(), false, AllSame("aa1aa"))
}

func (this *SuiteStrings) TestGetItem() {
	assert.Equal(this.T(), "a", GetItem(this.items, 0))
	assert.Equal(this.T(), "b", GetItem(this.items, 1))
	assert.Equal(this.T(), "c", GetItem(this.items, 2))
	assert.Equal(this.T(), "", GetItem(this.items, 3))
}
