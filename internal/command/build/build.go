package build

import (
	"Sheeter/internal"
	"Sheeter/internal/command/build/core"
	"Sheeter/internal/util"

	"github.com/spf13/cobra"
)

// Build 命令物件
var Build = &cobra.Command{
	Use:   "build [config]",
	Short: "build sheet",
	Long:  "build all the sheet in the configuration",
	Args:  cobra.ExactArgs(1),
	Run:   execute,
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	config, err := core.ReadConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if

	cmd.Printf("excelPath: %s\n", config.Global.ExcelPath)
	cmd.Printf("cppLibraryPath: %s\n", config.Global.CppLibraryPath)
	cmd.Printf("csNamespace: %s\n", config.Global.CsNamespace)
	cmd.Printf("goPackage: %s\n", config.Global.GoPackage)
	cmd.Printf("bom: %t\n", config.Global.Bom)
	cmd.Printf("lineOfNote: %d\n", config.Global.LineOfNote)
	cmd.Printf("lineOfField: %d\n", config.Global.LineOfField)
	cmd.Printf("lineOfData: %d\n", config.Global.LineOfData)

	for _, itor := range config.Elements {
		progress := util.NewProgressBar(internal.ProgressDefault, itor.GetFullName(), cmd.OutOrStdout())
		cargo := &core.Cargo{
			Progress: progress,
			Global:   &config.Global,
			Element:  &itor,
		}
		err := core.ReadSheet(cargo)

		if err != nil {
			cmd.Println(err)
			return
		} // if

		_ = progress.Finish()
	} // for
}
