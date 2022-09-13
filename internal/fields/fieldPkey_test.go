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
	workDir string
}

func (this *SuitePkey) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePkey) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
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

func (this *SuitePkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(0), result)

	result, err = target.ToJsonValue("123456789", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue(testdata.UnknownStr, false)
	assert.NotNil(this.T(), err)
}
