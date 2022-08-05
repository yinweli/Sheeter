package build

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonGo(t *testing.T) {
	suite.Run(t, new(SuiteTaskJsonGo))
}

type SuiteTaskJsonGo struct {
	suite.Suite
	workDir   string
	dataBytes []byte
}

func (this *SuiteTaskJsonGo) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`package sheeter

type RealData struct {
	Name0 int64  ` + "`json:\"name0\"`" + `
	Name1 bool   ` + "`json:\"name1\"`" + `
	Name2 int64  ` + "`json:\"name2\"`" + `
	Name3 string ` + "`json:\"name3\"`" + `
}
`)
}

func (this *SuiteTaskJsonGo) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJsonGo)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskJsonGo) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	target.columns = []*Column{
		{Name: "name0", Note: "note0", Field: &FieldPkey{}},
		{Name: "name1", Note: "note1", Field: &FieldBool{}},
		{Name: "name2", Note: "note2", Field: &FieldInt{}},
		{Name: "name3", Note: "note3", Field: &FieldText{}},
	}
	return target
}

func (this *SuiteTaskJsonGo) TestTaskJsonGo() {
	target := this.target()
	assert.Nil(this.T(), target.runJsonSchema())
	assert.Nil(this.T(), target.runJsonGo())
	testdata.CompareFile(this.T(), target.jsonGoFilePath(), this.dataBytes)
	target.close()

	target = this.target()
	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonGo())
	target.close()

	target = this.target()
	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonGo())
	target.close()
}
