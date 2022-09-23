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
	fileJsonCsStruct := filepath.Join(internal.PathJson, internal.PathCs, structName+internal.ExtCs)
	fileJsonCsReader := filepath.Join(internal.PathJson, internal.PathCs, readerName+internal.ExtCs)
	fileJsonGoStruct := filepath.Join(internal.PathJson, internal.PathGo, structName+internal.ExtGo)
	fileJsonGoReader := filepath.Join(internal.PathJson, internal.PathGo, readerName+internal.ExtGo)
	fileJsonDataName := structName + internal.ExtJsonData
	fileJsonDataPath := filepath.Join(internal.PathJson, internal.PathData, fileJsonDataName)

	target := this.target()
	assert.Equal(this.T(), fileJsonCsStruct, target.FileJsonCsStruct())
	assert.Equal(this.T(), fileJsonCsReader, target.FileJsonCsReader())
	assert.Equal(this.T(), fileJsonGoStruct, target.FileJsonGoStruct())
	assert.Equal(this.T(), fileJsonGoReader, target.FileJsonGoReader())
	assert.Equal(this.T(), fileJsonDataName, target.FileJsonDataName())
	assert.Equal(this.T(), fileJsonDataPath, target.FileJsonDataPath())
}
