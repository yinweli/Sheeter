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

func TestEnum(t *testing.T) {
	suite.Run(t, new(SuiteEnum))
}

type SuiteEnum struct {
	suite.Suite
	testdata.TestEnv
	excelName string
	sheetName string
}

func (this *SuiteEnum) SetupSuite() {
	this.Change("test-named-enum")
	this.excelName = "excelEnum"
	this.sheetName = "sheetEnum"
}

func (this *SuiteEnum) TearDownSuite() {
	this.Restore()
}

func (this *SuiteEnum) target() *Enum {
	target := &Enum{
		ExcelName: this.excelName,
		SheetName: this.sheetName,
	}
	return target
}

func (this *SuiteEnum) TestName() {
	name := this.excelName + utils.FirstUpper(this.sheetName)
	structName := name
	protoCsPath := filepath.Join(internal.EnumPath, internal.CsPath)
	protoGoPath := filepath.Join(internal.EnumPath, internal.GoPath)
	protoSchemaPath := filepath.Join(internal.EnumPath, internal.SchemaPath)
	protoName := structName + internal.EnumSchemaExt
	protoPath := filepath.Join(protoSchemaPath, protoName)

	target := this.target()
	assert.Equal(this.T(), protoCsPath, target.EnumCsPath())
	assert.Equal(this.T(), protoGoPath, target.EnumGoPath())
	assert.Equal(this.T(), protoSchemaPath, target.EnumSchemaPath())
	assert.Equal(this.T(), protoName, target.EnumName())
	assert.Equal(this.T(), protoPath, target.EnumPath())
}
