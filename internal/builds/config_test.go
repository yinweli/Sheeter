package builds

import (
	"strconv"
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
	workDir    string
	configFile Config
	configFlag Config
	elements   string
}

func (this *SuiteConfig) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.configFile = Config{
		Global: Global{
			LineOfField: 101,
			LineOfLayer: 102,
			LineOfNote:  103,
			LineOfData:  104,
		},
		Elements: []Element{
			{Excel: "excel1", Sheet: "sheet1"},
			{Excel: "excel2", Sheet: "sheet2"},
		},
	}
	this.configFlag = Config{
		Global: Global{
			LineOfField: 201,
			LineOfLayer: 202,
			LineOfNote:  203,
			LineOfData:  204,
		},
		Elements: []Element{
			{Excel: "excel3", Sheet: "sheet3"},
			{Excel: "excel4", Sheet: "sheet4"},
		},
	}
	this.elements = "excel3#sheet3,excel4#sheet4"
}

func (this *SuiteConfig) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteConfig) target() *Config {
	target := &Config{
		Global: Global{
			LineOfField: this.configFile.Global.LineOfField,
			LineOfLayer: this.configFile.Global.LineOfLayer,
			LineOfNote:  this.configFile.Global.LineOfNote,
			LineOfData:  this.configFile.Global.LineOfData,
		},
	}
	target.Elements = append(target.Elements, this.configFile.Elements...)
	return target
}

func (this *SuiteConfig) TestInitialize() {
	cmd := SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagConfig, testdata.ConfigNameReal))
	config := Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), this.configFile.Global.LineOfField, config.Global.LineOfField)
	assert.Equal(this.T(), this.configFile.Global.LineOfLayer, config.Global.LineOfLayer)
	assert.Equal(this.T(), this.configFile.Global.LineOfNote, config.Global.LineOfNote)
	assert.Equal(this.T(), this.configFile.Global.LineOfData, config.Global.LineOfData)
	assert.Equal(this.T(), this.configFile.Elements, config.Elements)

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfField, strconv.Itoa(this.configFlag.Global.LineOfField)))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfLayer, strconv.Itoa(this.configFlag.Global.LineOfLayer)))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfNote, strconv.Itoa(this.configFlag.Global.LineOfNote)))
	assert.Nil(this.T(), cmd.Flags().Set(flagLineOfData, strconv.Itoa(this.configFlag.Global.LineOfData)))
	assert.Nil(this.T(), cmd.Flags().Set(flagElements, this.elements))
	config = Config{}
	assert.Nil(this.T(), config.Initialize(cmd))
	assert.Equal(this.T(), this.configFlag.Global.LineOfField, config.Global.LineOfField)
	assert.Equal(this.T(), this.configFlag.Global.LineOfLayer, config.Global.LineOfLayer)
	assert.Equal(this.T(), this.configFlag.Global.LineOfNote, config.Global.LineOfNote)
	assert.Equal(this.T(), this.configFlag.Global.LineOfData, config.Global.LineOfData)
	assert.Equal(this.T(), this.configFlag.Elements, config.Elements)

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
}
