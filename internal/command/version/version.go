package version

import (
	"github.com/spf13/cobra"
)

const version = "1.0.4" // 版本字串

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show version",
		Long:  "show version",
		Run:   execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	cmd.Printf("%s\n", version)
}
