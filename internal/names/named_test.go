package names

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestNamed(t *testing.T) {
	suite.Run(t, new(SuiteNamed))
}

type SuiteNamed struct {
	suite.Suite
	excel            string
	sheet            string
	workDir          string
	structName       string
	readerName       string
	fileJson         string
	fileJsonCode     string
	fileJsonSchema   string
	fileJsonCsCode   string
	fileJsonCsReader string
	fileJsonGoCode   string
	fileJsonGoReader string
	token            string
}

func (this *SuiteNamed) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excel"
	this.sheet = "sheet"
	this.structName = "ExcelSheet"
	this.readerName = "ExcelSheetReader"
	this.fileJson = filepath.Join(internal.PathJson, "excelSheet.json")
	this.fileJsonCode = filepath.ToSlash(this.fileJson)
	this.fileJsonSchema = filepath.Join(internal.PathJsonSchema, "excelSheet.json")
	this.fileJsonCsCode = filepath.Join(internal.PathJsonCs, "excelSheet.cs")
	this.fileJsonCsReader = filepath.Join(internal.PathJsonCs, "excelSheetReader.cs")
	this.fileJsonGoCode = filepath.Join(internal.PathJsonGo, "excelSheet.go")
	this.fileJsonGoReader = filepath.Join(internal.PathJsonGo, "excelSheetReader.go")
	this.token = "#"
}

func (this *SuiteNamed) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteNamed) target() *Named {
	target := &Named{
		Excel: this.excel,
		Sheet: this.sheet,
	}
	return target
}

func (this *SuiteNamed) TestName() {
	target := this.target()
	assert.Equal(this.T(), internal.AppName, target.AppName())
	assert.Equal(this.T(), internal.AppName, target.Namespace())
	assert.Equal(this.T(), this.structName, target.StructName())
	assert.Equal(this.T(), this.readerName, target.ReaderName())
	assert.Equal(this.T(), this.fileJson, target.FileJson())
	assert.Equal(this.T(), this.fileJsonCode, target.FileJsonCode())
	assert.Equal(this.T(), this.fileJsonSchema, target.FileJsonSchema())
	assert.Equal(this.T(), this.fileJsonCsCode, target.FileJsonCsCode())
	assert.Equal(this.T(), this.fileJsonCsReader, target.FileJsonCsReader())
	assert.Equal(this.T(), this.fileJsonGoCode, target.FileJsonGoCode())
	assert.Equal(this.T(), this.fileJsonGoReader, target.FileJsonGoReader())
	assert.Equal(this.T(), this.fileJsonCode, target.FileJsonCode())
}

func (this *SuiteNamed) TestCombine() {
	target := this.target()
	assert.Equal(this.T(), "excelsheet", target.combine(params{}))
	assert.Equal(this.T(), "Excelsheet", target.combine(params{excelUpper: true}))
	assert.Equal(this.T(), "excelSheet", target.combine(params{sheetUpper: true}))
	assert.Equal(this.T(), "excelsheet#", target.combine(params{last: this.token}))
	assert.Equal(this.T(), filepath.Join(this.token, "excelsheet"), target.combine(params{path: this.token}))
	assert.Equal(this.T(), "excelsheet.#", target.combine(params{ext: this.token}))
}
