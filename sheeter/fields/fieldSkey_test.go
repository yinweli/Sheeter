package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestSkey(t *testing.T) {
	suite.Run(t, new(SuiteSkey))
}

type SuiteSkey struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteSkey) SetupSuite() {
	this.Change("test-field-skey")
}

func (this *SuiteSkey) TearDownSuite() {
	this.Restore()
}

func (this *SuiteSkey) target() *Skey {
	return &Skey{}
}

func (this *SuiteSkey) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"skey"}, target.Field())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TokenSkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TokenSkeyGo, target.ToTypeGo())
	assert.Equal(this.T(), sheeter.TokenSkeyProto, target.ToTypeProto())
}

func (this *SuiteSkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
