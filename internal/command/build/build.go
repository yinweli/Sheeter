package build

import (
	"Sheeter/internal/command/build/builder"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	buildConfig, errs := builder.ReadConfig(args[0])

	if errs != nil {
		for _, itor := range errs {
			cmd.Println(itor)
		} // for

		return
	} // if

	if len(args) != 1 { // 這裡用來阻擋單元測試時的虛假設定檔測試
		return
	} // if

	for _, itor := range buildConfig.Elements {
		cargo := builder.Cargo{
			Output:  cmd.OutOrStdout(),
			Global:  &buildConfig.Global,
			Element: &itor,
		}
		err := builder.ReadExcel(&cargo)

		if err != nil {
			cmd.Println(err)
			continue
		} // if
	} // for
}
