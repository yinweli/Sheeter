package util

import (
	"os"
	"path/filepath"
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
	workDir      string
	filePathReal string
	filePathFake string
	fileBytes    []byte
	fileBytesBom []byte
}

func (this *SuiteFileWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.filePathReal = "file/test.file"
	this.filePathFake = "?file/test.file"
	this.fileBytes = []byte("this is a string")
	this.fileBytesBom = bomPrefix
	this.fileBytesBom = append(this.fileBytesBom, this.fileBytes...)
}

func (this *SuiteFileWrite) TearDownSuite() {
	_ = os.RemoveAll(filepath.Dir(this.filePathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteFileWrite) TestFileWrite() {
	assert.Nil(this.T(), FileWrite(this.filePathReal, this.fileBytes, false))
	testdata.CompareFile(this.T(), this.filePathReal, this.fileBytes)

	assert.Nil(this.T(), FileWrite(this.filePathReal, this.fileBytes, true))
	testdata.CompareFile(this.T(), this.filePathReal, this.fileBytesBom)

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	assert.True(this.T(), testdata.IsWindows() && FileWrite(this.filePathFake, this.fileBytes, false) != nil)
}
