package build

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestBuild(t *testing.T) {
	suite.Run(t, new(SuiteBuild))
}

type SuiteBuild struct {
	suite.Suite
	testdata.TestEnv
	config string
}

func (this *SuiteBuild) SetupSuite() {
	this.TBegin("test-cmd-build", "build")
	this.config = "config.yaml"
}

func (this *SuiteBuild) TearDownSuite() {
	excels.CloseAll()
	this.TFinal()
}

func (this *SuiteBuild) TestNewCommand() {
	assert.NotNil(this.T(), NewCommand())
}

func (this *SuiteBuild) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set("config", this.config))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.CsPath, "ExcelTestReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.CsPath, "Sheeter.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.GoPath, "excelTestReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.GoPath, "sheeter.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, "excelTest.json"))
}
