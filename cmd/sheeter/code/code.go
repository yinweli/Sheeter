package code

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/builds"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "code",
		Short: "code initialize",
		Long:  "save reader template file, user can edit this file to change reader code",
		Run:   execute,
	}
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	code := builds.Code{}

	if err := code.Initialize(); err != nil {
		cmd.Println(fmt.Errorf("code initialize failed: %w", err))
		return
	} // if

	cmd.Println("code initialize complete")
}
