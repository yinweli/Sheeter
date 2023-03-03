package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteField) SetupSuite() {
	this.Change("test-field-parser")
}

func (this *SuiteField) TearDownSuite() {
	this.Restore()
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

	_, err = Parser(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
