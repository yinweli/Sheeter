package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestDouble(t *testing.T) {
	suite.Run(t, new(SuiteDouble))
}

type SuiteDouble struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteDouble) SetupSuite() {
	this.Change("test-field-double")
}

func (this *SuiteDouble) TearDownSuite() {
	this.Restore()
}

func (this *SuiteDouble) target() *Double {
	return &Double{}
}

func (this *SuiteDouble) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"double"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenDoubleCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenDoubleGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenDoubleProto, target.ToTypeProto())
}

func (this *SuiteDouble) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.123456, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
