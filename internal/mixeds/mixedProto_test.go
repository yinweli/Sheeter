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
	pathProtoCs := filepath.Join(internal.PathProto, internal.PathCs)
	pathProtoGo := filepath.Join(internal.PathProto, internal.PathGo)
	pathProtoSchema := filepath.Join(internal.PathProto, internal.PathSchema)
	fileProtoName := structName + internal.ExtProtoSchema
	pathProtoName := filepath.Join(pathProtoSchema, fileProtoName)
	fileProtoData := structName + internal.ExtProtoData
	pathProtoData := filepath.Join(internal.PathProto, internal.PathData, fileProtoData)
	pathProtoCsReader := filepath.Join(internal.PathProto, internal.PathCs, utils.FirstUpper(readerName)+internal.ExtCs)
	pathProtoGoReader := filepath.Join(internal.PathProto, internal.PathGo, readerName+internal.ExtGo)
	protoDepend := utils.FirstLower(this.excel) + "." + internal.ExtProtoSchema

	target := this.target()
	assert.Equal(this.T(), pathProtoCs, target.PathProtoCs())
	assert.Equal(this.T(), pathProtoGo, target.PathProtoGo())
	assert.Equal(this.T(), pathProtoSchema, target.PathProtoSchema())
	assert.Equal(this.T(), fileProtoName, target.FileProtoName())
	assert.Equal(this.T(), pathProtoName, target.PathProtoName())
	assert.Equal(this.T(), fileProtoData, target.FileProtoData())
	assert.Equal(this.T(), pathProtoData, target.PathProtoData())
	assert.Equal(this.T(), pathProtoCsReader, target.PathProtoCsReader())
	assert.Equal(this.T(), pathProtoGoReader, target.PathProtoGoReader())
	assert.Equal(this.T(), protoDepend, target.ProtoDepend(this.excel))
}
