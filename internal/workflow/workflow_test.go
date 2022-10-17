package workflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestWorkflow(t *testing.T) {
	suite.Run(t, new(SuiteWorkflow))
}

type SuiteWorkflow struct {
	suite.Suite
	workDir string
	count   int
}

func (this *SuiteWorkflow) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.count = 5
}

func (this *SuiteWorkflow) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteWorkflow) target() *Workflow {
	return NewWorkflow("", this.count)
}

func (this *SuiteWorkflow) TestNewWorkflow() {
	assert.NotNil(this.T(), NewWorkflow("", this.count))
}

func (this *SuiteWorkflow) TestIncrement() {
	target := this.target()

	for i := 0; i < this.count; i++ {
		target.Error(fmt.Errorf("err%v", i))
		target.Increment()
	} // for

	errs := target.End()
	assert.Len(this.T(), errs, this.count)
}

func (this *SuiteWorkflow) TestAbort() {
	target := this.target()
	target.Error(fmt.Errorf("err"))
	target.Increment()
	target.Abort()
	errs := target.End()
	assert.Len(this.T(), errs, 1)
}
