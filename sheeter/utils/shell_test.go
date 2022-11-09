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
	testdata.TestEnv
}

func (this *SuiteShell) SetupSuite() {
	this.Change("test-shell")
}

func (this *SuiteShell) TearDownSuite() {
	this.Restore()
}

func (this *SuiteShell) TestShellRun() {
	assert.Nil(this.T(), ShellRun("go", "version"))
	assert.NotNil(this.T(), ShellRun("unknown"))
}
