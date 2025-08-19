package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestLkey(t *testing.T) {
	suite.Run(t, new(SuiteLkey))
}

type SuiteLkey struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteLkey) SetupSuite() {
	this.Env = testdata.EnvSetup("test-fields-lkey")
}

func (this *SuiteLkey) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteLkey) TestField() {
	target := &Lkey{}
	assert.Equal(this.T(), []string{"lkey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.IsType(this.T(), &Lkey{}, target.ToPkey())
	assert.Equal(this.T(), sheeter.TypeLkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeLkeyGo, target.ToTypeGo())
}

func (this *SuiteLkey) TestToJsonValue() {
	target := &Lkey{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)
}
