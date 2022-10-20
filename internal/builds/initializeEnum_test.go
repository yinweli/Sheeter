package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeEnum(t *testing.T) {
	suite.Run(t, new(SuiteInitializeEnum))
}

type SuiteInitializeEnum struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeEnum) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeEnum) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeEnum) target() *initializeEnum {
	target := &initializeEnum{
		Global: &Global{
			LineOfEnum: 2,
		},
		Named: &nameds.Named{ExcelName: testdata.ExcelReal, SheetName: testdata.SheetEnum},
	}
	return target
}

func (this *SuiteInitializeEnum) TestInitializeEnum() {
	target := this.target()
	assert.Nil(this.T(), InitializeEnum(target))
	assert.NotNil(this.T(), target.layoutEnum)

	assert.Nil(this.T(), InitializeEnum(nil))

	target = this.target()
	target.Named.ExcelName = "dep"
	target.Named.SheetName = "oT"
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.ExcelName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.SheetName = testdata.UnknownStr
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.SheetName = "Enum2"
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidEnum
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidEnumDupl
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidIndex
	assert.NotNil(this.T(), InitializeEnum(target))

	target = this.target()
	target.Named.ExcelName = testdata.ExcelInvalidIndexDupl
	assert.NotNil(this.T(), InitializeEnum(target))
}
