package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestShell(t *testing.T) {
	suite.Run(t, new(SuiteShell))
}

type SuiteShell struct {
	suite.Suite
	workDir string
}

func (this *SuiteShell) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteShell) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteShell) TestShellRun() {
	assert.Nil(this.T(), ShellRun("go", "version"))
	assert.NotNil(this.T(), ShellRun("unknown"))
}
