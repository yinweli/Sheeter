package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
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

func (this *SuitePkey) TestField() {
	target := this.target()
	assert.Equal(this.T(), "pkey", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), internal.TokenIntCs, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenIntGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenIntProto, target.ToTypeProto())
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
