package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestName(t *testing.T) {
	suite.Run(t, new(SuiteName))
}

type SuiteName struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteName) SetupSuite() {
	this.Change("test-name")
}

func (this *SuiteName) TearDownSuite() {
	this.Restore()
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

func (this *SuiteName) TestIsDataSheetName() {
	input := "sheetData"
	assert.True(this.T(), IsDataSheetName(input))
	assert.True(this.T(), IsDataSheetName(sheeter.SignData+input))
	assert.False(this.T(), IsDataSheetName(sheeter.SignEnum+input))
	assert.False(this.T(), IsDataSheetName(sheeter.SignIgnore+input))
}

func (this *SuiteName) TestIsEnumSheetName() {
	input := "sheetEnum"
	assert.True(this.T(), IsEnumSheetName(sheeter.SignEnum+input))
	assert.False(this.T(), IsIgnoreSheetName(input))
}

func (this *SuiteName) TestIsIgnoreSheetName() {
	input := "sheetIgnore"
	assert.True(this.T(), IsIgnoreSheetName(sheeter.SignIgnore+input))
	assert.True(this.T(), IsIgnoreSheetName(strings.ToUpper(sheeter.SignIgnore+input)))
	assert.True(this.T(), IsIgnoreSheetName(strings.ToLower(sheeter.SignIgnore+input)))
	assert.False(this.T(), IsIgnoreSheetName(input))
}

func (this *SuiteName) TestRemoveSheetPrefix() {
	input := "sheetRemove"
	assert.Equal(this.T(), input, RemoveSheetPrefix(sheeter.SignData+input))
	assert.Equal(this.T(), input, RemoveSheetPrefix(sheeter.SignEnum+input))
	assert.Equal(this.T(), input, RemoveSheetPrefix(input))
}
