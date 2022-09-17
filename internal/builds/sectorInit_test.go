package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestSectorInit(t *testing.T) {
	suite.Run(t, new(SuiteSectorInit))
}

type SuiteSectorInit struct {
	suite.Suite
	workDir string
}

func (this *SuiteSectorInit) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteSectorInit) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSectorInit) target() *Sector {
	target := &Sector{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return target
}

func (this *SuiteSectorInit) TestSectorInit() {
	target := this.target()
	assert.Nil(this.T(), SectorInit(target))
	assert.NotNil(this.T(), target.Excel)
	target.Close()

	target = this.target()
	target.LineOfField = 10
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.LineOfLayer = 10
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.LineOfNote = 10
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanAll
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanField
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidField
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayer
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayout
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyZero
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyDupl
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), SectorInit(target))
	target.Close()
}
