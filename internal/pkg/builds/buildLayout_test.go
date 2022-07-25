package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildLayout(t *testing.T) {
	suite.Run(t, new(SuiteBuildLayout))
}

type SuiteBuildLayout struct {
	suite.Suite
}

func (this *SuiteBuildLayout) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.Path(testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteBuildLayout) TestBuildLayout() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.NotNil(this.T(), target.builder)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	target.LineOfField = 10
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	target.LineOfLayer = 10
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	target.LineOfNote = 10
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameCleanAll)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameCleanField)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidField)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidLayer)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidLayout)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidPkeyZero)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidPkeyDupl)
	assert.NotNil(this.T(), buildLayout(target))
	target.close()
}
