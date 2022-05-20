package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConfig_Check(t *testing.T) {
	buildConfig := BuildConfig{}

	recoveryBuildConfig(&buildConfig)
	assert.Equal(t, true, buildConfig.Check(), "check failed")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.ExcelPath = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.ExcelPath)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathJson = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathJson)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathGo = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathGo)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathCs = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathCs)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathCpp = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathCpp)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.LineOfNote = 3
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.LineOfNote)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.LineOfField = 3
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.LineOfField)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements = []Element{}
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements empty)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements[0].ExcelName = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements.ExcelName)")

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements[0].SheetName = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements.SheetName)")

}

// recoveryBuildConfig 復原建立表格設定資料
func recoveryBuildConfig(buildConfig *BuildConfig) {
	buildConfig.Global.ExcelPath = "testme"
	buildConfig.Global.OutputPathJson = "testme"
	buildConfig.Global.OutputPathGo = "testme"
	buildConfig.Global.OutputPathCs = "testme"
	buildConfig.Global.OutputPathCpp = "testme"
	buildConfig.Global.CppLibraryPath = "testme"
	buildConfig.Global.LineOfNote = 0
	buildConfig.Global.LineOfField = 1
	buildConfig.Global.LineOfData = 2
	buildConfig.Elements = []Element{
		Element{ExcelName: "testme", SheetName: "testme"},
	}
}
