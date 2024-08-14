package version

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
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
	stdColor := utils.NewStdColor(cmd.OutOrStdout(), cmd.ErrOrStderr())
	stdColor.Out("%v version %v", sheeter.AppName, sheeter.Version)
}
