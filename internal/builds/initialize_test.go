package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitialize) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitialize) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitialize) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteInitialize) TestInitialize() {
	target := this.target()
	assert.Nil(this.T(), Initialize(target))
	assert.NotNil(this.T(), target.excel)
	target.Close()

	target = this.target()
	target.LineOfField = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.LineOfLayer = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.LineOfNote = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanAll
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanField
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidField
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayer
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayout
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyZero
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyDupl
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), Initialize(target))
	target.Close()
}
