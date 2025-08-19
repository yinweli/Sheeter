package pipelines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestPipeline(t *testing.T) {
	suite.Run(t, new(SuitePipeline))
}

type SuitePipeline struct {
	suite.Suite
	testdata.Env
}

func (this *SuitePipeline) SetupSuite() {
	this.Env = testdata.EnvSetup("test-pipelines-pipeline")
}

func (this *SuitePipeline) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuitePipeline) TestPipeline() {
	result, err := Pipeline[int]("name", []int{0, 1}, []PipelineFunc[int]{
		func(material int, result chan any) error {
			result <- material
			return nil
		},
		func(material int, result chan any) error {
			return fmt.Errorf("err")
		},
	})
	assert.Len(this.T(), result, 2)
	assert.Len(this.T(), err, 2)

	result, err = Pipeline[int]("name", []int{}, []PipelineFunc[int]{})
	assert.Empty(this.T(), result)
	assert.Empty(this.T(), err)
}
