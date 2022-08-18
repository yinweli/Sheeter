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

func (this *SuiteBuildLayout) content() *Content {
	content := &Content{
		Path:        testdata.RootPath,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return content
}

func (this *SuiteBuildLayout) TestBuildLayout() {
	content := this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), content.builder)
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	content.LineOfField = 10
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	content.LineOfLayer = 10
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	content.LineOfNote = 10
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameCleanAll)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameCleanField)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidField)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidLayer)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidLayout)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidPkeyZero)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidPkeyDupl)
	assert.NotNil(this.T(), buildLayout(content))
	content.close()
}
