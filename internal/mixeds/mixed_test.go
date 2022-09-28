package mixeds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestMixed(t *testing.T) {
	suite.Run(t, new(SuiteMixed))
}

type SuiteMixed struct {
	suite.Suite
	workDir string
	excel   string
	sheet   string
}

func (this *SuiteMixed) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelMixed"
	this.sheet = "sheetMixed"
}

func (this *SuiteMixed) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteMixed) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}

func (this *SuiteMixed) TestNewMixed() {
	assert.NotNil(this.T(), NewMixed(this.excel, this.sheet))
}

func (this *SuiteMixed) TestName() {
	structName := utils.FirstUpper(this.excel) + utils.FirstUpper(this.sheet)
	readerName := structName + internal.Reader
	storerName := structName + internal.Storer
	storerMessage := internal.NamespaceProto + "." + storerName

	target := this.target()
	assert.Equal(this.T(), internal.AppName, target.AppName())
	assert.Equal(this.T(), internal.NamespaceJson, target.NamespaceJson())
	assert.Equal(this.T(), internal.NamespaceProto, target.NamespaceProto())
	assert.Equal(this.T(), structName, target.StructName())
	assert.Equal(this.T(), readerName, target.ReaderName())
	assert.Equal(this.T(), storerName, target.StorerName())
	assert.Equal(this.T(), internal.StorerDatas, target.StorerDatas())
	assert.Equal(this.T(), storerMessage, target.StorerMessage())
}

func (this *SuiteMixed) TestCombine() {
	token := "#"

	target := this.target()
	assert.Equal(this.T(), this.excel+this.sheet, target.combine(params{}))
	assert.Equal(this.T(), utils.FirstUpper(this.excel)+this.sheet, target.combine(params{excelUpper: true}))
	assert.Equal(this.T(), this.excel+utils.FirstUpper(this.sheet), target.combine(params{sheetUpper: true}))
	assert.Equal(this.T(), this.excel+this.sheet+token, target.combine(params{last: token}))
	assert.Equal(this.T(), this.excel+this.sheet+"."+token, target.combine(params{ext: token}))
}
