package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestContent(t *testing.T) {
	suite.Run(t, new(SuiteContent))
}

type SuiteContent struct {
	suite.Suite
	workDir string
}

func (this *SuiteContent) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteContent) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteContent) target() *Content {
	target := &Content{
		Bom:         true,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteContent) TestStructName() {
	assert.Equal(this.T(), "RealData", this.target().StructName())
}

func (this *SuiteContent) TestFileJson() {
	assert.Equal(this.T(), filepath.Join(pathJson, "realData.json"), this.target().FileJson())
}

func (this *SuiteContent) TestFileJsonSchema() {
	assert.Equal(this.T(), filepath.Join(pathJsonSchema, "realData.json"), this.target().FileJsonSchema())
}

func (this *SuiteContent) TestCombine() {
	target := this.target()

	assert.Equal(this.T(), "Realdata", target.combine(params{excelUpper: true}))
	assert.Equal(this.T(), "realdata", target.combine(params{excelUpper: false}))
	assert.Equal(this.T(), "realData", target.combine(params{sheetUpper: true}))
	assert.Equal(this.T(), "realdata", target.combine(params{sheetUpper: false}))
	assert.Equal(this.T(), "realdata.txt", target.combine(params{ext: "txt"}))
}

func (this *SuiteContent) TestGetRows() {
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

func (this *SuiteContent) TestGetColumns() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := target.GetColumns(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{"name0#pkey", "empty#empty", "name1#bool", "name2#int", "name3#text", "name2#int", "name3#text",
			"name2#int", "name3#text"}, cols)

	_, err = target.GetColumns(10)
	assert.NotNil(this.T(), err)

	_, err = target.GetColumns(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.GetColumns(1)
	assert.NotNil(this.T(), err)

	target.Close()
}

func TestContents(t *testing.T) {
	suite.Run(t, new(SuiteContents))
}

type SuiteContents struct {
	suite.Suite
	workDir   string
	lineEmpty string
	lineNew   string
}

func (this *SuiteContents) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.lineEmpty = ""
	this.lineNew = "\n"
}

func (this *SuiteContents) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteContents) target() *Contents {
	return &Contents{
		Contents: []*Content{{}, {}},
	}
}

func (this *SuiteContents) TestName() {
	assert.Equal(this.T(), internal.Title, this.target().AppName())
	assert.Equal(this.T(), internal.Title, this.target().Namespace())
	assert.Equal(this.T(), pathJson, this.target().PathJson())
	assert.Equal(this.T(), pathJsonSchema, this.target().PathJsonSchema())
	assert.Equal(this.T(), pathJsonCs, this.target().PathJsonCs())
	assert.Equal(this.T(), pathJsonGo, this.target().PathJsonGo())
	assert.Equal(this.T(), filepath.Join(pathJsonCs, fileJsonCsCode), this.target().FileJsonCsCode())
	assert.Equal(this.T(), filepath.Join(pathJsonCs, fileJsonCsReader), this.target().FileJsonCsReader())
	assert.Equal(this.T(), filepath.Join(pathJsonGo, fileJsonGoCode), this.target().FileJsonGoCode())
	assert.Equal(this.T(), filepath.Join(pathJsonGo, fileJsonGoReader), this.target().FileJsonGoReader())
}

func (this *SuiteContents) TestLine() {
	target := this.target()

	assert.Equal(this.T(), this.lineEmpty, target.SetLine())
	assert.Equal(this.T(), this.lineNew, target.NewLine())
	assert.Equal(this.T(), this.lineEmpty, target.NewLine())
	assert.Equal(this.T(), this.lineEmpty, target.NewLine())
}
