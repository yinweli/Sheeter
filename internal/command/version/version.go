package version

import (
	"github.com/yinweli/Sheeter/internal"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version for sheeter",
		Long:  "version for sheeter",
		Run:   execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	cmd.Printf("%s version %s\n", internal.Title, internal.Version)
}
