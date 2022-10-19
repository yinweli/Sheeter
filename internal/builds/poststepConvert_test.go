package builds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepConvert(t *testing.T) {
	suite.Run(t, new(SuitePoststepConvert))
}

type SuitePoststepConvert struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepConvert) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	_ = os.MkdirAll(internal.CsPath, os.ModePerm)
	_ = os.MkdirAll(internal.GoPath, os.ModePerm)
}

func (this *SuitePoststepConvert) TearDownSuite() {
	_ = os.RemoveAll(internal.CsPath)
	_ = os.RemoveAll(internal.GoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepConvert) target() *poststepConvert {
	target := &poststepConvert{
		include:  "",
		outputCs: internal.CsPath,
		outputGo: internal.GoPath,
		source:   testdata.ProtoNameTest,
	}
	return target
}

func (this *SuitePoststepConvert) TestPoststepConvertCs() {
	assert.Nil(this.T(), PoststepConvertCs(this.target()))
	assert.FileExists(this.T(), filepath.Join(internal.CsPath, "test1.cs"))

	assert.Nil(this.T(), PoststepConvertCs(nil))
}

func (this *SuitePoststepConvert) TestPoststepConvertGo() {
	assert.Nil(this.T(), PoststepConvertGo(this.target()))
	assert.FileExists(this.T(), filepath.Join(internal.GoPath, "test1.pb.go")) // proto轉換go都會帶pb後綴

	assert.Nil(this.T(), PoststepConvertGo(nil))
}
