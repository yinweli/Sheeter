package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestProto(t *testing.T) {
	suite.Run(t, new(SuiteProto))
}

type SuiteProto struct {
	suite.Suite
	workDir string
}

func (this *SuiteProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteProto) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteProto) TestParseProto() {
	file, err := ParseProto(testdata.TestProto)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), file)
	assert.NotNil(this.T(), file.FindMessage("sheeter.Test"))
	assert.NotNil(this.T(), file.FindMessage("sheeter.TestReader"))

	_, err = ParseProto(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
