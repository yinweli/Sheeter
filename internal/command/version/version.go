package version

import (
	"Sheeter/internal"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令
func NewCommand() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "version",
		Short: "show version",
		Long:  "show version",
		Run:   execute,
	}
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	cmd.Printf("%s %s", internal.Title, internal.Version)
}
