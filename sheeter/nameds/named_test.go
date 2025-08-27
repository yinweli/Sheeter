package nameds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/fields"
	"github.com/yinweli/Sheeter/v3/sheeter/layouts"
	"github.com/yinweli/Sheeter/v3/testdata"
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
		Primary: &layouts.LayoutData{
			Field: &fields.Int{},
		},
	}
	assert.Equal(this.T(), sheeter.Application, target.AppName())
	assert.Equal(this.T(), sheeter.Namespace, target.Namespace())
	assert.Equal(this.T(), "ExcelSheet", target.StructName())
	assert.Equal(this.T(), "excel#sheet", target.StructNote())
	assert.Equal(this.T(), "ExcelSheetReader", target.ReaderName())
	assert.Equal(this.T(), "excelSheet", target.JsonName())
	assert.Equal(this.T(), sheeter.ExtJson, target.JsonExt())
	assert.Equal(this.T(), "excelSheet.json", target.DataFile())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathJson, "excelSheet.json"), target.DataPath())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathCs, "ExcelSheetReader.cs"), target.ReaderPathCs())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathCs, "Sheeter.cs"), target.SheeterPathCs())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathCs, "Helper.cs"), target.HelperPathCs())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathGo, "excelSheetReader.go"), target.ReaderPathGo())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathGo, "sheeter.go"), target.SheeterPathGo())
	assert.Equal(this.T(), filepath.Join("output", sheeter.PathGo, "helper.go"), target.HelperPathGo())
	assert.Equal(this.T(), "Int32", target.PrimaryCs())
	assert.Equal(this.T(), "int32", target.PrimaryGo())
	assert.Equal(this.T(), "TestString", target.FirstUpper("testString"))
	assert.Equal(this.T(), "testString", target.FirstLower("TestString"))
}
