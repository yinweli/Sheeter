package excels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestExcel(t *testing.T) {
	suite.Run(t, new(SuiteExcel))
}

type SuiteExcel struct {
	suite.Suite
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteExcel) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = testdata.ExcelNameReal
	this.sheet = testdata.SheetName
}

func (this *SuiteExcel) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteExcel) target() *Excel {
	return &Excel{}
}

func (this *SuiteExcel) TestOpen() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))
	assert.True(this.T(), target.IsOpen())
	target.Close()
	assert.False(this.T(), target.IsOpen())

	target = this.target()
	assert.NotNil(this.T(), target.Open(testdata.UnknownStr))
}

func (this *SuiteExcel) TestGetLine() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))

	line, err := target.GetLine(this.sheet, 1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	line.Close()

	line, err = target.GetLine(this.sheet, 10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	line.Close()

	_, err = target.GetLine(this.sheet, 0)
	assert.NotNil(this.T(), err)

	_, err = target.GetLine(testdata.UnknownStr, 1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteExcel) TestGetData() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))

	data, err := target.GetData(this.sheet, 1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{
			"name0#pkey",
			"empty#empty",
			"name1#bool",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
		}, data)

	_, err = target.GetData(this.sheet, 10)
	assert.NotNil(this.T(), err)

	_, err = target.GetData(this.sheet, 0)
	assert.NotNil(this.T(), err)

	_, err = target.GetData(testdata.UnknownStr, 1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteExcel) TestSheetExist() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))
	assert.True(this.T(), target.SheetExist(this.sheet))
	assert.False(this.T(), target.SheetExist(testdata.UnknownStr))
	target.Close()
}

func (this *SuiteExcel) TestLine() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))
	line, err := target.GetLine(this.sheet, 1)
	assert.Nil(this.T(), err)
	assert.True(this.T(), line.Next())
	data, err := line.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{
		"",
		"",
		"{S",
		"{[]A",
		"",
		"/",
		"",
		"/",
		"}}",
	}, data)
	line.Close()
	target.Close()
}
