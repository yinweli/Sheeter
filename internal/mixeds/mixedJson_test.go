package mixeds

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
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelJson"
	this.sheet = "sheetJson"
}

func (this *SuiteJson) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJson) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}

func (this *SuiteJson) TestName() {
	name := this.excel + utils.FirstUpper(this.sheet)
	structName := name + "."
	readerName := name + internal.Reader + "."
	jsonDataFile := structName + internal.JsonDataExt
	jsonDataPath := filepath.Join(internal.JsonPath, internal.DataPath, structName+internal.JsonDataExt)
	jsonCsStructPath := filepath.Join(internal.JsonPath, internal.CsPath, structName+internal.CsExt)
	jsonCsReaderPath := filepath.Join(internal.JsonPath, internal.CsPath, readerName+internal.CsExt)
	jsonCsDepotPath := filepath.Join(internal.JsonPath, internal.CsPath, internal.Depot+"."+internal.CsExt)
	jsonGoStructPath := filepath.Join(internal.JsonPath, internal.GoPath, structName+internal.GoExt)
	jsonGoReaderPath := filepath.Join(internal.JsonPath, internal.GoPath, readerName+internal.GoExt)
	jsonGoDepotPath := filepath.Join(internal.JsonPath, internal.GoPath, internal.Depot+"."+internal.GoExt)

	target := this.target()
	assert.Equal(this.T(), name, target.JsonDataName())
	assert.Equal(this.T(), internal.JsonDataExt, target.JsonDataExt())
	assert.Equal(this.T(), jsonDataFile, target.JsonDataFile())
	assert.Equal(this.T(), jsonDataPath, target.JsonDataPath())
	assert.Equal(this.T(), jsonCsStructPath, target.JsonCsStructPath())
	assert.Equal(this.T(), jsonCsReaderPath, target.JsonCsReaderPath())
	assert.Equal(this.T(), jsonCsDepotPath, target.JsonCsDepotPath())
	assert.Equal(this.T(), jsonGoStructPath, target.JsonGoStructPath())
	assert.Equal(this.T(), jsonGoReaderPath, target.JsonGoReaderPath())
	assert.Equal(this.T(), jsonGoDepotPath, target.JsonGoDepotPath())
}
