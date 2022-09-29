package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTextArray(t *testing.T) {
	suite.Run(t, new(SuiteTextArray))
}

type SuiteTextArray struct {
	suite.Suite
	workDir string
}

func (this *SuiteTextArray) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteTextArray) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTextArray) target() *TextArray {
	return &TextArray{}
}

func (this *SuiteTextArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), "textArray", target.Type())
	assert.Equal(this.T(), true, target.IsShow())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), internal.TokenStringCs+internal.TokenArray, target.ToTypeCs())
	assert.Equal(this.T(), internal.TokenArray+internal.TokenStringGo, target.ToTypeGo())
	assert.Equal(this.T(), internal.TokenRepeated+" "+internal.TokenStringProto, target.ToTypeProto())
}

func (this *SuiteTextArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack", true)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{}, result)

	result, err = target.ToJsonValue("ball,book,pack", false)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, result)
}
