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
	result, err := Pipeline[int]("name1", []int{0, 1}, []Execute[int]{
		func(material int) Output {
			return Output{
				Result: []any{material},
			}
		},
		func(material int) Output {
			return Output{
				Error: fmt.Errorf("error"),
			}
		},
	})
	assert.Len(this.T(), result, 2)
	assert.Len(this.T(), err, 2)

	result, err = Pipeline[int]("name2", nil, nil)
	assert.Empty(this.T(), result)
	assert.Empty(this.T(), err)
}
