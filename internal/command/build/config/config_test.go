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
	var result bool
	var errors []string

	fakeConfig(t, &config)
	result, errors = config.Check()
	assert.Equal(t, true, result, "check result failed")
	assert.Equal(t, 0, len(errors), "check errors failed")

	fakeConfig(t, &config)
	config.Global.ExcelPath = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.ExcelPath)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.ExcelPath)")

	fakeConfig(t, &config)
	config.Global.OutputPathJson = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathJson)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathJson)")

	fakeConfig(t, &config)
	config.Global.OutputPathCpp = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCpp)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathCpp)")

	fakeConfig(t, &config)
	config.Global.OutputPathCs = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCs)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathCs)")

	fakeConfig(t, &config)
	config.Global.OutputPathGo = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathGo)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.OutputPathGo)")

	fakeConfig(t, &config)
	config.Global.GoPackage = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.GoPackage)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.GoPackage)")

	fakeConfig(t, &config)
	config.Global.LineOfNote = 3
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfNote)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.LineOfNote)")

	fakeConfig(t, &config)
	config.Global.LineOfField = 3
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfField)")
	assert.Equal(t, 1, len(errors), "check errors failed(Global.LineOfField)")

	fakeConfig(t, &config)
	config.Elements = []Element{}
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements empty)")
	assert.Equal(t, 1, len(errors), "check errors failed(Elements empty)")

	fakeConfig(t, &config)
	config.Elements[0].ExcelName = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.ExcelName)")
	assert.Equal(t, 1, len(errors), "check errors failed(Elements.ExcelName)")

	fakeConfig(t, &config)
	config.Elements[0].SheetName = ""
	result, errors = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.SheetName)")
	assert.Equal(t, 1, len(errors), "check errors failed(Elements.SheetName)")

}

// fakeConfig 取得假的編譯設定
func fakeConfig(t *testing.T, config *Config) {
	path := testdata.Path("config.yaml")
	file, err := ioutil.ReadFile(path)

	assert.NotNil(t, file, "load config failed")
	assert.Nil(t, err, err)

	err = yaml.Unmarshal(file, config)

	assert.Nil(t, err, err)
}
