package tmpls

import (
	"github.com/spf13/cobra"
)

const flagClean = "clean" // 旗標名稱: 清理模板

// SetFlags 設定命令旗標
func SetFlags(cmd *cobra.Command) *cobra.Command {
	flags := cmd.Flags()
	flags.BoolP(flagClean, "c", false, "clean template file")
	return cmd
}
