package build

import (
	"Sheeter/internal/command/build/core"

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
	_, err := core.ReadConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if
}
