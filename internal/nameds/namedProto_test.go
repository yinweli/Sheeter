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

func TestProto(t *testing.T) {
	suite.Run(t, new(SuiteProto))
}

type SuiteProto struct {
	suite.Suite
	workDir   string
	excelName string
	sheetName string
}

func (this *SuiteProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excelName = "excelProto"
	this.sheetName = "sheetProto"
}

func (this *SuiteProto) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteProto) target() *Proto {
	target := &Proto{
		ExcelName: this.excelName,
		SheetName: this.sheetName,
	}
	return target
}

func (this *SuiteProto) TestName() {
	name := this.excelName + utils.FirstUpper(this.sheetName)
	structName := name + "."
	readerName := name + internal.Reader + "."
	protoCsPath := filepath.Join(internal.ProtoPath, internal.CsPath)
	protoGoPath := filepath.Join(internal.ProtoPath, internal.GoPath)
	protoSchemaPath := filepath.Join(internal.ProtoPath, internal.SchemaPath)
	protoName := structName + internal.ProtoSchemaExt
	protoPath := filepath.Join(protoSchemaPath, protoName)
	protoDataFile := structName + internal.ProtoDataExt
	protoDataPath := filepath.Join(internal.ProtoPath, internal.DataPath, protoDataFile)
	protoReaderCsPath := filepath.Join(internal.ProtoPath, internal.CsPath, utils.FirstUpper(readerName)+internal.CsExt)
	protoDepotCsPath := filepath.Join(internal.ProtoPath, internal.CsPath, utils.FirstUpper(internal.Depot)+"."+internal.CsExt)
	protoReaderGoPath := filepath.Join(internal.ProtoPath, internal.GoPath, readerName+internal.GoExt)
	protoDepotGoPath := filepath.Join(internal.ProtoPath, internal.GoPath, utils.FirstUpper(internal.Depot)+"."+internal.GoExt)
	protoDepend := utils.FirstLower(this.excelName) + "." + internal.ProtoSchemaExt

	target := this.target()
	assert.Equal(this.T(), protoCsPath, target.ProtoCsPath())
	assert.Equal(this.T(), protoGoPath, target.ProtoGoPath())
	assert.Equal(this.T(), protoSchemaPath, target.ProtoSchemaPath())
	assert.Equal(this.T(), protoName, target.ProtoName())
	assert.Equal(this.T(), protoPath, target.ProtoPath())
	assert.Equal(this.T(), name, target.ProtoDataName())
	assert.Equal(this.T(), internal.ProtoDataExt, target.ProtoDataExt())
	assert.Equal(this.T(), protoDataFile, target.ProtoDataFile())
	assert.Equal(this.T(), protoDataPath, target.ProtoDataPath())
	assert.Equal(this.T(), protoReaderCsPath, target.ProtoReaderCsPath())
	assert.Equal(this.T(), protoDepotCsPath, target.ProtoDepotCsPath())
	assert.Equal(this.T(), protoReaderGoPath, target.ProtoReaderGoPath())
	assert.Equal(this.T(), protoDepotGoPath, target.ProtoDepotGoPath())
	assert.Equal(this.T(), protoDepend, target.ProtoDepend(this.excelName))
}
