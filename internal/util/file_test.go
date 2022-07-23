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
	workDir       string
	filePathReal  string
	filePathFake1 string
	filePathFake2 string
	fileBytes     []byte
	fileBytesBom  []byte
}

func (this *SuiteFileWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.filePathReal = "test/test.txt"
	this.filePathFake1 = "????.txt"
	this.filePathFake2 = "????/????.txt"
	this.fileBytes = []byte("this is a string")
	this.fileBytesBom = append(bomPrefix, this.fileBytes...)
}

func (this *SuiteFileWrite) TearDownSuite() {
	_ = os.RemoveAll(path.Dir(this.filePathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFileWrite) check(filepath string, expected []byte) {
	actual, err := os.ReadFile(filepath)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), expected, actual)
}

func (this *SuiteFileWrite) TestFileWrite() {
	assert.Nil(this.T(), FileWrite(this.filePathReal, this.fileBytes, false))
	this.check(this.filePathReal, this.fileBytes)
}

func (this *SuiteFileWrite) TestFileWriteBom() {
	assert.Nil(this.T(), FileWrite(this.filePathReal, this.fileBytes, true))
	this.check(this.filePathReal, this.fileBytesBom)
}

func (this *SuiteFileWrite) TestFileWriteFailed() {
	assert.NotNil(this.T(), FileWrite(this.filePathFake1, this.fileBytes, false))
	assert.NotNil(this.T(), FileWrite(this.filePathFake2, this.fileBytes, false))
}
