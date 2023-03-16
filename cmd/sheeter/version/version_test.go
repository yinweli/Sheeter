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
	testdata.TestData
}

func (this *SuiteVersion) SetupSuite() {
	this.TBegin("test-cmd-version", "")
}

func (this *SuiteVersion) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteVersion) TestNewCommand() {
	assert.NotNil(this.T(), NewCommand())
}

func (this *SuiteVersion) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Execute())
}
