package builds

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepProto(t *testing.T) {
	suite.Run(t, new(SuitePoststepProto))
}

type SuitePoststepProto struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststepProto) TearDownSuite() {
	_ = os.Remove(internal.FileProtoCsBat)
	_ = os.Remove(internal.FileProtoCsSh)
	_ = os.Remove(internal.FileProtoGoBat)
	_ = os.Remove(internal.FileProtoGoSh)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepProto) target() *Runtime {
	target := &Runtime{
		Struct: []*RuntimeStruct{
			{Mixed: mixeds.NewMixed("test", "data")},
		},
	}
	return target
}

func (this *SuitePoststepProto) TestPoststepProtoCsBat() {
	proto := filepath.Join(internal.PathProto, internal.PathSchema)
	code := filepath.Join(internal.PathProto, internal.PathCs)
	file := filepath.Join(internal.PathProto, internal.PathSchema, "testData.proto")
	data := []byte(fmt.Sprintf(`REM generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./%s --csharp_out=./%s ./%s
`, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsBat(target))
	testdata.CompareFile(this.T(), internal.FileProtoCsBat, data)
}

func (this *SuitePoststepProto) TestPoststepProtoCsSh() {
	proto := filepath.Join(internal.PathProto, internal.PathSchema)
	code := filepath.Join(internal.PathProto, internal.PathCs)
	file := filepath.Join(internal.PathProto, internal.PathSchema, "testData.proto")
	data := []byte(fmt.Sprintf(`# generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./%s --csharp_out=./%s ./%s
`, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsSh(target))
	testdata.CompareFile(this.T(), internal.FileProtoCsSh, data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoBat() {
	proto := filepath.Join(internal.PathProto, internal.PathSchema)
	code := filepath.Join(internal.PathProto, internal.PathGo)
	file := filepath.Join(internal.PathProto, internal.PathSchema, "testData.proto")
	data := []byte(fmt.Sprintf(`REM generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./%s --go_out=./%s ./%s
`, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoBat(target))
	testdata.CompareFile(this.T(), internal.FileProtoGoBat, data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoSh() {
	proto := filepath.Join(internal.PathProto, internal.PathSchema)
	code := filepath.Join(internal.PathProto, internal.PathGo)
	file := filepath.Join(internal.PathProto, internal.PathSchema, "testData.proto")
	data := []byte(fmt.Sprintf(`# generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./%s --go_out=./%s ./%s
`, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoSh(target))
	testdata.CompareFile(this.T(), internal.FileProtoGoSh, data)
}
