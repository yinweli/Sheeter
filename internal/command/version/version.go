package version

import (
	"Sheeter/internal"

	"github.com/spf13/cobra"
)

// Version 命令物件
var Version = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version",
	Run:   execute,
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	cmd.Printf("%s %s", internal.Title, internal.Version)
}
