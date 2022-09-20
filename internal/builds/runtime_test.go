package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestRuntime(t *testing.T) {
	suite.Run(t, new(SuiteRuntime))
}

type SuiteRuntime struct {
	suite.Suite
	workDir string
}

func (this *SuiteRuntime) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteRuntime) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteRuntime) runtimeSector() *RuntimeSector {
	sector := &RuntimeSector{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return sector
}

func (this *SuiteRuntime) TestGetRows() {
	sector := this.runtimeSector()
	sector.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	rows, err := sector.GetRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = sector.GetRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = sector.GetRows(0)
	assert.NotNil(this.T(), err)

	sector.Sheet = testdata.UnknownStr
	_, err = sector.GetRows(1)
	assert.NotNil(this.T(), err)

	sector.Close()
}

func (this *SuiteRuntime) TestGetColumns() {
	sector := this.runtimeSector()
	sector.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := sector.GetColumns(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{
			"name0#pkey",
			"empty#empty",
			"name1#bool",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
		}, cols)

	_, err = sector.GetColumns(10)
	assert.NotNil(this.T(), err)

	_, err = sector.GetColumns(0)
	assert.NotNil(this.T(), err)

	sector.Sheet = testdata.UnknownStr
	_, err = sector.GetColumns(1)
	assert.NotNil(this.T(), err)

	sector.Close()
}
