package codes

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestCode(t *testing.T) {
	suite.Run(t, new(SuiteCode))
}

type SuiteCode struct {
	suite.Suite
	workDir string
	name    string
	code1   string
	code2   string
}

func (this *SuiteCode) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.name = "code.txt"
	this.code1 = "code1"
	this.code2 = "code2"
}

func (this *SuiteCode) TearDownSuite() {
	_ = os.RemoveAll(internal.PathCode)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteCode) target() *Code {
	target := &Code{
		Name: this.name,
		Code: this.code1,
	}
	return target
}

func (this *SuiteCode) TestInitialize() {
	cmd := SetFlags(&cobra.Command{})
	assert.Nil(this.T(), utils.WriteFile(JsonCsReader.path(), []byte(this.code1)))
	assert.Nil(this.T(), utils.WriteFile(JsonGoReader.path(), []byte(this.code2)))
	assert.Nil(this.T(), Initialize(cmd))
	assert.Equal(this.T(), this.code1, JsonCsReader.Code)
	assert.Equal(this.T(), this.code2, JsonGoReader.Code)
	testdata.CompareFile(this.T(), JsonCsReader.path(), []byte(JsonCsReader.Code))
	testdata.CompareFile(this.T(), JsonGoReader.path(), []byte(JsonGoReader.Code))

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagClean, strconv.FormatBool(true)))
	assert.Nil(this.T(), Initialize(cmd))
	testdata.CompareFile(this.T(), JsonCsReader.path(), []byte(JsonCsReader.Code))
	testdata.CompareFile(this.T(), JsonGoReader.path(), []byte(JsonGoReader.Code))
}

func (this *SuiteCode) TestLoad() {
	target := this.target()
	assert.Nil(this.T(), target.load())
	assert.Equal(this.T(), this.code1, target.Code)

	target = this.target()
	assert.Nil(this.T(), utils.WriteFile(target.path(), []byte(this.code2)))
	assert.Nil(this.T(), target.load())
	assert.Equal(this.T(), this.code2, target.Code)
}

func (this *SuiteCode) TestSave() {
	target := this.target()
	assert.Nil(this.T(), target.save())
	testdata.CompareFile(this.T(), target.path(), []byte(target.Code))
}

func (this *SuiteCode) TestPath() {
	target := this.target()
	assert.Equal(this.T(), filepath.Join(internal.PathCode, target.Name), target.path())
}
