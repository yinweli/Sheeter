package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeSheetEnum(t *testing.T) {
	suite.Run(t, new(SuiteInitializeSheetEnum))
}

type SuiteInitializeSheetEnum struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeSheetEnum) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeSheetEnum) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeSheetEnum) target() *initializeSheetEnum {
	target := &initializeSheetEnum{
		Global: &Global{
			LineOfEnum: 2,
		},
		Named: &nameds.Named{ExcelName: testdata.ExcelReal, SheetName: testdata.SheetEnum},
	}
	return target
}

func (this *SuiteInitializeSheetEnum) TestInitializeSheetEnum() {
	result := make(chan any, 1)
	target := this.target()
	assert.Nil(this.T(), InitializeSheetEnum(target, result))
	assert.NotNil(this.T(), target.layoutEnum)

	assert.Nil(this.T(), InitializeSheetEnum(nil, result))

	target = this.target()
	target.Named.ExcelName = "dep"
	target.Named.SheetName = "oT"
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.SheetName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.SheetName = "Enum2"
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidEnum
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidEnumDupl
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidIndex
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidIndexDupl
	assert.NotNil(this.T(), InitializeSheetEnum(target, result))
}
