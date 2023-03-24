package builds

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	testdata.Env
	folderSuccess     string
	folderFailed      string
	folderSearchExcel string
	folderSearchSheet string
	excelSuccess1     string
	excelSuccess2     string
	excelSuccess3     string
	excelFailed       string
	sheet1            string
	sheet2            string
}

func (this *SuiteInitialize) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-builds-initialize", "initialize")
	this.folderSuccess = "success"
	this.folderFailed = "failed"
	this.folderSearchExcel = "searchExcel"
	this.folderSearchSheet = "searchSheet"
	this.excelSuccess1 = "success1.xlsx"
	this.excelSuccess2 = "success2.xlsx"
	this.excelSuccess3 = "success3.xlsx"
	this.excelFailed = "failed.xlsx"
	this.sheet1 = "Test1"
	this.sheet2 = "Test2"
}

func (this *SuiteInitialize) TearDownSuite() {
	excels.CloseAll()
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteInitialize) TestInitialize() {
	config := &Config{
		Source:  []string{this.folderSuccess},
		Exclude: []string{"exclude#exclude"},
	}
	context, err := Initialize(config)
	time.Sleep(testdata.Timeout)
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), context, 7)

	for _, itor := range context {
		assert.NotNil(this.T(), itor.Excel)
		assert.NotNil(this.T(), itor.Sheet)
		assert.NotEmpty(this.T(), itor.ExcelName)
		assert.NotEmpty(this.T(), itor.SheetName)
	} // for

	_, err = Initialize(&Config{
		Source: []string{this.folderFailed},
	})
	time.Sleep(testdata.Timeout)
	assert.Len(this.T(), err, 3)
}

func (this *SuiteInitialize) TestSearchExcel() {
	result := make(chan any, sheeter.MaxExcel)
	assert.Nil(this.T(), searchExcel(this.folderSearchExcel, result))
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess1), <-result)
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess2), <-result)
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess3), <-result)
}

func (this *SuiteInitialize) TestSearchSheet() {
	result := make(chan any, sheeter.MaxExcel)
	assert.Nil(this.T(), searchSheet(filepath.Join(this.folderSearchSheet, this.excelSuccess1), result))
	prepare := (<-result).(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), this.excelSuccess1, prepare.ExcelName)
	assert.Equal(this.T(), this.sheet1, prepare.SheetName)
	prepare = (<-result).(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), this.excelSuccess1, prepare.ExcelName)
	assert.Equal(this.T(), this.sheet2, prepare.SheetName)

	assert.NotNil(this.T(), searchSheet(filepath.Join(this.folderSearchSheet, this.excelFailed), result))
}
