package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestBool(t *testing.T) {
	suite.Run(t, new(SuiteBool))
}

type SuiteBool struct {
	suite.Suite
	workDir string
}

func (this *SuiteBool) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteBool) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBool) target() *Bool {
	return &Bool{}
}

func (this *SuiteBool) TestField() {
	target := this.target()
	assert.Equal(this.T(), "bool", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenBoolCs, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenBoolGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenBoolProto, target.ToTypeProto())
}

func (this *SuiteBool) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, result)

	result, err = target.ToJsonValue("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
