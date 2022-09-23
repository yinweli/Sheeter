package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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
	_ = os.Remove(this.target().FileProtoCsBat())
	_ = os.Remove(this.target().FileProtoCsSh())
	_ = os.Remove(this.target().FileProtoGoBat())
	_ = os.Remove(this.target().FileProtoGoSh())
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepProto) target() *Runtime {
	target := &Runtime{
		Struct: []*RuntimeStruct{
			{Mixed: mixeds.NewMixed("test", "data1")},
			{Mixed: mixeds.NewMixed("test", "data2")},
		},
	}
	return target
}

func (this *SuitePoststepProto) TestPoststepProtoCsBat() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./schema --csharp_out=./cs ./schema\testData1.proto
protoc --experimental_allow_proto3_optional --proto_path=./schema --csharp_out=./cs ./schema\testData2.proto
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsBat(target))
	testdata.CompareFile(this.T(), target.FileProtoCsBat(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoCsSh() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./schema --csharp_out=./cs ./schema\testData1.proto
protoc --experimental_allow_proto3_optional --proto_path=./schema --csharp_out=./cs ./schema\testData2.proto
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsSh(target))
	testdata.CompareFile(this.T(), target.FileProtoCsSh(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoBat() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./schema --go_out=./go ./schema\testData1.proto
protoc --experimental_allow_proto3_optional --proto_path=./schema --go_out=./go ./schema\testData2.proto
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoBat(target))
	testdata.CompareFile(this.T(), target.FileProtoGoBat(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoSh() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.
protoc --experimental_allow_proto3_optional --proto_path=./schema --go_out=./go ./schema\testData1.proto
protoc --experimental_allow_proto3_optional --proto_path=./schema --go_out=./go ./schema\testData2.proto
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoSh(target))
	testdata.CompareFile(this.T(), target.FileProtoGoSh(), data)
}
