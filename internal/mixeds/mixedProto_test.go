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
	pathProtoSchema := internal.PathProto
	pathProtoCs := filepath.Join(internal.PathProto, internal.PathCs)
	pathProtoGo := filepath.Join(internal.PathProto, internal.PathGo)
	fileProtoSchema := filepath.Join(internal.PathProto, structName+internal.ExtProtoSchema)
	fileProtoCsReader := filepath.Join(internal.PathProto, internal.PathCs, readerName+internal.ExtCs)
	fileProtoGoReader := filepath.Join(internal.PathProto, internal.PathGo, readerName+internal.ExtGo)
	fileProtoDataName := structName + internal.ExtProtoData
	fileProtoDataPath := filepath.Join(internal.PathProto, internal.PathData, fileProtoDataName)
	fileProtoBat := filepath.Join(internal.PathProto, internal.FileProtoBat)
	fileProtoSh := filepath.Join(internal.PathProto, internal.FileProtoSh)
	protoDepend := this.excel + "." + internal.ExtProtoSchema

	target := this.target()
	assert.Equal(this.T(), pathProtoSchema, target.PathProtoSchema())
	assert.Equal(this.T(), pathProtoCs, target.PathProtoCs())
	assert.Equal(this.T(), pathProtoGo, target.PathProtoGo())
	assert.Equal(this.T(), fileProtoSchema, target.FileProtoSchema())
	assert.Equal(this.T(), fileProtoCsReader, target.FileProtoCsReader())
	assert.Equal(this.T(), fileProtoGoReader, target.FileProtoGoReader())
	assert.Equal(this.T(), fileProtoDataName, target.FileProtoDataName())
	assert.Equal(this.T(), fileProtoDataPath, target.FileProtoDataPath())
	assert.Equal(this.T(), fileProtoBat, target.FileProtoBat())
	assert.Equal(this.T(), fileProtoSh, target.FileProtoSh())
	assert.Equal(this.T(), protoDepend, target.ProtoDepend(this.excel))
}
