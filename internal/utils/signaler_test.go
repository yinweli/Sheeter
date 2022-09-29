package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestSignaler(t *testing.T) {
	suite.Run(t, new(SuiteSignaler))
}

type SuiteSignaler struct {
	suite.Suite
	workDir string
}

func (this *SuiteSignaler) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteSignaler) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSignaler) TestNewWaitGroup() {
	assert.NotNil(this.T(), NewWaitGroup(0))
}
