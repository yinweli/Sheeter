package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/testdata"
)

func TestSector(t *testing.T) {
	suite.Run(t, new(SuiteSector))
}

type SuiteSector struct {
	suite.Suite
	workDir string
}

func (this *SuiteSector) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteSector) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSector) target() *Sector {
	target := &Sector{
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
	return target
}

func (this *SuiteSector) TestGetRows() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	rows, err := target.GetRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = target.GetRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = target.GetRows(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.GetRows(1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteSector) TestGetColumns() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := target.GetColumns(1)
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

	_, err = target.GetColumns(10)
	assert.NotNil(this.T(), err)

	_, err = target.GetColumns(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.GetColumns(1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func (this *SuiteSector) TestMergeSectorLayoutType() {
	sectors := []*Sector{
		{layoutType: layouts.NewLayoutType()},
		{layoutType: layouts.NewLayoutType()},
		{layoutType: layouts.NewLayoutType()},
	}
	layoutType, err := MergeSectorLayoutType(sectors)
	assert.NotNil(this.T(), layoutType)
	assert.Nil(this.T(), err)
}
