package tasks

import (
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskUtil(t *testing.T) {
	suite.Run(t, new(SuiteTaskUtil))
}

type SuiteTaskUtil struct {
	suite.Suite
}

func (this *SuiteTaskUtil) target() *Task {
	target := &Task{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteTaskUtil) TestTargetName() {
	assert.Equal(this.T(), "real.xlsx(Data)", this.target().targetName())
}

func (this *SuiteTaskUtil) TestNamespace() {
	assert.Equal(this.T(), "sheeter", this.target().namespace())
}

func (this *SuiteTaskUtil) TestStructName() {
	assert.Equal(this.T(), "RealData", this.target().structName())
}

func (this *SuiteTaskUtil) TestReaderName() {
	assert.Equal(this.T(), "RealDataReader", this.target().readerName())
}

func (this *SuiteTaskUtil) TestExcelName() {
	assert.Equal(this.T(), "real", this.target().excelName())
}

func (this *SuiteTaskUtil) TestSheetExists() {
	target := this.target()

	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.True(this.T(), target.sheetExists())

	target.Sheet = testdata.UnknownStr
	assert.False(this.T(), target.sheetExists())

	target.close()
}

func (this *SuiteTaskUtil) TestGetRows() {
	target := this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)

	rows, err := target.getRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = target.getRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = target.getRows(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.getRows(1)
	assert.NotNil(this.T(), err)

	target.close()
}

func (this *SuiteTaskUtil) TestGetRowContent() {
	target := this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)

	cols, err := target.getRowContent(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"name0#pkey", "name1#bool", "name2#int", "name3#text", "empty#empty"}, cols)

	cols, err = target.getRowContent(3)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"note0", "note1", "note2", "note3", "empty"}, cols)

	_, err = target.getRowContent(10)
	assert.NotNil(this.T(), err)

	_, err = target.getRowContent(0)
	assert.NotNil(this.T(), err)

	target.Sheet = testdata.UnknownStr
	_, err = target.getRowContent(1)
	assert.NotNil(this.T(), err)

	target.close()
}
