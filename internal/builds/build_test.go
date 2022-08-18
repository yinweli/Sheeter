package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuild(t *testing.T) {
	suite.Run(t, new(SuiteBuild))
}

type SuiteBuild struct {
	suite.Suite
	workDir string
}

func (this *SuiteBuild) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteBuild) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJson)
	_ = os.RemoveAll(pathJsonCs)
	_ = os.RemoveAll(pathJsonGo)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBuild) target() *Content {
	target := &Content{
		Path:        testdata.RootPath,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteBuild) TestBuild() {
	target := this.target()
	assert.Nil(this.T(), Build(target))
	target.close()
}
