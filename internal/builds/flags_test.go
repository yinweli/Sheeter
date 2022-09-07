package builds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestFlags(t *testing.T) {
	suite.Run(t, new(SuiteFlags))
}

type SuiteFlags struct {
	suite.Suite
	workDir string
}

func (this *SuiteFlags) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteFlags) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFlags) TestInitFlags() {
	cmd := InitFlags(&cobra.Command{})
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagConfig))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagBom))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfField))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfLayer))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfNote))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfData))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagElements))
}
