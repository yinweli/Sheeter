package tasks

import (
	"os"
	"testing"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	target := &Task{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
		columns: []*Column{
			{Name: "name0", Note: "note0", Field: &fields.Pkey{}},
			{Name: "name1", Note: "note1", Field: &fields.Bool{}},
			{Name: "name2", Note: "note2", Field: &fields.Int{}},
			{Name: "name3", Note: "note3", Field: &fields.Text{}},
		},
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
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonGo())
	target.close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonGo())
	target.close()
}
