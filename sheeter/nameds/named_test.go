package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestNamed(t *testing.T) {
	suite.Run(t, new(SuiteNamed))
}

type SuiteNamed struct {
	suite.Suite
	testdata.TestEnv
	excelName string
	sheetName string
}

func (this *SuiteNamed) SetupSuite() {
	this.Change("test-named")
	this.excelName = "excel"
	this.sheetName = "sheet"
}

func (this *SuiteNamed) TearDownSuite() {
	this.Restore()
}

func (this *SuiteNamed) target() *Named {
	target := &Named{
		ExcelName: this.excelName,
		SheetName: this.sheetName,
	}
	return target
}

func (this *SuiteNamed) TestName() {
	structName := utils.FirstUpper(this.excelName) + utils.FirstUpper(this.sheetName)
	readerName := structName + sheeter.Reader
	storerName := structName + sheeter.Storer

	target := this.target()
	assert.Equal(this.T(), sheeter.AppName, target.AppName())
	assert.Equal(this.T(), sheeter.AppName, target.JsonNamespace(true))
	assert.Equal(this.T(), sheeter.JsonNamespace, target.JsonNamespace(false))
	assert.Equal(this.T(), sheeter.AppName, target.ProtoNamespace(true))
	assert.Equal(this.T(), sheeter.ProtoNamespace, target.ProtoNamespace(false))
	assert.Equal(this.T(), sheeter.AppName, target.EnumNamespace(true))
	assert.Equal(this.T(), sheeter.EnumNamespace, target.EnumNamespace(false))
	assert.Equal(this.T(), structName, target.StructName())
	assert.Equal(this.T(), readerName, target.ReaderName())
	assert.Equal(this.T(), storerName, target.StorerName())
	assert.Equal(this.T(), sheeter.StorerDatas, target.StorerDatas())
	assert.Equal(this.T(), sheeter.AppName+"."+storerName, target.StorerMessage(true))
	assert.Equal(this.T(), sheeter.ProtoNamespace+"."+storerName, target.StorerMessage(false))
	assert.Equal(this.T(), "TestString", target.FirstUpper("testString"))
	assert.Equal(this.T(), "testString", target.FirstLower("TestString"))
	assert.Equal(this.T(), "8", target.Add(6, 2))
	assert.Equal(this.T(), "8", target.Add(2, 6))
	assert.Equal(this.T(), "4", target.Sub(6, 2))
	assert.Equal(this.T(), "-4", target.Sub(2, 6))
	assert.Equal(this.T(), "12", target.Mul(6, 2))
	assert.Equal(this.T(), "12", target.Mul(2, 6))
	assert.Equal(this.T(), "3", target.Div(6, 2))
	assert.Equal(this.T(), "0", target.Div(2, 6))
}
