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

func (this *SuiteExcel) TestGet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))

	sheet, err := target.Get(this.sheet)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	_, err = target.Get(testdata.UnknownStr)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteExcel) TestGetLine() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))

	line, err := target.GetLine(this.sheet, 1, 2)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 2)
	assert.Equal(this.T(), []string{
		"name0#pkey",
		"empty#empty",
		"name1#bool",
		"name2#int",
		"name3#text",
		"name2#int",
		"name3#text",
		"name2#int",
		"name3#text",
	}, line[1])
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
	}, line[2])

	_, err = target.GetLine(this.sheet, 100)
	assert.NotNil(this.T(), err)

	_, err = target.GetLine(testdata.UnknownStr, 1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteExcel) TestExist() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))
	assert.True(this.T(), target.Exist(this.sheet))
	assert.False(this.T(), target.Exist(testdata.UnknownStr))
	target.Close()
}

func (this *SuiteExcel) TestSheet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(this.excel))
	sheet, err := target.Get(this.sheet)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)

	assert.True(this.T(), sheet.Next())
	assert.True(this.T(), sheet.Nextn(2))
	assert.False(this.T(), sheet.Nextn(-1))
	data, err := sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{
			"note0",
			"empty",
			"note1",
			"note2",
			"note3",
			"note4",
			"note5",
			"note6",
			"note7",
		}, data)

	sheet.Close()
	target.Close()
}
