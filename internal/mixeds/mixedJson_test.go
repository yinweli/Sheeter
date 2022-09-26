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
	structName := this.excel + utils.FirstUpper(this.sheet) + "."
	readerName := this.excel + utils.FirstUpper(this.sheet) + internal.Reader + "."
	fileJsonData := structName + internal.ExtJsonData
	pathJsonData := filepath.Join(internal.PathJson, internal.PathData, structName+internal.ExtJsonData)
	pathJsonCsStruct := filepath.Join(internal.PathJson, internal.PathCs, structName+internal.ExtCs)
	pathJsonCsReader := filepath.Join(internal.PathJson, internal.PathCs, readerName+internal.ExtCs)
	pathJsonGoStruct := filepath.Join(internal.PathJson, internal.PathGo, structName+internal.ExtGo)
	pathJsonGoReader := filepath.Join(internal.PathJson, internal.PathGo, readerName+internal.ExtGo)

	target := this.target()
	assert.Equal(this.T(), fileJsonData, target.FileJsonData())
	assert.Equal(this.T(), pathJsonData, target.PathJsonData())
	assert.Equal(this.T(), pathJsonCsStruct, target.PathJsonCsStruct())
	assert.Equal(this.T(), pathJsonCsReader, target.PathJsonCsReader())
	assert.Equal(this.T(), pathJsonGoStruct, target.PathJsonGoStruct())
	assert.Equal(this.T(), pathJsonGoReader, target.PathJsonGoReader())
}
