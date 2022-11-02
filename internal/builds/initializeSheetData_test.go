package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeSheetData(t *testing.T) {
	suite.Run(t, new(SuiteInitializeSheetData))
}

type SuiteInitializeSheetData struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeSheetData) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeSheetData) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeSheetData) target() *initializeSheetData {
	target := &initializeSheetData{
		Global: &Global{
			LineOfName:  1,
			LineOfNote:  2,
			LineOfField: 3,
			LineOfLayer: 4,
		},
		Named: &nameds.Named{ExcelName: testdata.ExcelReal, SheetName: testdata.SheetData},
	}
	return target
}

func (this *SuiteInitializeSheetData) TestInitializeSheetData() {
	result := make(chan any, 1)
	target := this.target()
	assert.Nil(this.T(), InitializeSheetData(target, result))
	assert.Empty(this.T(), result)
	assert.NotNil(this.T(), target.excel)
	assert.NotNil(this.T(), target.layoutData)
	assert.NotNil(this.T(), target.layoutType)
	assert.NotNil(this.T(), target.layoutDepend)

	assert.Nil(this.T(), InitializeSheetData(nil, result))

	target = this.target()
	target.Global.LineOfName = -1
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Global.LineOfNote = -1
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Global.LineOfField = -1
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Global.LineOfLayer = -1
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = "Dep"
	target.Named.SheetName = "ot"
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.SheetName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.SheetName = "Data2"
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidFile
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelCleanAll
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelCleanField
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidField
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidLayer
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidLayout
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidPkeyZero
	assert.NotNil(this.T(), InitializeSheetData(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidPkeyDupl
	assert.NotNil(this.T(), InitializeSheetData(target, result))
}