package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildExcel(t *testing.T) {
	suite.Run(t, new(SuiteBuildExcel))
}

type SuiteBuildExcel struct {
	suite.Suite
}

func (this *SuiteBuildExcel) content() *Content {
	content := &Content{
		Path:  testdata.RootPath,
		Excel: testdata.ExcelNameReal,
		Sheet: testdata.SheetName,
	}
	return content
}

func (this *SuiteBuildExcel) TestReadExcel() {
	content := this.content()
	assert.Nil(this.T(), readExcel(content))
	assert.NotNil(this.T(), content.excel)
	content.close()

	content = this.content()
	content.Path = ""
	assert.NotNil(this.T(), readExcel(content))
	content.close()

	content = this.content()
	content.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), readExcel(content))
	content.close()

	content = this.content()
	content.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), readExcel(content))
	content.close()

	content = this.content()
	content.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), readExcel(content))
	content.close()
}
