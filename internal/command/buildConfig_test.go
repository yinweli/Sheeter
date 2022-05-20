package command

import "testing"

func TestBuildConfig_Check(t *testing.T) {
	buildConfig := BuildConfig{}

	recoveryBuildConfig(&buildConfig)

	if buildConfig.Check() == false {
		t.Error("check == false")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.ExcelPath = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Global.ExcelPath)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathJson = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Global.OutputPathJson)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathGo = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Global.OutputPathGo)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathCs = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Global.OutputPathCs)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.OutputPathCpp = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Global.OutputPathCpp)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.LineOfNote = 3

	if buildConfig.Check() == true {
		t.Error("check == true (Global.LineOfNote)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Global.LineOfField = 3

	if buildConfig.Check() == true {
		t.Error("check == true (Global.LineOfField)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements = []Element{}

	if buildConfig.Check() == true {
		t.Error("check == true (Elements empty)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements[0].ExcelName = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Elements.ExcelName)")
	}

	recoveryBuildConfig(&buildConfig)
	buildConfig.Elements[0].SheetName = ""

	if buildConfig.Check() == true {
		t.Error("check == true (Elements.SheetName)")
	}
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
