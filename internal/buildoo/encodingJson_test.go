package buildoo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingJson(t *testing.T) {
	suite.Run(t, new(SuiteEncodingJson))
}

type SuiteEncodingJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingJson) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingJson) target(excel string) *encodingJson {
	// TODO: TestEncodingJson
	return nil
}
