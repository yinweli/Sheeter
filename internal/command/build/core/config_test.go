package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := mockConfig()
	err := config.Check()
	assert.Nil(t, err)

	config = mockConfig()
	config.Global.ExcelPath = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.CppNamespace = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.CsNamespace = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.GoPackage = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfData = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 3
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 3
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Elements = []Element{}
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Elements[0].Excel = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Elements[0].Sheet = ""
	err = config.Check()
	assert.NotNil(t, err)
}

func TestGlobal(t *testing.T) {
	config := mockConfig()
	assert.Equal(t, config.Global.LineOfNote-1, config.Global.GetLineOfNote())
	assert.Equal(t, config.Global.LineOfField-1, config.Global.GetLineOfField())
	assert.Equal(t, config.Global.LineOfData-1, config.Global.GetLineOfData())
}

func TestElement(t *testing.T) {
	config := mockConfig()
	element := config.Elements[0]
	assert.Equal(t, fmt.Sprintf("%s(%s)", element.Excel, element.Sheet), element.GetFullName())
}

func mockConfig() *Config {
	return &Config{
		Global: Global{
			ExcelPath:      "test",
			CppLibraryPath: "nlohmann",
			CppNamespace:   "Sheeter",
			CsNamespace:    "Sheeter",
			GoPackage:      "sheeter",
			Bom:            true,
			LineOfNote:     1,
			LineOfField:    2,
			LineOfData:     3,
		},
		Elements: []Element{{
			Excel: "Test.xlsx",
			Sheet: "Data",
		}},
	}
}
