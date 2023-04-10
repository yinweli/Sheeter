package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
	testdata.Env
}

func (this *SuitePkey) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-pkey")
}

func (this *SuitePkey) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuitePkey) TestField() {
	target := &Pkey{}
	assert.Equal(this.T(), []string{"pkey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.IsType(this.T(), &Pkey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.ToTypeGo())
}

func (this *SuitePkey) TestToJsonValue() {
	target := &Pkey{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(123456789), result)
}
