package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := mockConfig()
	err := config.Check()
	assert.Nil(t, err)

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfData = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 3
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 3
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

func mockConfig() *Config {
	return &Config{
		Global: Global{
			ExcelPath:      "excel",
			CppLibraryPath: "cpp",
			Bom:            true,
			LineOfField:    1,
			LineOfNote:     2,
			LineOfData:     3,
		},
		Elements: []Element{{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		}},
	}
}
