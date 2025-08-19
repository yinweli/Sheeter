package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/testdata"
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
	this.Env = testdata.EnvSetup("test-builds-initialize", "initialize")
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
	testdata.EnvRestore(this.Env)
}

func (this *SuiteInitialize) TestInitialize() {
	result, err := Initialize(&Config{
		Source:  []string{this.folderSuccess},
		Exclude: []string{"exclude#exclude"},
	})
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), result, 7)

	for _, itor := range result {
		assert.NotNil(this.T(), itor.Excel)
		assert.NotNil(this.T(), itor.Sheet)
		assert.NotEmpty(this.T(), itor.ExcelName)
		assert.NotEmpty(this.T(), itor.SheetName)
	} // for

	_, err = Initialize(&Config{
		Source: []string{this.folderFailed},
	})
	assert.Len(this.T(), err, 3)
}

func (this *SuiteInitialize) TestSearchExcel() {
	result := searchExcel(this.folderSearchExcel)
	assert.Nil(this.T(), result.Error)
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess1), result.Result[0])
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess2), result.Result[1])
	assert.Equal(this.T(), filepath.Join(this.folderSearchExcel, this.excelSuccess3), result.Result[2])
}

func (this *SuiteInitialize) TestSearchSheet() {
	result := searchSheet(filepath.Join(this.folderSearchSheet, this.excelSuccess1))
	assert.Nil(this.T(), result.Error)
	prepare := result.Result[0].(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), this.excelSuccess1, prepare.ExcelName)
	assert.Equal(this.T(), this.sheet1, prepare.SheetName)
	prepare = result.Result[1].(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), this.excelSuccess1, prepare.ExcelName)
	assert.Equal(this.T(), this.sheet2, prepare.SheetName)

	result = searchSheet(filepath.Join(this.folderSearchSheet, this.excelFailed))
	assert.NotNil(this.T(), result.Error)
}
