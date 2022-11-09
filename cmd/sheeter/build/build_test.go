package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestBuild(t *testing.T) {
	suite.Run(t, new(SuiteBuild))
}

type SuiteBuild struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteBuild) SetupSuite() {
	this.Change("test-cmd-build")
}

func (this *SuiteBuild) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteBuild) TestExecute() {
	config := "config"
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealAll))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.DataPath, "realData.json"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.DataPath, "realData.bytes"))
	assert.FileExists(this.T(), filepath.Join(sheeter.EnumPath, sheeter.SchemaPath, "realEnum.proto"))
	_ = os.RemoveAll(sheeter.TmplPath)
	_ = os.RemoveAll(sheeter.JsonPath)
	_ = os.RemoveAll(sheeter.ProtoPath)
	_ = os.RemoveAll(sheeter.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealJson))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.DataPath, "realData.json"))
	_ = os.RemoveAll(sheeter.TmplPath)
	_ = os.RemoveAll(sheeter.JsonPath)
	_ = os.RemoveAll(sheeter.ProtoPath)
	_ = os.RemoveAll(sheeter.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealProto))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.DataPath, "realData.bytes"))
	_ = os.RemoveAll(sheeter.TmplPath)
	_ = os.RemoveAll(sheeter.JsonPath)
	_ = os.RemoveAll(sheeter.ProtoPath)
	_ = os.RemoveAll(sheeter.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealEnum))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.EnumPath, sheeter.SchemaPath, "realEnum.proto"))
	_ = os.RemoveAll(sheeter.TmplPath)
	_ = os.RemoveAll(sheeter.JsonPath)
	_ = os.RemoveAll(sheeter.ProtoPath)
	_ = os.RemoveAll(sheeter.EnumPath)
}
