package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestWrite(t *testing.T) {
	suite.Run(t, new(SuiteWrite))
}

type SuiteWrite struct {
	suite.Suite
	workDir string
	dirReal string
	dirFake string
}

func (this *SuiteWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dirReal = "write"
	this.dirFake = "?write"
}

func (this *SuiteWrite) TearDownSuite() {
	_ = os.RemoveAll(this.dirReal)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteWrite) TestFileName() {
	name := "test"
	path := filepath.Join("dir1", "dir2", "dir3", name+".txt")

	assert.Equal(this.T(), name, FileName(path))
}

func (this *SuiteWrite) TestExistFile() {
	assert.True(this.T(), ExistFile(testdata.ConfigNameReal))
	assert.False(this.T(), ExistFile(testdata.UnknownStr))
}

func (this *SuiteWrite) TestWriteFile() {
	name := "test.file"
	pathReal := filepath.Join(this.dirReal, name)
	pathFake := filepath.Join(this.dirFake, name)
	bytes := []byte("this is a string")

	assert.Nil(this.T(), WriteFile(pathReal, bytes))
	testdata.CompareFile(this.T(), pathReal, bytes)

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), WriteFile(pathFake, bytes))
	} // if
}

func (this *SuiteWrite) TestWriteTmpl() {
	name := "test.tmpl"
	pathReal := filepath.Join(this.dirReal, name)
	pathFake := filepath.Join(this.dirFake, name)
	contentReal := "{{$.Value}}"
	contentFake := "{{{$.Value}}"
	datas := map[string]string{"Value": "Value"}
	bytes := []byte("Value")

	assert.Nil(this.T(), WriteTmpl(pathReal, contentReal, datas))
	testdata.CompareFile(this.T(), pathReal, bytes)

	assert.NotNil(this.T(), WriteTmpl(pathReal, contentFake, nil))
	assert.NotNil(this.T(), WriteTmpl(pathReal, contentReal, "nothing!"))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), WriteTmpl(pathFake, contentReal, datas))
	} // if
}
