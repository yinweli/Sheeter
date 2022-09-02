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
	workDir     string
	bom         bool
	lineOfField int
	lineOfLayer int
	lineOfNote  int
	lineOfData  int
	element     []Element
}

func (this *SuiteConfig) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.bom = true
	this.lineOfField = 101
	this.lineOfLayer = 102
	this.lineOfNote = 103
	this.lineOfData = 104
	this.element = []Element{
		{"excel3", "sheet"},
		{"excel2", "sheet"},
		{"excel1", "sheet"},
	}
}

func (this *SuiteConfig) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteConfig) target() *Config {
	target := &Config{
		Global: Global{
			Bom:         this.bom,
			LineOfField: this.lineOfField,
			LineOfLayer: this.lineOfLayer,
			LineOfNote:  this.lineOfNote,
			LineOfData:  this.lineOfData,
		},
	}
	target.Elements = append(target.Elements, this.element...)
	return target
}

func (this *SuiteConfig) cmd(filename string) *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Flags().String(FlagConfig, filename, "")
	cmd.Flags().Bool(FlagBom, this.bom, "")
	cmd.Flags().Int(FlagLineOfField, this.lineOfField, "")
	cmd.Flags().Int(FlagLineOfLayer, this.lineOfLayer, "")
	cmd.Flags().Int(FlagLineOfNote, this.lineOfNote, "")
	cmd.Flags().Int(FlagLineOfData, this.lineOfData, "")
	cmd.Flags().StringSlice(FlagElements, []string{this.element[0].Excel + separateElement + this.element[0].Sheet}, "")
	return cmd
}

func (this *SuiteConfig) TestNewConfig() {
	config, err := NewConfig(this.cmd(testdata.ConfigNameReal))
	assert.NotNil(this.T(), config)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), this.bom, config.Global.Bom)
	assert.Equal(this.T(), this.lineOfField, config.Global.LineOfField)
	assert.Equal(this.T(), this.lineOfLayer, config.Global.LineOfLayer)
	assert.Equal(this.T(), this.lineOfNote, config.Global.LineOfNote)
	assert.Equal(this.T(), this.lineOfData, config.Global.LineOfData)
	assert.Equal(this.T(), this.element[0].Excel, config.Elements[0].Excel)
	assert.Equal(this.T(), this.element[0].Sheet, config.Elements[0].Sheet)

	config, err = NewConfig(this.cmd(testdata.UnknownStr))
	assert.Nil(this.T(), config)
	assert.NotNil(this.T(), err)

	config, err = NewConfig(this.cmd(testdata.ConfigNameFake))
	assert.Nil(this.T(), config)
	assert.NotNil(this.T(), err)
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

func (this *SuiteConfig) TestToContents() {
	target := this.target()
	contents := target.ToContents()
	assert.Equal(this.T(), this.bom, contents.Contents[0].Bom)
	assert.Equal(this.T(), this.lineOfField, contents.Contents[0].LineOfField)
	assert.Equal(this.T(), this.lineOfLayer, contents.Contents[0].LineOfLayer)
	assert.Equal(this.T(), this.lineOfNote, contents.Contents[0].LineOfNote)
	assert.Equal(this.T(), this.lineOfData, contents.Contents[0].LineOfData)
	assert.Equal(this.T(), this.element[2].Excel, contents.Contents[0].Excel)
	assert.Equal(this.T(), this.element[2].Sheet, contents.Contents[0].Sheet)
	assert.Equal(this.T(), this.element[1].Excel, contents.Contents[1].Excel)
	assert.Equal(this.T(), this.element[1].Sheet, contents.Contents[1].Sheet)
	assert.Equal(this.T(), this.element[0].Excel, contents.Contents[2].Excel)
	assert.Equal(this.T(), this.element[0].Sheet, contents.Contents[2].Sheet)
	assert.Len(this.T(), contents.Contents, 3)
}
