package mixeds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestMath(t *testing.T) {
	suite.Run(t, new(SuiteMath))
}

type SuiteMath struct {
	suite.Suite
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteMath) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelMath"
	this.sheet = "sheetMath"
}

func (this *SuiteMath) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteMath) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}

func (this *SuiteMath) TestMath() {
	target := this.target()
	assert.Equal(this.T(), "8", target.Add(6, 2))
	assert.Equal(this.T(), "8", target.Add(2, 6))
	assert.Equal(this.T(), "4", target.Sub(6, 2))
	assert.Equal(this.T(), "-4", target.Sub(2, 6))
	assert.Equal(this.T(), "12", target.Mul(6, 2))
	assert.Equal(this.T(), "12", target.Mul(2, 6))
	assert.Equal(this.T(), "3", target.Div(6, 2))
	assert.Equal(this.T(), "0", target.Div(2, 6))
}
