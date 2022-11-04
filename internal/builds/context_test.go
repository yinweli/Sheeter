package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestContext(t *testing.T) {
	suite.Run(t, new(SuiteContext))
}

type SuiteContext struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteContext) SetupSuite() {
	this.Change("test-context")
}

func (this *SuiteContext) TearDownSuite() {
	this.Restore()
}
