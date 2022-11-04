package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestVersion(t *testing.T) {
	suite.Run(t, new(SuiteVersion))
}

type SuiteVersion struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteVersion) SetupSuite() {
	this.Change("test-cmd-version")
}

func (this *SuiteVersion) TearDownSuite() {
	this.Restore()
}

func (this *SuiteVersion) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Execute())
}
