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

func TestProto(t *testing.T) {
	suite.Run(t, new(SuiteProto))
}

type SuiteProto struct {
	suite.Suite
	testdata.TestEnv
	excelName string
	sheetName string
}

func (this *SuiteProto) SetupSuite() {
	this.Change("test-named-proto")
	this.excelName = "excelProto"
	this.sheetName = "sheetProto"
}

func (this *SuiteProto) TearDownSuite() {
	this.Restore()
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
	structName := name
	readerName := name + sheeter.Reader
	protoCsPath := filepath.Join(sheeter.ProtoPath, sheeter.CsPath)
	protoGoPath := filepath.Join(sheeter.ProtoPath, sheeter.GoPath)
	protoSchemaPath := filepath.Join(sheeter.ProtoPath, sheeter.SchemaPath)
	protoName := structName + sheeter.ProtoSchemaExt
	protoPath := filepath.Join(protoSchemaPath, protoName)
	protoDataFile := structName + sheeter.ProtoDataExt
	protoDataPath := filepath.Join(sheeter.ProtoPath, sheeter.DataPath, protoDataFile)
	protoReaderCsPath := filepath.Join(sheeter.ProtoPath, sheeter.CsPath, utils.FirstUpper(readerName)+sheeter.CsExt)
	protoDepotCsPath := filepath.Join(sheeter.ProtoPath, sheeter.CsPath, utils.FirstUpper(sheeter.Depot)+sheeter.CsExt)
	protoReaderGoPath := filepath.Join(sheeter.ProtoPath, sheeter.GoPath, readerName+sheeter.GoExt)
	protoDepotGoPath := filepath.Join(sheeter.ProtoPath, sheeter.GoPath, sheeter.Depot+sheeter.GoExt)
	protoDepend := utils.FirstLower(this.excelName) + sheeter.ProtoSchemaExt

	target := this.target()
	assert.Equal(this.T(), protoCsPath, target.ProtoCsPath())
	assert.Equal(this.T(), protoGoPath, target.ProtoGoPath())
	assert.Equal(this.T(), protoSchemaPath, target.ProtoSchemaPath())
	assert.Equal(this.T(), protoName, target.ProtoName())
	assert.Equal(this.T(), protoPath, target.ProtoPath())
	assert.Equal(this.T(), name, target.ProtoDataName())
	assert.Equal(this.T(), sheeter.ProtoDataExt, target.ProtoDataExt())
	assert.Equal(this.T(), protoDataFile, target.ProtoDataFile())
	assert.Equal(this.T(), protoDataPath, target.ProtoDataPath())
	assert.Equal(this.T(), protoReaderCsPath, target.ProtoReaderCsPath())
	assert.Equal(this.T(), protoDepotCsPath, target.ProtoDepotCsPath())
	assert.Equal(this.T(), protoReaderGoPath, target.ProtoReaderGoPath())
	assert.Equal(this.T(), protoDepotGoPath, target.ProtoDepotGoPath())
	assert.Equal(this.T(), protoDepend, target.ProtoDepend(this.excelName))
}
