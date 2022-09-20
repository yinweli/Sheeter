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

func TestProto(t *testing.T) {
	suite.Run(t, new(SuiteProto))
}

type SuiteProto struct {
	suite.Suite
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelProto"
	this.sheet = "sheetProto"
}

func (this *SuiteProto) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteProto) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}

func (this *SuiteProto) TestName() {
	structName := this.excel + utils.FirstUpper(this.sheet) + "."
	readerName := this.excel + utils.FirstUpper(this.sheet) + internal.Reader + "."
	pathProtoSchema := internal.PathProtoSchema
	pathProtoCs := internal.PathProtoCs
	pathProtoGo := internal.PathProtoGo
	fileProtoDepend := this.excel + "." + internal.ExtProtoSchema
	fileProtoSchema := filepath.Join(internal.PathProtoSchema, structName+internal.ExtProtoSchema)
	fileProtoCsReader := filepath.Join(internal.PathProtoCs, readerName+internal.ExtCs)
	fileProtoGoReader := filepath.Join(internal.PathProtoGo, readerName+internal.ExtGo)
	fileProtoData := filepath.Join(internal.PathProtoData, structName+internal.ExtProtoData)
	fileProtoDataCode := filepath.ToSlash(fileProtoData)
	fileProtoBat := internal.PathProtoSchema + "." + internal.ExtBat
	fileProtoSh := internal.PathProtoSchema + "." + internal.ExtSh

	target := this.target()
	assert.Equal(this.T(), pathProtoSchema, target.PathProtoSchema())
	assert.Equal(this.T(), pathProtoCs, target.PathProtoCs())
	assert.Equal(this.T(), pathProtoGo, target.PathProtoGo())
	assert.Equal(this.T(), fileProtoDepend, target.FileProtoDepend(this.excel))
	assert.Equal(this.T(), fileProtoSchema, target.FileProtoSchema())
	assert.Equal(this.T(), fileProtoCsReader, target.FileProtoCsReader())
	assert.Equal(this.T(), fileProtoGoReader, target.FileProtoGoReader())
	assert.Equal(this.T(), fileProtoData, target.FileProtoData())
	assert.Equal(this.T(), fileProtoDataCode, target.FileProtoDataCode())
	assert.Equal(this.T(), fileProtoBat, target.FileProtoBat())
	assert.Equal(this.T(), fileProtoSh, target.FileProtoSh())
}
