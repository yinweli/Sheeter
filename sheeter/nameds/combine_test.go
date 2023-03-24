package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestCombine(t *testing.T) {
	suite.Run(t, new(SuiteCombine))
}

type SuiteCombine struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteCombine) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-nameds-combine")
}

func (this *SuiteCombine) TearDownSuite() {
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteCombine) TestCombine() {
	assert.Equal(this.T(), "aB", combine(&params{
		excelName: "a",
		sheetName: "b",
	}))
	assert.Equal(this.T(), "a", combine(&params{
		excelName: "a",
		sheetName: "a",
	}))
	assert.Equal(this.T(), "AB", combine(&params{
		excelUpper: true,
		excelName:  "a",
		sheetName:  "b",
	}))
	assert.Equal(this.T(), "aBc", combine(&params{
		excelName: "a",
		sheetName: "b",
		last:      "c",
	}))
	assert.Equal(this.T(), "aB.x", combine(&params{
		excelName: "a",
		sheetName: "b",
		ext:       ".x",
	}))
	assert.Equal(this.T(), "aBc.x", combine(&params{
		excelName: "a",
		sheetName: "b",
		last:      "c",
		ext:       ".x",
	}))
}
