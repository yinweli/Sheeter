package utils

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
	testdata.TestEnv
}

func (this *SuiteStrings) SetupSuite() {
	this.Change("test-strings")
}

func (this *SuiteStrings) TearDownSuite() {
	this.Restore()
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
	items := []string{"a", "b", "c"}
	assert.Equal(this.T(), "a", GetItem(items, 0))
	assert.Equal(this.T(), "b", GetItem(items, 1))
	assert.Equal(this.T(), "c", GetItem(items, 2))
	assert.Equal(this.T(), "", GetItem(items, 3))
}
