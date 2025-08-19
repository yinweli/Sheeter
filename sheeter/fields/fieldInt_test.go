package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestInt(t *testing.T) {
	suite.Run(t, new(SuiteInt))
}

type SuiteInt struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteInt) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-int")
}

func (this *SuiteInt) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteInt) TestField() {
	target := &Int{}
	assert.Equal(this.T(), []string{"int"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.IsType(this.T(), &Pkey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeIntCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeIntGo, target.ToTypeGo())
}

func (this *SuiteInt) TestToJsonValue() {
	target := &Int{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(123456789), result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(0), result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
