package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuitePkey) SetupSuite() {
	this.Change("test-field-pkey")
}

func (this *SuitePkey) TearDownSuite() {
	this.Restore()
}

func (this *SuitePkey) target() *Pkey {
	return &Pkey{}
}

func (this *SuitePkey) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"pkey"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenPkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenPkeyGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenPkeyProto, target.ToTypeProto())
}

func (this *SuitePkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
