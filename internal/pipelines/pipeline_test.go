package pipelines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestPipeline(t *testing.T) {
	suite.Run(t, new(SuitePipeline))
}

type SuitePipeline struct {
	suite.Suite
	workDir string
}

func (this *SuitePipeline) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePipeline) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePipeline) TestExecute() {
	material := []any{0, 1}
	executor := []Executor{
		func(data any) error {
			return fmt.Errorf("err")
		},
		func(data any) error {
			return fmt.Errorf("err")
		},
	}
	assert.Len(this.T(), Execute("name", material, executor), len(material)*len(executor))
}
