package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepFormat(t *testing.T) {
	suite.Run(t, new(SuitePoststepFormat))
}

type SuitePoststepFormat struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepFormat) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststepFormat) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepFormat) target() *poststepFormat {
	return &poststepFormat{}
}

func (this *SuitePoststepFormat) TestPoststepFormat() {
	target := this.target()
	assert.Nil(this.T(), PoststepFormatCs(target))
	assert.Nil(this.T(), PoststepFormatGo(target))
	assert.Nil(this.T(), PoststepFormatProto(target))
	assert.Nil(this.T(), PoststepFormatCs(nil))
	assert.Nil(this.T(), PoststepFormatGo(nil))
	assert.Nil(this.T(), PoststepFormatProto(nil))
}
