package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEntire(t *testing.T) {
	suite.Run(t, new(SuiteEntire))
}

type SuiteEntire struct {
	suite.Suite
	workDir          string
	structName       string
	readerName       string
	fileJson         string
	fileJsonCsCode   string
	fileJsonCsReader string
	fileJsonGoCode   string
	fileJsonGoReader string
	fileJsonCode     string
}

func (this *SuiteEntire) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.structName = "testStruct"
	this.readerName = "testReader"
	this.fileJson = "test\\Json"
	this.fileJsonCsCode = filepath.Join(internal.PathJsonCs, this.structName+"."+internal.ExtCs)
	this.fileJsonCsReader = filepath.Join(internal.PathJsonCs, this.readerName+"."+internal.ExtCs)
	this.fileJsonGoCode = filepath.Join(internal.PathJsonGo, this.structName+"."+internal.ExtGo)
	this.fileJsonGoReader = filepath.Join(internal.PathJsonGo, this.readerName+"."+internal.ExtGo)
	this.fileJsonCode = filepath.ToSlash(this.fileJson)
}

func (this *SuiteEntire) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEntire) target() *Entire {
	target := &Entire{
		Type: &layouts.Type{
			StructName: this.structName,
			ReaderName: this.readerName,
			FileJson:   this.fileJson,
		},
	}
	return target
}

func (this *SuiteEntire) TestName() {
	target := this.target()
	assert.Equal(this.T(), internal.AppName, target.AppName())
	assert.Equal(this.T(), internal.AppName, target.Namespace())
	assert.Equal(this.T(), this.fileJsonCsCode, target.FileJsonCsCode())
	assert.Equal(this.T(), this.fileJsonCsReader, target.FileJsonCsReader())
	assert.Equal(this.T(), this.fileJsonGoCode, target.FileJsonGoCode())
	assert.Equal(this.T(), this.fileJsonGoReader, target.FileJsonGoReader())
	assert.Equal(this.T(), this.fileJsonCode, target.FileJsonCode())
}

func (this *SuiteEntire) TestCombine() {
	target := this.target()
	assert.Equal(this.T(), filepath.Join("a", "b.c"), target.combine("a", "b", "c"))
}
