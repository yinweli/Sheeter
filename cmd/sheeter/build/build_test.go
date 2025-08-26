package build

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestBuild(t *testing.T) {
	suite.Run(t, new(SuiteBuild))
}

type SuiteBuild struct {
	suite.Suite
	testdata.Env
	config string
}

func (this *SuiteBuild) SetupSuite() {
	this.Env = testdata.EnvSetup("test-cmd-build", "build")
	this.config = "config.yaml"
}

func (this *SuiteBuild) TearDownSuite() {
	excels.CloseAll()
	testdata.EnvRestore(this.Env)
}

func (this *SuiteBuild) TestNewCommand() {
	assert.NotNil(this.T(), NewCommand())
}

func (this *SuiteBuild) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set("config", this.config))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.PathCs, "ExcelTestReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.PathCs, "Sheeter.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.PathGo, "excelTestReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.PathGo, "sheeter.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.PathJson, "excelTest.json"))
}
