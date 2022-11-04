package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializePick(t *testing.T) {
	suite.Run(t, new(SuiteInitializePick))
}

type SuiteInitializePick struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteInitializePick) SetupSuite() {
	this.Change("test-initializePick")
}

func (this *SuiteInitializePick) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteInitializePick) target() *Config {
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
	}
	return target
}

func (this *SuiteInitializePick) TestInitializePick() {
	target := this.target()
	context := &Context{
		Global: &target.Global,
	}
	sheetData := &initializeSheetData{
		Global: &target.Global,
		Named:  &nameds.Named{ExcelName: testdata.ExcelReal, SheetName: testdata.SheetData},
	}
	sheetEnum := &initializeSheetEnum{
		Global: &target.Global,
		Named:  &nameds.Named{ExcelName: testdata.ExcelReal, SheetName: testdata.SheetEnum},
	}
	sheet := []any{sheetData, sheetEnum}

	for _, itor := range sheet {
		assert.Nil(this.T(), InitializeSheetData(itor, nil))
		assert.Nil(this.T(), InitializeSheetEnum(itor, nil))
	} // for

	assert.Nil(this.T(), InitializePick(sheet, context))
}
