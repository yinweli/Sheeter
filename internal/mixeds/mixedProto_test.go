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
	name := this.excel + utils.FirstUpper(this.sheet)
	structName := name + "."
	readerName := name + internal.Reader + "."
	protoCsPath := filepath.Join(internal.PathProto, internal.PathCs)
	protoGoPath := filepath.Join(internal.PathProto, internal.PathGo)
	protoSchemaPath := filepath.Join(internal.PathProto, internal.PathSchema)
	protoName := structName + internal.ExtProtoSchema
	protoPath := filepath.Join(protoSchemaPath, protoName)
	protoDataFile := structName + internal.ExtProtoData
	protoDataPath := filepath.Join(internal.PathProto, internal.PathData, protoDataFile)
	protoCsReaderPath := filepath.Join(internal.PathProto, internal.PathCs, utils.FirstUpper(readerName)+internal.ExtCs)
	protoGoReaderPath := filepath.Join(internal.PathProto, internal.PathGo, readerName+internal.ExtGo)
	protoDepend := utils.FirstLower(this.excel) + "." + internal.ExtProtoSchema

	target := this.target()
	assert.Equal(this.T(), protoCsPath, target.ProtoCsPath())
	assert.Equal(this.T(), protoGoPath, target.ProtoGoPath())
	assert.Equal(this.T(), protoSchemaPath, target.ProtoSchemaPath())
	assert.Equal(this.T(), protoName, target.ProtoName())
	assert.Equal(this.T(), protoPath, target.ProtoPath())
	assert.Equal(this.T(), name, target.ProtoDataName())
	assert.Equal(this.T(), internal.ExtProtoData, target.ProtoDataExt())
	assert.Equal(this.T(), protoDataFile, target.ProtoDataFile())
	assert.Equal(this.T(), protoDataPath, target.ProtoDataPath())
	assert.Equal(this.T(), protoCsReaderPath, target.ProtoCsReaderPath())
	assert.Equal(this.T(), protoGoReaderPath, target.ProtoGoReaderPath())
	assert.Equal(this.T(), protoDepend, target.ProtoDepend(this.excel))
}
