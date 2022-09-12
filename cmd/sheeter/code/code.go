package code

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/builds"
)

const flagClean = "clean" // 旗標名稱: 清理模板

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "code initialize",
		Long:  "save code template file, user can edit this file to change generation code",
		Run:   execute,
	}
	cmd.Flags().BoolP(flagClean, "c", false, "clean code template file")
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	code := builds.Code{}
	flags := cmd.Flags()

	if flags.Changed(flagClean) {
		if clean, err := flags.GetBool(flagClean); err == nil && clean {
			code.Clean()
		} // if
	} // if

	if err := code.Initialize(); err != nil {
		cmd.Println(fmt.Errorf("code initialize failed: %w", err))
		return
	} // if

	cmd.Println("code initialize complete")
}
