package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteInitialize) SetupSuite() {
	this.Change("test-initialize")
}

func (this *SuiteInitialize) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteInitialize) target() *Config {
	target := &Config{
		Global: Global{
			ExportJson:      true,
			ExportProto:     true,
			ExportEnum:      true,
			SimpleNamespace: false,
			LineOfName:      1,
			LineOfNote:      2,
			LineOfField:     3,
			LineOfLayer:     4,
			LineOfData:      5,
			LineOfEnum:      2,
		},
		Inputs: []string{
			testdata.ExcelReal + sheeter.SeparateSheet + testdata.SheetData,
			testdata.ExcelReal + sheeter.SeparateSheet + testdata.SheetEnum,
		},
	}
	return target
}

func (this *SuiteInitialize) TestInitialize() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.NotNil(this.T(), context)
	assert.NotNil(this.T(), context.Global)
	assert.NotEmpty(this.T(), context.Generate)
	assert.NotEmpty(this.T(), context.Encoding)
	assert.NotEmpty(this.T(), context.Poststep)
}

func (this *SuiteInitialize) TestPreparePath() {
	result := []any{"test1", "test2"}
	config := []string{"test1", "test2", "test2"}
	assert.Equal(this.T(), result, preparePath(config))
}

func (this *SuiteInitialize) TestPrepareExcel() {
	result := []any{"test1", "test2", "test3"}
	config := []string{"test1", "test2"}
	native := []any{"test2", "test3"}
	assert.Equal(this.T(), result, prepareExcel(config, native))
}

func (this *SuiteInitialize) TestPrepareSheet() {
	result := []any{
		&initializeSheetData{
			Named: &nameds.Named{ExcelName: "test1", SheetName: sheeter.SignData + "sheet"},
		},
		&initializeSheetEnum{
			Named: &nameds.Named{ExcelName: "test2", SheetName: sheeter.SignEnum + "sheet"},
		},
		&initializeSheetData{
			Named: &nameds.Named{ExcelName: "test3", SheetName: sheeter.SignData + "sheet"},
		},
		&initializeSheetEnum{
			Named: &nameds.Named{ExcelName: "test4", SheetName: sheeter.SignEnum + "sheet"},
		},
	}
	config := []Sheet{
		{ExcelName: "test1", SheetName: sheeter.SignData + "sheet"},
		{ExcelName: "test2", SheetName: sheeter.SignEnum + "sheet"},
		{ExcelName: "test1", SheetName: sheeter.SignData + "sheet"},
		{ExcelName: "testx", SheetName: "x"},
	}
	native := []any{
		&initializeSheetData{
			Named: &nameds.Named{ExcelName: "test3", SheetName: sheeter.SignData + "sheet"},
		},
		&initializeSheetData{
			Named: &nameds.Named{ExcelName: "test3", SheetName: sheeter.SignData + "sheet"},
		},
		&initializeSheetEnum{
			Named: &nameds.Named{ExcelName: "test4", SheetName: sheeter.SignEnum + "sheet"},
		},
		&initializeSheetEnum{
			Named: &nameds.Named{ExcelName: "test4", SheetName: sheeter.SignEnum + "sheet"},
		},
	}
	assert.Equal(this.T(), result, prepareSheet(config, native, nil))
}
