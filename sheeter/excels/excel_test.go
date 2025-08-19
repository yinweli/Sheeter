package excels

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thedatashed/xlsxreader"

	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestExcel(t *testing.T) {
	suite.Run(t, new(SuiteExcel))
}

type SuiteExcel struct {
	suite.Suite
	testdata.Env
	excelSuccess string
	sheet1       string
	sheet2       string
}

func (this *SuiteExcel) SetupSuite() {
	this.Env = testdata.EnvSetup("test-excels-excel", "excel")
	this.excelSuccess = "success.xlsx"
	this.sheet1 = "Test1"
	this.sheet2 = "Test2"
}

func (this *SuiteExcel) TearDownSuite() {
	CloseAll()
	testdata.EnvRestore(this.Env)
}

func (this *SuiteExcel) TestOpen() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	target.Close()

	target = &Excel{}
	assert.NotNil(this.T(), target.Open(testdata.Unknown))
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

	_, err = target.Get(testdata.Unknown)
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

	_, err = target.GetLine(testdata.Unknown, 1)
	assert.NotNil(this.T(), err)
	CloseAll()
}

func (this *SuiteExcel) TestSheets() {
	target := &Excel{}
	assert.Empty(this.T(), target.Sheet())
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	assert.Equal(this.T(), []string{this.sheet1, this.sheet2}, target.Sheet())
	CloseAll()
}

func (this *SuiteExcel) TestExist() {
	target := &Excel{}
	assert.Nil(this.T(), target.Open(this.excelSuccess))
	assert.True(this.T(), target.Exist(this.sheet1))
	assert.True(this.T(), target.Exist(this.sheet2))
	assert.False(this.T(), target.Exist(testdata.Unknown))
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
	assert.Equal(this.T(), 2, sheet.Line())

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

func (this *SuiteExcel) TestFormatNumerical() {
	assert.Equal(this.T(), testdata.Unknown, formatNumerical(xlsxreader.Cell{
		Value: testdata.Unknown,
		Type:  xlsxreader.TypeString,
	}))
	assert.Equal(this.T(), "0."+strings.Repeat("0", 19)+"1", formatNumerical(xlsxreader.Cell{
		Value: "1e-20",
		Type:  xlsxreader.TypeNumerical,
	}))
	assert.Equal(this.T(), "1"+strings.Repeat("0", 20), formatNumerical(xlsxreader.Cell{
		Value: "1e+20",
		Type:  xlsxreader.TypeNumerical,
	}))
	assert.Equal(this.T(), testdata.Unknown, formatNumerical(xlsxreader.Cell{
		Value: testdata.Unknown,
		Type:  xlsxreader.TypeNumerical,
	}))
	assert.Equal(this.T(), "1e?20", formatNumerical(xlsxreader.Cell{
		Value: "1e?20",
		Type:  xlsxreader.TypeNumerical,
	}))
	assert.Equal(this.T(), "1e+400", formatNumerical(xlsxreader.Cell{
		Value: "1e+400",
		Type:  xlsxreader.TypeNumerical,
	}))
}
