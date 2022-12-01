package builds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepConvert(t *testing.T) {
	suite.Run(t, new(SuitePoststepConvert))
}

type SuitePoststepConvert struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuitePoststepConvert) SetupSuite() {
	this.Change("test-poststepConvert")
	_ = os.MkdirAll(sheeter.CsPath, os.ModePerm)
	_ = os.MkdirAll(sheeter.GoPath, os.ModePerm)
}

func (this *SuitePoststepConvert) TearDownSuite() {
	this.Restore()
}

func (this *SuitePoststepConvert) target() *poststepConvert {
	target := &poststepConvert{
		include:  "",
		outputCs: sheeter.CsPath,
		outputGo: sheeter.GoPath,
		source:   testdata.ProtoTest,
	}
	return target
}

func (this *SuitePoststepConvert) TestPoststepConvertCs() {
	assert.Nil(this.T(), PoststepConvertCs(this.target(), nil))
	assert.FileExists(this.T(), filepath.Join(sheeter.CsPath, "Test1.cs"))

	assert.Nil(this.T(), PoststepConvertCs(nil, nil))
}

func (this *SuitePoststepConvert) TestPoststepConvertGo() {
	assert.Nil(this.T(), PoststepConvertGo(this.target(), nil))
	assert.FileExists(this.T(), filepath.Join(sheeter.GoPath, "test1.pb.go")) // proto轉換go都會帶pb後綴

	assert.Nil(this.T(), PoststepConvertGo(nil, nil))
}