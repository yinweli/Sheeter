package builds

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestContext(t *testing.T) {
	suite.Run(t, new(SuiteContext))
}

type SuiteContext struct {
	suite.Suite
	workDir string
}

func (this *SuiteContext) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteContext) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteContext) target() *Context {
	target := &Context{
		Sector: []*ContextSector{
			{excel: &excels.Excel{}},
		},
	}
	return target
}

func (this *SuiteContext) TestClose() {
	this.target().Close()
}