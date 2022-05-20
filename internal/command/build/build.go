package build

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令
func NewCommand() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
