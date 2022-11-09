package excels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestExcel(t *testing.T) {
	suite.Run(t, new(SuiteExcel))
}

type SuiteExcel struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteExcel) SetupSuite() {
	this.Change("test-excel")
}

func (this *SuiteExcel) TearDownSuite() {
	CloseAll()
	this.Restore()
}

func (this *SuiteExcel) target() *Excel {
	return &Excel{}
}

func (this *SuiteExcel) TestOpen() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelReal))
	assert.True(this.T(), target.IsOpen())
	target.Close()
	assert.False(this.T(), target.IsOpen())

	target = this.target()
	assert.NotNil(this.T(), target.Open(testdata.UnknownStr))

	CloseAll()
}

func (this *SuiteExcel) TestGet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelReal))

	sheet, err := target.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	_, err = target.Get(testdata.UnknownStr)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestGetLine() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelReal))

	line, err := target.GetLine(testdata.SheetData, 1, 2)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 2)
	assert.Equal(this.T(), []string{
		"name0",
		"empty",
		"name1",
		"name2",
		"name3",
		"name2",
		"name3",
		"name2",
		"name3",
	}, line[1])
	assert.Equal(this.T(), []string{
		"note0",
		"empty",
		"note1",
		"note2",
		"note3",
		"note4",
		"note5",
		"note6",
		"note7",
	}, line[2])

	_, err = target.GetLine(testdata.SheetData, -1)
	assert.NotNil(this.T(), err)

	_, err = target.GetLine(testdata.UnknownStr, 1)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestSheets() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelReal))
	assert.Equal(this.T(), []string{sheeter.SignData + "Data", sheeter.SignEnum + "Enum"}, target.Sheets())
	CloseAll()
}

func (this *SuiteExcel) TestExist() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelReal))
	assert.True(this.T(), target.Exist(testdata.SheetData))
	assert.False(this.T(), target.Exist(testdata.UnknownStr))
	CloseAll()
}

func (this *SuiteExcel) TestSheet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelSheet))

	sheet, err := target.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)

	assert.True(this.T(), sheet.Next())
	data, err := sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{
			"name0",
			"name1",
			"name2",
			"name3",
			"empty",
		}, data)

	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), data)

	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{
		"pkey",
		"bool",
		"int",
		"string",
		"empty",
	}, data)
	sheet.Close()

	sheet, err = target.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	assert.True(this.T(), sheet.Nextn(2))
	assert.False(this.T(), sheet.Nextn(-1))
	sheet.Close()

	sheet, err = target.Get(testdata.SheetData)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	_, err = sheet.Data()
	assert.NotNil(this.T(), err)
	sheet.Close()

	CloseAll()
}

func (this *SuiteExcel) TestColumnToIndex() {
	assert.Equal(this.T(), 1, columnToIndex("A"))
	assert.Equal(this.T(), 111, columnToIndex("DG"))
	assert.Equal(this.T(), 222, columnToIndex("HN"))
	assert.Equal(this.T(), 333, columnToIndex("LU"))
	assert.Equal(this.T(), 444, columnToIndex("QB"))
	assert.Equal(this.T(), 555, columnToIndex("UI"))
	assert.Equal(this.T(), 666, columnToIndex("YP"))
	assert.Equal(this.T(), 777, columnToIndex("ACW"))
	assert.Panics(this.T(), func() { columnToIndex("") })
	assert.Panics(this.T(), func() { columnToIndex("0") })
	assert.Panics(this.T(), func() { columnToIndex("?") })
}
