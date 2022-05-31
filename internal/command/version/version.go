package version

import (
	"Sheeter/internal"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show version",
		Long:  "show version",
		RunE:  execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) error {
	cmd.Printf("%s %s", internal.Title, internal.Version)

	return nil
}
