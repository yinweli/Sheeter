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
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Elements: []Element{
			{Excel: "excel1", Sheet: "sheet1"},
			{Excel: "excel2", Sheet: "sheet2"},
		},
	}
	return target
}

func (this *SuiteConfig) TestInitialize() {
	cmd := SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.ConfigNameReal))
	config := Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), 101, config.Global.LineOfField)
	assert.Equal(this.T(), 102, config.Global.LineOfLayer)
	assert.Equal(this.T(), 103, config.Global.LineOfNote)
	assert.Equal(this.T(), 104, config.Global.LineOfData)
	assert.Equal(this.T(), []string{"tag1", "tag2"}, config.Global.Excludes)
	assert.Equal(this.T(), []Element{{Excel: "excel1", Sheet: "sheet1"}, {Excel: "excel2", Sheet: "sheet2"}}, config.Elements)

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfField, "201"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfLayer, "202"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfNote, "203"))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfData, "204"))
	assert.Nil(this.T(), cmd.Flags().Set(flagExcludes, "tag3,tag4"))
	assert.Nil(this.T(), cmd.Flags().Set(flagElements, "excel3#sheet3,excel4#sheet4"))
	config = Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), 201, config.Global.LineOfField)
	assert.Equal(this.T(), 202, config.Global.LineOfLayer)
	assert.Equal(this.T(), 203, config.Global.LineOfNote)
	assert.Equal(this.T(), 204, config.Global.LineOfData)
	assert.Equal(this.T(), []string{"tag3", "tag4"}, config.Global.Excludes)
	assert.Equal(this.T(), []Element{{Excel: "excel3", Sheet: "sheet3"}, {Excel: "excel4", Sheet: "sheet4"}}, config.Elements)

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.ConfigNameFake))
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
	target.Global.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfLayer = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfField = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfLayer = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 999
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[0].Excel = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[0].Sheet = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[1].Excel = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[1].Sheet = ""
	assert.NotNil(this.T(), target.Check())
}
