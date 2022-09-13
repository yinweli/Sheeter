package codes

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
	workDir string
}

func (this *SuiteConfig) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteConfig) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteConfig) TestSetFlags() {
	cmd := SetFlags(&cobra.Command{})
	assert.NotNil(this.T(), cmd)
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagClean))
}
