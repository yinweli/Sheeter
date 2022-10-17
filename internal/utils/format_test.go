package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestFormat(t *testing.T) {
	suite.Run(t, new(SuiteFormat))
}

type SuiteFormat struct {
	suite.Suite
	workDir string
}

func (this *SuiteFormat) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteFormat) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFormat) TestFormat() {
	assert.Nil(this.T(), Format())
}
