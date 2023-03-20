package excels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestExcel(t *testing.T) {
	suite.Run(t, new(SuiteExcel))
}

type SuiteExcel struct {
	suite.Suite
	testdata.TestEnv
	excelSuccess string
	sheet1       string
	sheet2       string
}

func (this *SuiteExcel) SetupSuite() {
	this.TBegin("test-excels-excel", "excel")
	this.excelSuccess = "success.xlsx"
	this.sheet1 = "Test1"
	this.sheet2 = "Test2"
}

func (this *SuiteExcel) TearDownSuite() {
	CloseAll()
	this.TFinal()
}

func (this *SuiteExcel) TestOpen() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	assert.True(this.T(), target.IsOpen())
	target.Close()
	assert.False(this.T(), target.IsOpen())

	target = &Excel{}
	assert.NotNil(this.T(), target.Open(this.Unknown))

	CloseAll()
}

func (this *SuiteExcel) TestGet() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))

	sheet, err := target.Get(this.sheet1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	sheet, err = target.Get(this.sheet2)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	_, err = target.Get(this.Unknown)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestGetLine() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))

	line, err := target.GetLine(this.sheet1, 1, 2, 3)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 3)
	assert.Equal(this.T(), []string{"value1", "value2"}, line[1])
	assert.Equal(this.T(), []string{"value3", "value4"}, line[2])
	assert.Equal(this.T(), []string{}, line[3])

	line, err = target.GetLine(this.sheet2, 1, 2, 3)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 3)
	assert.Equal(this.T(), []string{"value5", "value6"}, line[1])
	assert.Equal(this.T(), []string{"value7", "value8"}, line[2])
	assert.Equal(this.T(), []string{}, line[3])

	_, err = target.GetLine(this.sheet1, -1)
	assert.NotNil(this.T(), err)

	_, err = target.GetLine(this.Unknown, 1)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestSheets() {
	target := &Excel{}
	assert.Equal(this.T(), []string{}, target.Sheets())
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	assert.Equal(this.T(), []string{this.sheet1, this.sheet2}, target.Sheets())
	CloseAll()
}

func (this *SuiteExcel) TestExist() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	assert.True(this.T(), target.Exist(this.sheet1))
	assert.True(this.T(), target.Exist(this.sheet2))
	assert.False(this.T(), target.Exist(this.Unknown))
	CloseAll()
}

func (this *SuiteExcel) TestSheet() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))

	sheet, err := target.Get(this.sheet1)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Next())
	data, err := sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value1", "value2"}, data)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value3", "value4"}, data)

	sheet, err = target.Get(this.sheet2)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value5", "value6"}, data)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value7", "value8"}, data)

	sheet, err = target.Get(this.sheet1)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(2))
	assert.False(this.T(), sheet.Nextn(-1))
	sheet.Close()

	sheet, err = target.Get(this.sheet1)
	assert.Nil(this.T(), err)
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
