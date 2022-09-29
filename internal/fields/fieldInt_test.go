package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInt(t *testing.T) {
	suite.Run(t, new(SuiteInt))
}

type SuiteInt struct {
	suite.Suite
	workDir string
}

func (this *SuiteInt) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInt) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInt) target() *Int {
	return &Int{}
}

func (this *SuiteInt) TestField() {
	target := this.target()
	assert.Equal(this.T(), "int", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenIntCs, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenIntGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenIntProto, target.ToTypeProto())
}

func (this *SuiteInt) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
