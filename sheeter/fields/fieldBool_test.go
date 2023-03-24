package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestBool(t *testing.T) {
	suite.Run(t, new(SuiteBool))
}

type SuiteBool struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteBool) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-fields-bool")
}

func (this *SuiteBool) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteBool) TestField() {
	target := &Bool{}
	assert.Equal(this.T(), []string{"bool"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.IsType(this.T(), &Pkey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeBoolCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeBoolGo, target.ToTypeGo())
}

func (this *SuiteBool) TestToJsonValue() {
	target := &Bool{}

	result, err := target.ToJsonValue("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, result)

	result, err = target.ToJsonValue("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)

	result, err = target.ToJsonValue("")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)

	_, err = target.ToJsonValue(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
