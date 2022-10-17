package nameds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	workDir   string
	excelName string
	sheetName string
}

func (this *SuiteJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excelName = "excelJson"
	this.sheetName = "sheetJson"
}

func (this *SuiteJson) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJson) target() *Json {
	target := &Json{
		ExcelName: this.excelName,
		SheetName: this.sheetName,
	}
	return target
}

func (this *SuiteJson) TestName() {
	name := this.excelName + utils.FirstUpper(this.sheetName)
	structName := name + "."
	readerName := name + internal.Reader + "."
	jsonDataFile := structName + internal.JsonDataExt
	jsonDataPath := filepath.Join(internal.JsonPath, internal.DataPath, structName+internal.JsonDataExt)
	jsonStructCsPath := filepath.Join(internal.JsonPath, internal.CsPath, structName+internal.CsExt)
	jsonReaderCsPath := filepath.Join(internal.JsonPath, internal.CsPath, readerName+internal.CsExt)
	jsonDepotCsPath := filepath.Join(internal.JsonPath, internal.CsPath, internal.Depot+"."+internal.CsExt)
	jsonStructGoPath := filepath.Join(internal.JsonPath, internal.GoPath, structName+internal.GoExt)
	jsonReaderGoPath := filepath.Join(internal.JsonPath, internal.GoPath, readerName+internal.GoExt)
	jsonDepotGoPath := filepath.Join(internal.JsonPath, internal.GoPath, internal.Depot+"."+internal.GoExt)

	target := this.target()
	assert.Equal(this.T(), name, target.JsonDataName())
	assert.Equal(this.T(), internal.JsonDataExt, target.JsonDataExt())
	assert.Equal(this.T(), jsonDataFile, target.JsonDataFile())
	assert.Equal(this.T(), jsonDataPath, target.JsonDataPath())
	assert.Equal(this.T(), jsonStructCsPath, target.JsonStructCsPath())
	assert.Equal(this.T(), jsonReaderCsPath, target.JsonReaderCsPath())
	assert.Equal(this.T(), jsonDepotCsPath, target.JsonDepotCsPath())
	assert.Equal(this.T(), jsonStructGoPath, target.JsonStructGoPath())
	assert.Equal(this.T(), jsonReaderGoPath, target.JsonReaderGoPath())
	assert.Equal(this.T(), jsonDepotGoPath, target.JsonDepotGoPath())
}
