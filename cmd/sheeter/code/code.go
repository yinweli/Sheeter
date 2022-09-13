package code

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/codes"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "code initialize",
		Long:  "code initialize, user can edit code file to change generation result",
		Run:   execute,
	}
	codes.SetFlags(cmd)
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	if err := codes.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("code initialize failed: %w", err))
		return
	} // if

	cmd.Println("code initialize complete")
}
