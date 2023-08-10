package nameds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestNamed(t *testing.T) {
	suite.Run(t, new(SuiteNamed))
}

type SuiteNamed struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteNamed) SetupSuite() {
	this.Env = testdata.EnvSetup("test-nameds-named")
}

func (this *SuiteNamed) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteNamed) TestName() {
	target := &Named{
		Output:    "output",
		ExcelName: "excel",
		SheetName: "sheet",
	}
	assert.Equal(this.T(), sheeter.AppName, target.AppName())
	assert.Equal(this.T(), sheeter.Namespace, target.Namespace())
	assert.Equal(this.T(), "ExcelSheet", target.StructName())
	assert.Equal(this.T(), "excel # sheet", target.StructNote())
	assert.Equal(this.T(), "ExcelSheetReader", target.ReaderName())
	assert.Equal(this.T(), "excelSheet", target.JsonName())
	assert.Equal(this.T(), sheeter.JsonExt, target.JsonExt())
	assert.Equal(this.T(), "excelSheet.json", target.DataFile())
	assert.Equal(this.T(), filepath.Join("output", sheeter.JsonPath, "excelSheet.json"), target.DataPath())
	assert.Equal(this.T(), filepath.Join("output", sheeter.CsPath, "ExcelSheetReader.cs"), target.ReaderPathCs())
	assert.Equal(this.T(), filepath.Join("output", sheeter.CsPath, "Sheeter.cs"), target.SheeterPathCs())
	assert.Equal(this.T(), filepath.Join("output", sheeter.GoPath, "excelSheetReader.go"), target.ReaderPathGo())
	assert.Equal(this.T(), filepath.Join("output", sheeter.GoPath, "sheeter.go"), target.SheeterPathGo())
	assert.Equal(this.T(), "TestString", target.FirstUpper("testString"))
	assert.Equal(this.T(), "testString", target.FirstLower("TestString"))
}
