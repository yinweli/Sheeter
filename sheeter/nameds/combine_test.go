package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestCombine(t *testing.T) {
	suite.Run(t, new(SuiteCombine))
}

type SuiteCombine struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteCombine) SetupSuite() {
	this.Change("test-combine")
}

func (this *SuiteCombine) TearDownSuite() {
	this.Restore()
}

func (this *SuiteCombine) TestCombine() {
	assert.Equal(this.T(), "ab", combine(&params{
		excelName: "a",
		sheetName: sheeter.SignData + "b",
	}))
	assert.Equal(this.T(), "ab", combine(&params{
		excelName: "a",
		sheetName: sheeter.SignEnum + "b",
	}))
	assert.Equal(this.T(), "ab", combine(&params{
		excelName: "a",
		sheetName: "b",
	}))
	assert.Equal(this.T(), "Ab", combine(&params{
		excelName:  "a",
		excelUpper: true,
		sheetName:  "b",
	}))
	assert.Equal(this.T(), "aB", combine(&params{
		excelName:  "a",
		sheetName:  "b",
		sheetUpper: true,
	}))
	assert.Equal(this.T(), "AB", combine(&params{
		excelName:  "a",
		excelUpper: true,
		sheetName:  "b",
		sheetUpper: true,
	}))
	assert.Equal(this.T(), "abc", combine(&params{
		excelName: "a",
		sheetName: "b",
		last:      "c",
	}))
	assert.Equal(this.T(), "ab.x", combine(&params{
		excelName: "a",
		sheetName: "b",
		ext:       ".x",
	}))
	assert.Equal(this.T(), "abc.x", combine(&params{
		excelName: "a",
		sheetName: "b",
		last:      "c",
		ext:       ".x",
	}))
}
