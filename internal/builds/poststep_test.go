package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststep(t *testing.T) {
	suite.Run(t, new(SuitePoststep))
}

type SuitePoststep struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststep) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststep) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}
