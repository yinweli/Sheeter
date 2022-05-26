package build

import (
	"fmt"

	"Sheeter/internal"
	"Sheeter/internal/command/build/core"
	"Sheeter/internal/util"

	"github.com/spf13/cobra"
)

// NewCommand 取得命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}
	cmd.Flags().BoolP("all", "a", false, "generate all file")
	cmd.Flags().BoolP("json", "j", false, "generate json file")
	cmd.Flags().BoolP("cpp", "c", false, "generate cpp file")
	cmd.Flags().BoolP("cs", "#", false, "generate cs file")
	cmd.Flags().BoolP("go", "g", false, "generate go file")

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	b, _ := cmd.Flags().GetBool("all") // TODO: 做命令旗標
	fmt.Printf("all:%t\n", b)
	b, _ = cmd.Flags().GetBool("json")
	fmt.Printf("json:%t\n", b)
	b, _ = cmd.Flags().GetBool("cpp")
	fmt.Printf("cpp:%t\n", b)
	b, _ = cmd.Flags().GetBool("cs")
	fmt.Printf("cs:%t\n", b)
	b, _ = cmd.Flags().GetBool("go")
	fmt.Printf("go:%t\n", b)

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

		err := core.ReadSheet(cargo, 5) // TODO: 用命令的flags計算task值

		if err != nil {
			cmd.Println(err)
			return
		} // if

		_ = progress.Finish()
	} // for
}
