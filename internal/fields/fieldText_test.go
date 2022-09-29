package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestText(t *testing.T) {
	suite.Run(t, new(SuiteText))
}

type SuiteText struct {
	suite.Suite
	workDir string
}

func (this *SuiteText) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteText) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteText) target() *Text {
	return &Text{}
}

func (this *SuiteText) TestField() {
	target := this.target()
	assert.Equal(this.T(), "text", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenStringCs, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenStringGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenStringProto, target.ToTypeProto())
}

func (this *SuiteText) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
