package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestEncoding(t *testing.T) {
	suite.Run(t, new(SuiteEncoding))
}

type SuiteEncoding struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncoding) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncoding) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}
