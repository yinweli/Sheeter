package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestSector(t *testing.T) {
	suite.Run(t, new(SuiteSector))
}

type SuiteSector struct {
	suite.Suite
	workDir string
	token   string
}

func (this *SuiteSector) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.token = "#"
}

func (this *SuiteSector) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSector) target() *Sector {
	target := &Sector{
		Global: Global{
			Bom:         true,
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

func (this *SuiteSector) TestName() {
	target := this.target()
	assert.Equal(this.T(), internal.AppName, target.AppName())
	assert.Equal(this.T(), "realdata", target.Namespace())
	assert.Equal(this.T(), internal.Struct, target.StructName())
	assert.Equal(this.T(), internal.Reader, target.ReaderName())
	assert.Equal(this.T(), filepath.Join(internal.PathJson, "realData.json"), target.FileJson())
	assert.Equal(this.T(), filepath.Join(internal.PathJsonSchema, "realData.json"), target.FileJsonSchema())
	assert.Equal(this.T(), filepath.Join(internal.PathJsonCs, "realData.cs"), target.FileJsonCsCode())
	assert.Equal(this.T(), filepath.Join(internal.PathJsonCs, "realDataReader.cs"), target.FileJsonCsReader())
	assert.Equal(this.T(), filepath.Join(internal.PathJsonGo, "realData.go"), target.FileJsonGoCode())
	assert.Equal(this.T(), filepath.Join(internal.PathJsonGo, "realDataReader.go"), target.FileJsonGoReader())
}

func (this *SuiteSector) TestCombine() {
	target := this.target()
	assert.Equal(this.T(), "realdata", target.combine(params{}))
	assert.Equal(this.T(), "Realdata", target.combine(params{excelUpper: true}))
	assert.Equal(this.T(), "realData", target.combine(params{sheetUpper: true}))
	assert.Equal(this.T(), "real#data", target.combine(params{middle: this.token}))
	assert.Equal(this.T(), "realdata#", target.combine(params{last: this.token}))
	assert.Equal(this.T(), "realdata.#", target.combine(params{ext: this.token}))
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
