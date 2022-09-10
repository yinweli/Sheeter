package builds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/util"
	"github.com/yinweli/Sheeter/testdata"
)

func TestCode(t *testing.T) {
	suite.Run(t, new(SuiteCode))
}

type SuiteCode struct {
	suite.Suite
	workDir      string
	jsonCsReader string
	jsonGoReader string
	fileName     string
	fileData     string
	filePreset   string
}

func (this *SuiteCode) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.jsonCsReader = filepath.Join(internal.PathCode, internal.FileCodeJsonCsReader)
	this.jsonGoReader = filepath.Join(internal.PathCode, internal.FileCodeJsonGoReader)
	this.fileName = "test.txt"
	this.fileData = "test"
	this.filePreset = "preset"
}

func (this *SuiteCode) TearDownSuite() {
	_ = os.RemoveAll(internal.PathCode)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteCode) target() *Code {
	return &Code{}
}

func (this *SuiteCode) TestInitialize() {
	target := this.target()
	assert.Nil(this.T(), target.Initialize())
	assert.Equal(this.T(), jsonCsReader, target.JsonCsReader)
	assert.Equal(this.T(), jsonGoReader, target.JsonGoReader)
	testdata.CompareFile(this.T(), this.jsonCsReader, []byte(jsonCsReader))
	testdata.CompareFile(this.T(), this.jsonGoReader, []byte(jsonGoReader))
}

func (this *SuiteCode) TestLoad() {
	target := this.target()

	tmpl, err := target.load(this.fileName, this.filePreset)
	assert.Equal(this.T(), this.filePreset, tmpl)
	assert.Nil(this.T(), err)

	assert.Nil(this.T(), util.WriteFile(filepath.Join(internal.PathCode, this.fileName), []byte(this.fileData), false))
	tmpl, err = target.load(this.fileName, this.filePreset)
	assert.Equal(this.T(), this.fileData, tmpl)
	assert.Nil(this.T(), err)
}
