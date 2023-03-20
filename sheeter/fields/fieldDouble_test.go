package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestDouble(t *testing.T) {
	suite.Run(t, new(SuiteDouble))
}

type SuiteDouble struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteDouble) SetupSuite() {
	this.TBegin("test-fields-double", "")
}

func (this *SuiteDouble) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteDouble) TestField() {
	target := &Double{}
	assert.Equal(this.T(), []string{"double"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Nil(this.T(), target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeDoubleCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeDoubleGo, target.ToTypeGo())
}

func (this *SuiteDouble) TestToJsonValue() {
	target := &Double{}

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float64(0.123456), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float64(0), result)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
