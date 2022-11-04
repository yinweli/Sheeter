package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
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
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.DataPath, "realData.json"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.DataPath, "realData.bytes"))
	assert.FileExists(this.T(), filepath.Join(internal.EnumPath, internal.SchemaPath, "realEnum.proto"))
	_ = os.RemoveAll(internal.TmplPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	_ = os.RemoveAll(internal.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealJson))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.DataPath, "realData.json"))
	_ = os.RemoveAll(internal.TmplPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	_ = os.RemoveAll(internal.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealProto))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.DataPath, "realData.bytes"))
	_ = os.RemoveAll(internal.TmplPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	_ = os.RemoveAll(internal.EnumPath)

	cmd = NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigRealEnum))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(internal.EnumPath, internal.SchemaPath, "realEnum.proto"))
	_ = os.RemoveAll(internal.TmplPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	_ = os.RemoveAll(internal.EnumPath)
}
