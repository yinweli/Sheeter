package build

import (
	"Sheeter/internal/command/build/core"
	"Sheeter/internal/util"

	"github.com/spf13/cobra"
)

const flagJson string = "json" // 命令旗標: 輸出json
const flagCpp string = "cpp"   // 命令旗標: 輸出c++
const flagCs string = "cs"     // 命令旗標: 輸出c#
const flagGo string = "go"     // 命令旗標: 輸出go

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		RunE:  execute,
	}
	cmd.Flags().BoolP(flagJson, "j", false, "generate json file")
	cmd.Flags().BoolP(flagCpp, "c", false, "generate cpp file")
	cmd.Flags().BoolP(flagCs, "s", false, "generate cs file")
	cmd.Flags().BoolP(flagGo, "g", false, "generate go file")

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) error {
	// TODO: 重構成Executor
	// TODO: 做WriteGo
	// TODO: WriteCpp, WriteCs, WriteGo, WriteJson的單元測試

	config, err := core.ReadConfig(args[0])

	if err != nil {
		return err
	} // if

	toJson := flag(cmd, flagJson)
	toCpp := flag(cmd, flagCpp)
	toCs := flag(cmd, flagCs)

	cmd.Printf("toJson: %t\n", toJson)
	cmd.Printf("toCpp: %t\n", toCpp)
	cmd.Printf("toCs: %t\n", toCs)

	for _, itor := range config.Elements {
		progress := util.NewProgressBar(itor.GetFullName(), cmd.OutOrStdout())
		cargo := &core.Cargo{
			Progress: progress,
			Global:   &config.Global,
			Element:  &itor,
		}
		err = core.ReadSheet(cargo, task(cmd))

		if err != nil {
			return err
		} // if

		if toJson {
			_, err := core.WriteJson(cargo)

			if err != nil {
				return err
			} // if
		} // if

		if toCpp {
			_, err := core.WriteCpp(cargo)

			if err != nil {
				return err
			} // if
		} // if

		if toCs {
			_, err := core.WriteCs(cargo)

			if err != nil {
				return err
			} // if
		} // if

		_ = progress.Finish()
	} // for

	return nil
}

// task 計算工作數量
func task(cmd *cobra.Command) int {
	var task int

	if flag(cmd, flagJson) {
		task++
	} // if

	if flag(cmd, flagCpp) {
		task++
	} // if

	if flag(cmd, flagCs) {
		task++
	} // if

	if flag(cmd, flagGo) {
		task++
	} // if

	return task + 1 // + 1 是包括讀取設定與表格這項工作
}

// flag 取得命令旗標
func flag(cmd *cobra.Command, name string) bool {
	result, err := cmd.Flags().GetBool(name)

	if err != nil {
		return false
	} // if

	return result
}
