package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestTerm(t *testing.T) {
	suite.Run(t, new(SuiteTerm))
}

type SuiteTerm struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteTerm) SetupSuite() {
	this.Env = testdata.EnvSetup("test-utils-term")
}

func (this *SuiteTerm) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteTerm) TestMergeTerm() {
	target := MergeTerm("name$excel1#sheet&excel2#sheet")
	assert.Equal(this.T(), "name", target.Name())
	assert.Equal(this.T(), []SheetTerm{"excel1#sheet", "excel2#sheet"}, target.Member())
	target = "excel1#sheet&excel2#sheet"
	assert.Equal(this.T(), "", target.Name())
	assert.Equal(this.T(), []SheetTerm{}, target.Member())
}

func (this *SuiteTerm) TestSheetTerm() {
	target := SheetTerm("excel#sheet")
	assert.True(this.T(), target.Match("excel", "sheet"))
	assert.False(this.T(), target.Match(testdata.Unknown, "sheet"))
	assert.False(this.T(), target.Match("excel", testdata.Unknown))
	assert.False(this.T(), target.Match("", ""))
}
