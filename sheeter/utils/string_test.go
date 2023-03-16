package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteString) SetupSuite() {
	this.TBegin("test-utils-string", "")
}

func (this *SuiteString) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteString) TestFirstUpper() {
	assert.Equal(this.T(), "", FirstUpper(""))
	assert.Equal(this.T(), "TestString", FirstUpper("testString"))
}

func (this *SuiteString) TestFirstLower() {
	assert.Equal(this.T(), "", FirstLower(""))
	assert.Equal(this.T(), "testString", FirstLower("TestString"))
}

func (this *SuiteString) TestAllSame() {
	assert.Equal(this.T(), true, AllSame(""))
	assert.Equal(this.T(), true, AllSame("aaaaa"))
	assert.Equal(this.T(), false, AllSame("aa1aa"))
}

func (this *SuiteString) TestCombine() {
	assert.ElementsMatch(this.T(), []string{"a", "b", "c", "1", "2", "3"}, Combine([]string{"a", "b", "c"}, []any{"1", "2", "3"}))
}

func (this *SuiteString) TestGetItem() {
	item := []string{"a", "b", "c"}
	assert.Equal(this.T(), "a", GetItem(item, 0))
	assert.Equal(this.T(), "b", GetItem(item, 1))
	assert.Equal(this.T(), "c", GetItem(item, 2))
	assert.Equal(this.T(), "", GetItem(item, 3))
}

func (this *SuiteString) TestGetUnique() {
	item := []string{"a", "b", "c", "a", "b", "c"}
	assert.ElementsMatch(this.T(), []string{"a", "b", "c"}, GetUnique(item))
}
