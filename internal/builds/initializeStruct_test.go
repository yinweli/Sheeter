package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeStruct(t *testing.T) {
	suite.Run(t, new(SuiteInitializeStruct))
}

type SuiteInitializeStruct struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeStruct) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeStruct) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeStruct) target() *Context {
	target := &Context{
		Config: &Config{
			Global: Global{
				LineOfName:  1,
				LineOfNote:  2,
				LineOfField: 3,
				LineOfLayer: 4,
			},
		},
		Sector: []*ContextSector{
			{
				Element: Element{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetName,
				},
			},
			{
				Element: Element{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetName,
				},
			},
		},
	}
	return target
}

func (this *SuiteInitializeStruct) TestInitializeStruct() {
	target := this.target()
	assert.Nil(this.T(), initializeSector(target, target.Sector[0]))
	assert.Nil(this.T(), initializeSector(target, target.Sector[1]))
	assert.Nil(this.T(), initializeStruct(target))

	structNames := []string{}

	for _, itor := range target.Struct {
		structName := mixeds.NewMixed(itor.types.Excel, itor.types.Sheet).StructName()
		structNames = append(structNames, structName)
	} // for

	assert.ElementsMatch(this.T(), []string{"RealData", "S", "A"}, structNames)
}
