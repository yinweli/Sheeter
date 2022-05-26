package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := mockConfig()
	err := config.Check()
	assert.Nil(t, err, "check failed")

	config = mockConfig()
	config.Global.ExcelPath = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.ExcelPath)")

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.CppLibraryPath)")

	config = mockConfig()
	config.Global.CsNamespace = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.CsNamespace)")

	config = mockConfig()
	config.Global.GoPackage = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.GoPackage)")

	config = mockConfig()
	config.Global.LineOfNote = 0
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfNote <= 0)")

	config = mockConfig()
	config.Global.LineOfField = 0
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfField <= 0)")

	config = mockConfig()
	config.Global.LineOfData = 0
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfData <= 0)")

	config = mockConfig()
	config.Global.LineOfNote = 3
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfNote)")

	config = mockConfig()
	config.Global.LineOfField = 3
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfField)")

	config = mockConfig()
	config.Elements = []Element{}
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements empty)")

	config = mockConfig()
	config.Elements[0].Excel = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements.Excel)")

	config = mockConfig()
	config.Elements[0].Sheet = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements.Sheet)")
}

func TestGlobal(t *testing.T) {
	config := mockConfig()
	assert.Equal(t, config.Global.LineOfNote-1, config.Global.GetLineOfNote(), "get line of note failed")
	assert.Equal(t, config.Global.LineOfField-1, config.Global.GetLineOfField(), "get line of field failed")
	assert.Equal(t, config.Global.LineOfData-1, config.Global.GetLineOfData(), "get line of data failed")
}

func mockConfig() *Config {
	return &Config{
		Global: Global{
			ExcelPath:      "test",
			CppLibraryPath: "nlohmann",
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
