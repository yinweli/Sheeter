package util

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFileWrite(t *testing.T) {
	suite.Run(t, new(SuiteFileWrite))
}

type SuiteFileWrite struct {
	suite.Suite
	workDir          string
	realFilePath     string
	realFileBytes    []byte
	realFileBytesBom []byte
	fakeFilePath1    string
	fakeFilePath2    string
}

func (this *SuiteFileWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.realFilePath = "test/test.txt"
	this.realFileBytes = []byte("this is a string")
	this.realFileBytesBom = append(bomprefix, this.realFileBytes...)
	this.fakeFilePath1 = "????.txt"
	this.fakeFilePath2 = "????/????.txt"
}

func (this *SuiteFileWrite) TearDownSuite() {
	_ = os.RemoveAll(path.Dir(this.realFilePath))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFileWrite) check(filepath string, expected []byte) {
	actual, err := os.ReadFile(filepath)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), expected, actual)
}

func (this *SuiteFileWrite) TestFileWrite() {
	assert.Nil(this.T(), FileWrite(this.realFilePath, this.realFileBytes, false))
	this.check(this.realFilePath, this.realFileBytes)
}

func (this *SuiteFileWrite) TestFileWriteBom() {
	assert.Nil(this.T(), FileWrite(this.realFilePath, this.realFileBytes, true))
	this.check(this.realFilePath, this.realFileBytesBom)
}

func (this *SuiteFileWrite) TestFileWriteFailed() {
	assert.NotNil(this.T(), FileWrite(this.fakeFilePath1, this.realFileBytes, false))
	assert.NotNil(this.T(), FileWrite(this.fakeFilePath2, this.realFileBytes, false))
}
