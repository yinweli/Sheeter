package nameds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	testdata.TestEnv
	excelName string
	sheetName string
}

func (this *SuiteJson) SetupSuite() {
	this.Change("test-named-json")
	this.excelName = "excelJson"
	this.sheetName = "sheetJson"
}

func (this *SuiteJson) TearDownSuite() {
	this.Restore()
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
	structName := name
	readerName := name + sheeter.Reader
	jsonDataFile := structName + sheeter.JsonDataExt
	jsonDataPath := filepath.Join(sheeter.JsonPath, sheeter.DataPath, structName+sheeter.JsonDataExt)
	jsonStructCsPath := filepath.Join(sheeter.JsonPath, sheeter.CsPath, utils.FirstUpper(structName)+sheeter.CsExt)
	jsonReaderCsPath := filepath.Join(sheeter.JsonPath, sheeter.CsPath, utils.FirstUpper(readerName)+sheeter.CsExt)
	jsonDepotCsPath := filepath.Join(sheeter.JsonPath, sheeter.CsPath, utils.FirstUpper(sheeter.Depot)+sheeter.CsExt)
	jsonStructGoPath := filepath.Join(sheeter.JsonPath, sheeter.GoPath, structName+sheeter.GoExt)
	jsonReaderGoPath := filepath.Join(sheeter.JsonPath, sheeter.GoPath, readerName+sheeter.GoExt)
	jsonDepotGoPath := filepath.Join(sheeter.JsonPath, sheeter.GoPath, sheeter.Depot+sheeter.GoExt)

	target := this.target()
	assert.Equal(this.T(), name, target.JsonDataName())
	assert.Equal(this.T(), sheeter.JsonDataExt, target.JsonDataExt())
	assert.Equal(this.T(), jsonDataFile, target.JsonDataFile())
	assert.Equal(this.T(), jsonDataPath, target.JsonDataPath())
	assert.Equal(this.T(), jsonStructCsPath, target.JsonStructCsPath())
	assert.Equal(this.T(), jsonReaderCsPath, target.JsonReaderCsPath())
	assert.Equal(this.T(), jsonDepotCsPath, target.JsonDepotCsPath())
	assert.Equal(this.T(), jsonStructGoPath, target.JsonStructGoPath())
	assert.Equal(this.T(), jsonReaderGoPath, target.JsonReaderGoPath())
	assert.Equal(this.T(), jsonDepotGoPath, target.JsonDepotGoPath())
}
