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
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("%s version %s\n", internal.Title, internal.Version)
		},
	}
	return cmd
}
