package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestShell(t *testing.T) {
	suite.Run(t, new(SuiteShell))
}

type SuiteShell struct {
	suite.Suite
}

func (this *SuiteShell) TestShellRun() {
	assert.Nil(this.T(), ShellRun("go", "version"))
	assert.NotNil(this.T(), ShellRun("unknown"))
}

func (this *SuiteShell) TestShellExist() {
	assert.True(this.T(), ShellExist("go"))
	assert.False(this.T(), ShellExist("unknown"))
}
