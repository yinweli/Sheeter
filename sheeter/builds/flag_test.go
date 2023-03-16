package builds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestFlag(t *testing.T) {
	suite.Run(t, new(SuiteFlag))
}

type SuiteFlag struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteFlag) SetupSuite() {
	this.TBegin("test-builds-flag", "")
}

func (this *SuiteFlag) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteFlag) TestSetFlag() {
	cmd := SetFlag(&cobra.Command{})
	assert.NotNil(this.T(), cmd)
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagConfig))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagSource))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagOutput))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagExclude))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagTag))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagAutoKey))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfTag))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfName))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfNote))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfField))
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagLineOfData))
}
