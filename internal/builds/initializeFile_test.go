package builds

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeFile(t *testing.T) {
	suite.Run(t, new(SuiteInitializeFile))
}

type SuiteInitializeFile struct {
	suite.Suite
	workDir string
	path    string
	file1   string
	file2   string
}

func (this *SuiteInitializeFile) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.path = "file"
	this.file1 = filepath.Join(this.path, "file1"+internal.ExcelExt)
	this.file2 = filepath.Join(this.path, "file2"+internal.ExcelExt)
	_ = os.MkdirAll(this.path, os.ModePerm)
	_ = os.WriteFile(this.file1, []byte{}, fs.ModePerm)
	_ = os.WriteFile(this.file2, []byte{}, fs.ModePerm)
}

func (this *SuiteInitializeFile) TearDownSuite() {
	_ = os.RemoveAll(this.path)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeFile) TestInitializeFile() {
	result := make(chan any, 100)
	assert.Nil(this.T(), InitializeFile(this.path, result))
	assert.Len(this.T(), result, 2)
	assert.Equal(this.T(), this.file1, <-result)
	assert.Equal(this.T(), this.file2, <-result)
}
