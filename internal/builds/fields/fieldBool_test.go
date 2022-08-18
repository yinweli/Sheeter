package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBool(t *testing.T) {
	suite.Run(t, new(SuiteBool))
}

type SuiteBool struct {
	suite.Suite
}

func (this *SuiteBool) target() *Bool {
	return &Bool{}
}

func (this *SuiteBool) TestType() {
	assert.Equal(this.T(), "bool", this.target().Type())
}

func (this *SuiteBool) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteBool) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteBool) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)
	result, err = target.ToJsonValue("true", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, result)
	result, err = target.ToJsonValue("false", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)
	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
