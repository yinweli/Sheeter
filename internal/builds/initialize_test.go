package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitialize) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitialize) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}
