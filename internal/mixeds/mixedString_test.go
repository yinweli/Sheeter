package mixeds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteString) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelString"
	this.sheet = "sheetString"
}

func (this *SuiteString) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteString) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}

func (this *SuiteString) TestString() {
	target := this.target()
	assert.Equal(this.T(), "TestString", target.FirstUpper("testString"))
	assert.Equal(this.T(), "testString", target.FirstLower("TestString"))
}
