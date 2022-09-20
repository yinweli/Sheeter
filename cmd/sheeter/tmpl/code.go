package tmpl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/tmpls"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tmpl",
		Short: "tmpl initialize",
		Long:  "tmpl initialize, user can edit tmpl file to change generation result",
		Run:   execute,
	}
	tmpls.SetFlags(cmd)
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	if err := tmpls.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("tmpl initialize failed: %w", err))
		return
	} // if

	cmd.Println("tmpl initialize complete")
}
