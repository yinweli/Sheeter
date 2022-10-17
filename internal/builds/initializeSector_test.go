package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeSector(t *testing.T) {
	suite.Run(t, new(SuiteInitializeSector))
}

type SuiteInitializeSector struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeSector) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeSector) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeSector) target() *Context {
	target := &Context{
		Config: &Config{
			Global: Global{
				LineOfName:  1,
				LineOfNote:  2,
				LineOfField: 3,
				LineOfLayer: 4,
			},
		},
		Sector: []*ContextSector{
			{
				Element: Element{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetData,
				},
			},
		},
	}
	return target
}

func (this *SuiteInitializeSector) TestInitializeSector() {
	target := this.target()
	sector := target.Sector[0]
	assert.Nil(this.T(), initializeSector(target, sector))
	assert.NotNil(this.T(), sector.excel)
	assert.NotNil(this.T(), sector.layoutJson)
	assert.NotNil(this.T(), sector.layoutType)
	assert.NotNil(this.T(), sector.layoutDepend)
	sector.Close()

	target = this.target()
	target.Global.LineOfName = -1
	sector = target.Sector[0]
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	target.Global.LineOfNote = -1
	sector = target.Sector[0]
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	target.Global.LineOfField = -1
	sector = target.Sector[0]
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	target.Global.LineOfLayer = -1
	sector = target.Sector[0]
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = "Dep"
	sector.Sheet = "ot"
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameCleanAll
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameCleanField
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidField
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidLayer
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidLayout
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidPkeyZero
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()

	target = this.target()
	sector = target.Sector[0]
	sector.Excel = testdata.ExcelNameInvalidPkeyDupl
	assert.NotNil(this.T(), initializeSector(target, sector))
	sector.Close()
}
