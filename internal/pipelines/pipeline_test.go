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

func (this *SuitePipeline) TestPipeline() {
	material := []any{0, 1}
	funcs := []PipelineFunc{
		func(material any, result chan any) error {
			result <- material
			return nil
		},
		func(material any, result chan any) error {
			return fmt.Errorf("err")
		},
	}

	result, errs := Pipeline("name", material, funcs)
	assert.Len(this.T(), result, len(material))
	assert.Len(this.T(), errs, len(material))

	result, errs = Pipeline("name", []any{}, []PipelineFunc{})
	assert.Empty(this.T(), result)
	assert.Empty(this.T(), errs)
}
