package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestGenerate(t *testing.T) {
	suite.Run(t, new(SuiteGenerate))
}

type SuiteGenerate struct {
	suite.Suite
	workDir string
}

func (this *SuiteGenerate) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteGenerate) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}
