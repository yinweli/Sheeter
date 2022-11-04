package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeExcel(t *testing.T) {
	suite.Run(t, new(SuiteInitializeExcel))
}

type SuiteInitializeExcel struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteInitializeExcel) SetupSuite() {
	this.Change("test-initializeExcel")
}

func (this *SuiteInitializeExcel) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteInitializeExcel) TestInitializeExcel() {
	result := make(chan any, 2)
	assert.Nil(this.T(), InitializeExcel(testdata.ExcelReal, result))
	assert.Len(this.T(), result, 2)
	sheet1 := (<-result).(*initializeSheetData)
	assert.Equal(this.T(), internal.SignData+"Data", sheet1.SheetName)
	assert.NotNil(this.T(), sheet1.excel)
	sheet2 := (<-result).(*initializeSheetEnum)
	assert.Equal(this.T(), internal.SignEnum+"Enum", sheet2.SheetName)
	assert.NotNil(this.T(), sheet2.excel)

	assert.Nil(this.T(), InitializeExcel(0, result))
}
