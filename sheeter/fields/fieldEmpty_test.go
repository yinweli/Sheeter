package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestEmpty(t *testing.T) {
	suite.Run(t, new(SuiteEmpty))
}

type SuiteEmpty struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteEmpty) SetupSuite() {
	this.Change("test-field-empty")
}

func (this *SuiteEmpty) TearDownSuite() {
	this.Restore()
}

func (this *SuiteEmpty) target() *Empty {
	return &Empty{}
}

func (this *SuiteEmpty) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"empty", "ignore"}, target.Field())
	assert.Equal(this.T(), false, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), "", target.ToTypeCs())
	assert.Equal(this.T(), "", target.ToTypeGo())
	assert.Equal(this.T(), "", target.ToTypeProto())
}

func (this *SuiteEmpty) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("test")
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), result)
}
