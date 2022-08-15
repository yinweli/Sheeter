package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
}

func (this *SuitePkey) target() *Pkey {
	return &Pkey{}
}

func (this *SuitePkey) TestType() {
	assert.Equal(this.T(), "pkey", this.target().Type())
}

func (this *SuitePkey) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuitePkey) TestIsPkey() {
	assert.Equal(this.T(), true, this.target().IsPkey())
}

func (this *SuitePkey) TestToJsonDefault() {
	assert.Equal(this.T(), int64(0), this.target().ToJsonDefault())
}

func (this *SuitePkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
