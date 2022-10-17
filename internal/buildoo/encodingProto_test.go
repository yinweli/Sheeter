package buildoo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingProto(t *testing.T) {
	suite.Run(t, new(SuiteEncodingProto))
}

type SuiteEncodingProto struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingProto) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingProto) target(excel string) *encodingProto {
	// TODO: TestEncodingProto
	return nil
}
