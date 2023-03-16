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
	testdata.TestData
	configSuccess string
	configFailed  string
	path1         string
	path2         string
	path3         string
	path4         string
}

func (this *SuiteConfig) SetupSuite() {
	this.TBegin("test-builds-config", "config")
	this.configSuccess = "success.yaml"
	this.configFailed = "failed.yaml"
	this.path1 = "path"
	this.path2 = "path/file.xlsx"
	this.path3 = "path/file.fail"
	this.path4 = "????"
}

func (this *SuiteConfig) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteConfig) TestInitialize() {
	cmd := SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, this.configSuccess))
	config := Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), []string{"path", "path/path", "path/path.xlsx"}, config.Source)
	assert.Equal(this.T(), "output", config.Output)
	assert.Equal(this.T(), "tag", config.Global.Tag)
	assert.Equal(this.T(), false, config.Global.AutoKey)
	assert.Equal(this.T(), 101, config.Global.LineOfTag)
	assert.Equal(this.T(), 102, config.Global.LineOfName)
	assert.Equal(this.T(), 103, config.Global.LineOfNote)
	assert.Equal(this.T(), 104, config.Global.LineOfField)
	assert.Equal(this.T(), 105, config.Global.LineOfData)

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagSource, "path/excel1.xlsx,path/excel2.xlsx"))
	assert.Nil(this.T(), cmd.Flags().Set(flagOutput, "path/output"))
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
	assert.Equal(this.T(), "TAG", config.Global.Tag)
	assert.Equal(this.T(), true, config.Global.AutoKey)
	assert.Equal(this.T(), 201, config.Global.LineOfTag)
	assert.Equal(this.T(), 202, config.Global.LineOfName)
	assert.Equal(this.T(), 203, config.Global.LineOfNote)
	assert.Equal(this.T(), 204, config.Global.LineOfField)
	assert.Equal(this.T(), 205, config.Global.LineOfData)

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, this.configFailed))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))

	cmd = SetFlag(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, this.Unknown))
	config = Config{}
	assert.NotNil(this.T(), config.Initialize(cmd))
}

func (this *SuiteConfig) TestFile() {
	config := Config{Source: []string{this.path1, this.path2, this.path3, this.path4}}
	assert.Equal(this.T(), []string{this.path2}, config.File())
}

func (this *SuiteConfig) TestPath() {
	config := Config{Source: []string{this.path1, this.path2, this.path3, this.path4}}
	assert.Equal(this.T(), []string{this.path1}, config.Path())
}

func (this *SuiteConfig) TestCheck() {
	original := Config{
		Global: Global{
			LineOfTag:   1,
			LineOfName:  2,
			LineOfNote:  3,
			LineOfField: 4,
			LineOfData:  5,
		},
	}

	target := original
	assert.Nil(this.T(), target.Check())

	target = original
	target.Global.LineOfTag = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfName = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfTag = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfName = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfNote = 999
	assert.NotNil(this.T(), target.Check())

	target = original
	target.Global.LineOfField = 999
	assert.NotNil(this.T(), target.Check())
}
