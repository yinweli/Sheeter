package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteBoolArray))
}

type SuiteBoolArray struct {
	suite.Suite
	workDir string
}

func (this *SuiteBoolArray) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteBoolArray) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBoolArray) target() *BoolArray {
	return &BoolArray{}
}

func (this *SuiteBoolArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "boolArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenBool+internal.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenArray+internal.TokenBool, target.ToTypeGo())
}

func (this *SuiteBoolArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{}, result)

	result, err = target.ToJsonValue("true,false,true,false,true", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
