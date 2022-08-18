package version

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal"
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
func execute(cmd *cobra.Command, _ []string) {
	cmd.Printf("%s version %s\n", internal.Title, internal.Version)
}
