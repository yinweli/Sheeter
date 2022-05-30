package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeGenerate(t *testing.T) {
	code := "{{excelName .Element.Excel}}#{{sheetName .Element.Sheet}}#{{memberName .Global.GoPackage}}"
	cargo := &Cargo{
		Global: &Global{
			GoPackage: "testPackage",
		},
		Element: &Element{
			Excel: "testExcel",
			Sheet: "testSheet",
		},
	}

	result, err := CodeGenerate(code, cargo)
	assert.Nil(t, err)
	assert.Equal(t, "TestExcel#TestSheet#TestPackage", result)

	result, err = CodeGenerate("{{{.Unknown}}", cargo)
	assert.NotNil(t, err)

	result, err = CodeGenerate(code, nil)
	assert.NotNil(t, err)
}
