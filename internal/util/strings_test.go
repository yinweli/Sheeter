package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestStrings(t *testing.T) {
	suite.Run(t, new(SuiteStrings))
}

type SuiteStrings struct {
	suite.Suite
}

func (this *SuiteStrings) TestFirstUpper() {
	assert.Equal(this.T(), "", FirstUpper(""))
	assert.Equal(this.T(), "TestString", FirstUpper("testString"))
}

func (this *SuiteStrings) TestFirstLower() {
	assert.Equal(this.T(), "", FirstLower(""))
	assert.Equal(this.T(), "testString", FirstLower("TestString"))
}
