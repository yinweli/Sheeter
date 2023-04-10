package builds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
	testdata.Env
	configSuccess string
	configFailed  string
}

func (this *SuiteConfig) SetupSuite() {
	this.Env = testdata.EnvSetup("test-builds-config", "config")
	this.configSuccess = "success.yaml"
	this.configFailed = "failed.yaml"
}

func (this *SuiteConfig) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteConfig) TestInitialize() {
	cmd := SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, this.configSuccess))
	config := Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), []string{"path", "path/path", "path/path.xlsx"}, config.Source)
	assert.Equal(this.T(), "output", config.Output)
	assert.Equal(this.T(), []string{"excel1#sheet", "excel2#sheet"}, config.Exclude)
	assert.Equal(this.T(), "tag", config.Tag)
	assert.Equal(this.T(), false, config.AutoKey)
	assert.Equal(this.T(), 101, config.LineOfTag)
	assert.Equal(this.T(), 102, config.LineOfName)
	assert.Equal(this.T(), 103, config.LineOfNote)
	assert.Equal(this.T(), 104, config.LineOfField)
	assert.Equal(this.T(), 105, config.LineOfData)

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagSource, "path/excel1.xlsx,path/excel2.xlsx"))
	assert.Nil(this.T(), cmd.Flags().Set(flagOutput, "path/output"))
	assert.Nil(this.T(), cmd.Flags().Set(flagExclude, "excel3#sheet,excel4#sheet"))
	assert.Nil(this.T(), cmd.Flags().Set(flagTag, "TAG"))
	assert.Nil(this.T(), cmd.Flags().Set(flagAutoKey, "true"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfTag, "201"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfName, "202"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfNote, "203"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfField, "204"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfData, "205"))
	config = Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), []string{"path/excel1.xlsx", "path/excel2.xlsx"}, config.Source)
	assert.Equal(this.T(), "path/output", config.Output)
	assert.Equal(this.T(), []string{"excel3#sheet", "excel4#sheet"}, config.Exclude)
	assert.Equal(this.T(), "TAG", config.Tag)
	assert.Equal(this.T(), true, config.AutoKey)
	assert.Equal(this.T(), 201, config.LineOfTag)
	assert.Equal(this.T(), 202, config.LineOfName)
	assert.Equal(this.T(), 203, config.LineOfNote)
	assert.Equal(this.T(), 204, config.LineOfField)
	assert.Equal(this.T(), 205, config.LineOfData)

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, this.configFailed))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.Unknown))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))
}

func (this *SuiteConfig) TestFile() {
	config := Config{Source: []string{"path", "path/file.xlsx", "path/file.fail", "????"}}
	assert.Equal(this.T(), []string{"path/file.xlsx"}, config.File())
}

func (this *SuiteConfig) TestPath() {
	config := Config{Source: []string{"path", "path/file.xlsx", "path/file.fail", "????"}}
	assert.Equal(this.T(), []string{"path"}, config.Path())
}

func (this *SuiteConfig) TestExamine() {
	config := Config{Exclude: []string{"excel1#sheet1", "excel2#sheet2"}}
	assert.True(this.T(), config.Examine("excel1", "sheet1"))
	assert.True(this.T(), config.Examine("excel2", "sheet2"))
	assert.False(this.T(), config.Examine("excel3", "sheet3"))
}

func (this *SuiteConfig) TestCheck() {
	original := Config{
		LineOfTag:   1,
		LineOfName:  2,
		LineOfNote:  3,
		LineOfField: 4,
		LineOfData:  5,
	}

	target := original
	assert.Nil(this.T(), target.Check())

	target = original
	target.LineOfTag = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfName = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfTag = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfName = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfNote = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.LineOfField = 999
	assert.NotNil(this.T(), target.Check())
}
