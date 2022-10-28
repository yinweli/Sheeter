package builds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
	workDir string
}

func (this *SuiteConfig) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteConfig) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
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
