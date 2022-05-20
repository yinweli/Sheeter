package config

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig_Check(t *testing.T) {
	var config Config
	var errors []string
	var result bool

	fakeConfig(t, &config)
	errors, result = config.Check()
	assert.Equal(t, 0, len(errors), "check errors failed")
	assert.Equal(t, true, result, "check result failed")

	fakeConfig(t, &config)
	config.Global.ExcelPath = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.ExcelPath)")
	assert.Equal(t, false, result, "check result failed(Global.ExcelPath)")

	fakeConfig(t, &config)
	config.Global.OutputPathJson = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathJson)")
	assert.Equal(t, false, result, "check result failed(Global.OutputPathJson)")

	fakeConfig(t, &config)
	config.Global.OutputPathCpp = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathCpp)")
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCpp)")

	fakeConfig(t, &config)
	config.Global.OutputPathCs = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathCs)")
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCs)")

	fakeConfig(t, &config)
	config.Global.OutputPathGo = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathGo)")
	assert.Equal(t, false, result, "check result failed(Global.OutputPathGo)")

	fakeConfig(t, &config)
	config.Global.GoPackage = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.GoPackage)")
	assert.Equal(t, false, result, "check result failed(Global.GoPackage)")

	fakeConfig(t, &config)
	config.Global.LineOfNote = 3
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.LineOfNote)")
	assert.Equal(t, false, result, "check result failed(Global.LineOfNote)")

	fakeConfig(t, &config)
	config.Global.LineOfField = 3
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Global.LineOfField)")
	assert.Equal(t, false, result, "check result failed(Global.LineOfField)")

	fakeConfig(t, &config)
	config.Elements = []Element{}
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Elements empty)")
	assert.Equal(t, false, result, "check result failed(Elements empty)")

	fakeConfig(t, &config)
	config.Elements[0].ExcelName = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Elements.ExcelName)")
	assert.Equal(t, false, result, "check result failed(Elements.ExcelName)")

	fakeConfig(t, &config)
	config.Elements[0].SheetName = ""
	errors, result = config.Check()
	assert.Equal(t, 1, len(errors), "check errors failed(Elements.SheetName)")
	assert.Equal(t, false, result, "check result failed(Elements.SheetName)")
}

// fakeConfig 取得假的編譯設定
func fakeConfig(t *testing.T, config *Config) {
	filename := testdata.Path("config.yaml")
	file, err := ioutil.ReadFile(filename)

	assert.NotNil(t, file, "load config failed")
	assert.Nil(t, err, err)

	err = yaml.Unmarshal(file, config)

	assert.Nil(t, err, err)
}
