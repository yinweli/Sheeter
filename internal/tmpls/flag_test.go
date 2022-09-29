package tmpls

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
	workDir string
}

func (this *SuiteFlag) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteFlag) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFlag) TestSetFlags() {
	cmd := SetFlags(&cobra.Command{})
	assert.NotNil(this.T(), cmd)
	assert.NotNil(this.T(), cmd.Flags().Lookup(flagClean))
}