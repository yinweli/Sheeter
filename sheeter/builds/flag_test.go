package builds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestFlag(t *testing.T) {
	suite.Run(t, new(SuiteFlag))
}

type SuiteFlag struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteFlag) SetupSuite() {
	this.Change("test-build-flag")
}

func (this *SuiteFlag) TearDownSuite() {
	this.Restore()
}

func (this *SuiteFlag) TestSetFlags() {
	cmd := SetFlags(&cobra.Command{})
	assert.NotNil(this.T(), cmd)
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagConfig))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagExportJson))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagExportProto))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagExportEnum))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagSimpleNamespace))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfTag))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfName))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfNote))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfField))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfLayer))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfData))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfEnum))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagTags))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagInputs))
}
