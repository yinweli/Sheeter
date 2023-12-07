package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestWrite(t *testing.T) {
	suite.Run(t, new(SuiteWrite))
}

type SuiteWrite struct {
	suite.Suite
	testdata.Env
	fileExist string
	pathExist string
}

func (this *SuiteWrite) SetupSuite() {
	this.Env = testdata.EnvSetup("test-utils-write", "write")
	this.fileExist = "exist.txt"
	this.pathExist, _ = os.Getwd()
}

func (this *SuiteWrite) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteWrite) TestFileName() {
	assert.Equal(this.T(), "test", FileName(filepath.Join("dir1", "dir2", "dir3", "test.txt")))
}

func (this *SuiteWrite) TestFileExist() {
	assert.True(this.T(), FileExist(this.fileExist))
	assert.False(this.T(), FileExist(this.pathExist))
	assert.False(this.T(), FileExist(testdata.Unknown))
}

func (this *SuiteWrite) TestWriteFile() {
	path := filepath.Join("path", "test.file")
	data := []byte("this is a string")

	assert.Nil(this.T(), WriteFile(path, data))
	assert.True(this.T(), testdata.CompareFile(path, data))
}

func (this *SuiteWrite) TestWriteTmpl() {
	path := filepath.Join("write", "write.tmpl")
	contentReal := "{{$.Value}}"
	contentFake := "{{{$.Value}}"

	assert.Nil(this.T(), WriteTmpl(path, contentReal, map[string]string{"Value": "Value"}))
	assert.True(this.T(), testdata.CompareFile(path, []byte("Value")))

	assert.NotNil(this.T(), WriteTmpl(path, contentFake, nil))
	assert.NotNil(this.T(), WriteTmpl(path, contentReal, "nothing!"))
}
