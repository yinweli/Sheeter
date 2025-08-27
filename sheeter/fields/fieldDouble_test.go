package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestDouble(t *testing.T) {
	suite.Run(t, new(SuiteDouble))
}

type SuiteDouble struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteDouble) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-double")
}

func (this *SuiteDouble) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteDouble) TestField() {
	target := &Double{}
	assert.Equal(this.T(), []string{"double"}, target.Field())
	assert.Equal(this.T(), "Double", target.ToTypeCs())
	assert.Equal(this.T(), "float64", target.ToTypeGo())
}

func (this *SuiteDouble) TestToJsonValue() {
	target := &Double{}

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float64(0.123456), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float64(0), result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
