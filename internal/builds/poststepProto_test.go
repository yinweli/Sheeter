package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepProto(t *testing.T) {
	suite.Run(t, new(SuitePoststepProto))
}

type SuitePoststepProto struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststepProto) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}
