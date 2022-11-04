package builds

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
	testdata.TestEnv
	excel     string
	sheet     string
	path      string
	pathx     string
	file      string
	filex     string
	fileExcel string
	fileSheet string
	empty     string
}

func (this *SuiteConfig) SetupSuite() {
	this.Change("test-config")
	this.excel = "excel"
	this.sheet = "sheet"
	this.path = "path"
	this.pathx = this.path + internal.SeparateSheet + "x"
	this.file = filepath.Join(this.path, "file.x")
	this.filex = filepath.Join(this.path, "file.x"+internal.SeparateSheet+"x")
	this.fileExcel = filepath.Join(this.path, this.excel+internal.ExcelExt)
	this.fileSheet = this.fileExcel + internal.SeparateSheet + this.sheet
	this.empty = "empty" + internal.SeparateSheet + "x"
	_ = os.MkdirAll(this.path, os.ModePerm)
	_ = os.WriteFile(this.file, []byte{}, fs.ModePerm)
	_ = os.WriteFile(this.fileExcel, []byte{}, fs.ModePerm)
}

func (this *SuiteConfig) TearDownSuite() {
	this.Restore()
}

func (this *SuiteConfig) target() *Config {
	target := &Config{
		Global: Global{
			LineOfName:  1,
			LineOfNote:  2,
			LineOfField: 3,
			LineOfLayer: 4,
			LineOfData:  5,
			LineOfEnum:  6,
		},
	}
	return target
}

func (this *SuiteConfig) TestInitialize() {
	cmd := SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.ConfigReal))
	config := Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), true, config.Global.ExportJson)
	assert.Equal(this.T(), true, config.Global.ExportProto)
	assert.Equal(this.T(), false, config.Global.ExportEnum)
	assert.Equal(this.T(), false, config.Global.SimpleNamespace)
	assert.Equal(this.T(), 101, config.Global.LineOfName)
	assert.Equal(this.T(), 102, config.Global.LineOfNote)
	assert.Equal(this.T(), 103, config.Global.LineOfField)
	assert.Equal(this.T(), 104, config.Global.LineOfLayer)
	assert.Equal(this.T(), 105, config.Global.LineOfData)
	assert.Equal(this.T(), 106, config.Global.LineOfEnum)
	assert.Equal(this.T(), []string{"tag1", "tag2"}, config.Global.Excludes)
	assert.Equal(this.T(), []string{"path", "path/path", "path/path.xlsx", "path/path.xlsx#sheet"}, config.Inputs)

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagExportJson, "true"))
	assert.Nil(this.T(), cmd.Flags().Set(flagExportProto, "true"))
	assert.Nil(this.T(), cmd.Flags().Set(flagExportEnum, "true"))
	assert.Nil(this.T(), cmd.Flags().Set(flagSimpleNamespace, "true"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfName, "201"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfNote, "202"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfField, "203"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfLayer, "204"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfData, "205"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfEnum, "206"))
	assert.Nil(this.T(), cmd.Flags().Set(flagExcludes, "tag3,tag4"))
	assert.Nil(this.T(), cmd.Flags().Set(flagInputs, "excel1#sheet1,excel2#sheet2"))
	config = Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), true, config.Global.ExportJson)
	assert.Equal(this.T(), true, config.Global.ExportProto)
	assert.Equal(this.T(), true, config.Global.ExportEnum)
	assert.Equal(this.T(), true, config.Global.SimpleNamespace)
	assert.Equal(this.T(), 201, config.Global.LineOfName)
	assert.Equal(this.T(), 202, config.Global.LineOfNote)
	assert.Equal(this.T(), 203, config.Global.LineOfField)
	assert.Equal(this.T(), 204, config.Global.LineOfLayer)
	assert.Equal(this.T(), 205, config.Global.LineOfData)
	assert.Equal(this.T(), 206, config.Global.LineOfEnum)
	assert.Equal(this.T(), []string{"tag3", "tag4"}, config.Global.Excludes)
	assert.Equal(this.T(), []string{"excel1#sheet1", "excel2#sheet2"}, config.Inputs)

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.ConfigFake))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.UnknownStr))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))
}

func (this *SuiteConfig) TestPath() {
	config := Config{Inputs: []string{this.path}}
	assert.Equal(this.T(), []string{this.path}, config.Path())

	config = Config{Inputs: []string{this.file}}
	assert.Empty(this.T(), config.Path())

	config = Config{Inputs: []string{this.empty}}
	assert.Empty(this.T(), config.Path())
}

func (this *SuiteConfig) TestExcel() {
	config := Config{Inputs: []string{this.fileExcel}}
	assert.Equal(this.T(), []string{this.fileExcel}, config.Excel())

	config = Config{Inputs: []string{this.path}}
	assert.Empty(this.T(), config.Excel())

	config = Config{Inputs: []string{this.file}}
	assert.Empty(this.T(), config.Excel())

	config = Config{Inputs: []string{this.empty}}
	assert.Empty(this.T(), config.Excel())
}

func (this *SuiteConfig) TestSheet() {
	config := Config{Inputs: []string{this.fileSheet}}
	assert.Equal(this.T(), []Sheet{{ExcelName: this.fileExcel, SheetName: this.sheet}}, config.Sheet())

	config = Config{Inputs: []string{this.path}}
	assert.Empty(this.T(), config.Sheet())

	config = Config{Inputs: []string{this.pathx}}
	assert.Empty(this.T(), config.Sheet())

	config = Config{Inputs: []string{this.file}}
	assert.Empty(this.T(), config.Sheet())

	config = Config{Inputs: []string{this.filex}}
	assert.Empty(this.T(), config.Sheet())

	config = Config{Inputs: []string{this.empty}}
	assert.Empty(this.T(), config.Sheet())
}

func (this *SuiteConfig) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfName = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfLayer = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfEnum = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfName = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfField = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfLayer = 999
	assert.NotNil(this.T(), target.Check())
}
