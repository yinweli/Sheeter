package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestVersion(t *testing.T) {
	suite.Run(t, new(SuiteVersion))
}

type SuiteVersion struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteVersion) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-cmd-version")
}

func (this *SuiteVersion) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteVersion) TestNewCommand() {
	assert.NotNil(this.T(), NewCommand())
}

func (this *SuiteVersion) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Execute())
}
