package tasks

import (
	"testing"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskColumn(t *testing.T) {
	suite.Run(t, new(SuiteTaskColumn))
}

type SuiteTaskColumn struct {
	suite.Suite
	columns []*Column
}

func (this *SuiteTaskColumn) SetupSuite() {
	this.columns = []*Column{
		{Name: "name0", Note: "note0", Field: &fields.Pkey{}},
		{Name: "name1", Note: "note1", Field: &fields.Bool{}},
		{Name: "name2", Note: "note2", Field: &fields.Int{}},
		{Name: "name3", Note: "note3", Field: &fields.Text{}},
		{Name: "empty", Note: "empty", Field: &fields.Empty{}},
	}
}

func (this *SuiteTaskColumn) target() *Task {
	target := &Task{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.RealExcel,
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteTaskColumn) TestTaskColumn() {
	target := this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.Nil(this.T(), target.runColumn())
	assert.Equal(this.T(), this.columns, target.columns)
	target.close()

	target = this.target()
	target.LineOfField = 10
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.LineOfNote = 10
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect2Excel)
	assert.Nil(this.T(), target.runColumn()) // 測試其實會成功
	assert.Equal(this.T(), 4, len(target.columns))
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect3Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect4Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect5Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect6Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect7Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect8Excel)
	assert.NotNil(this.T(), target.runColumn())
	target.close()
}

func TestFromList(t *testing.T) {
	suite.Run(t, new(SuiteFromList))
}

type SuiteFromList struct {
	suite.Suite
	lists []string
}

func (this *SuiteFromList) SetupSuite() {
	this.lists = []string{"a", "b", "c"}
}

func (this *SuiteFromList) TestFromList() {
	assert.Equal(this.T(), "a", fromList(this.lists, 0))
	assert.Equal(this.T(), "b", fromList(this.lists, 1))
	assert.Equal(this.T(), "c", fromList(this.lists, 2))
	assert.Equal(this.T(), "", fromList(this.lists, 3))
}
