package builds

import (
	"github.com/spf13/cobra"
)

const flagConfig = "config"           // 旗標名稱: 設定檔案路徑
const flagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
const flagLineOfLayer = "lineOfLayer" // 旗標名稱: 階層行號
const flagLineOfNote = "lineOfNote"   // 旗標名稱: 註解行號
const flagLineOfData = "lineOfData"   // 旗標名稱: 資料行號
const flagElements = "elements"       // 旗標名稱: 項目列表

// SetFlags 設定命令旗標
func SetFlags(cmd *cobra.Command) *cobra.Command {
	flags := cmd.Flags()
	flags.String(flagConfig, "", "config file path")
	flags.Int(flagLineOfField, 0, "line of field")
	flags.Int(flagLineOfLayer, 0, "line of layer")
	flags.Int(flagLineOfNote, 0, "line of note")
	flags.Int(flagLineOfData, 0, "line of data")
	flags.StringSlice(flagElements, []string{}, "element lists(excel:sheet,excel:sheet,excel:sheet,...)")
	return cmd
}
