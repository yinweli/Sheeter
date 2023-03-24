package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteField) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-fields-parser")
}

func (this *SuiteField) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteField) TestParser() {
	field, err := Parser("boolArray")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), (&BoolArray{}).Field(), field.Field())

	field, err = Parser("[]bool")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), (&BoolArray{}).Field(), field.Field())

	field, err = Parser("bool[]")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), (&BoolArray{}).Field(), field.Field())

	_, err = Parser(testdata.Unknown)
	assert.NotNil(this.T(), err)
}
