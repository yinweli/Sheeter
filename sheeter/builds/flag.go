package builds

import (
	"github.com/spf13/cobra"
)

const flagConfig = "config"           // 旗標名稱: 設定檔案路徑
const flagSource = "source"           // 旗標名稱: 來源列表
const flagTag = "tag"                 // 旗標名稱: 標籤列表
const flagAutoKey = "autokey"         // 旗標名稱: 自動選取索引
const flagLineOfTag = "lineOfTag"     // 旗標名稱: 標籤行號
const flagLineOfName = "lineOfName"   // 旗標名稱: 名稱行號
const flagLineOfNote = "lineOfNote"   // 旗標名稱: 註解行號
const flagLineOfField = "lineOfField" // 旗標名稱: 欄位行號
const flagLineOfData = "lineOfData"   // 旗標名稱: 資料行號

// SetFlag 設定命令旗標
func SetFlag(cmd *cobra.Command) *cobra.Command {
	flag := cmd.Flags()
	flag.String(flagConfig, "", "config file path")
	flag.StringSlice(flagSource, []string{}, "source file/folder")
	flag.String(flagTag, "", "tag list that determines the columns to be output")
	flag.Bool(flagAutoKey, false, "automatically select the second column as an key")
	flag.Int(flagLineOfTag, 0, "line of tag")
	flag.Int(flagLineOfName, 0, "line of name")
	flag.Int(flagLineOfNote, 0, "line of note")
	flag.Int(flagLineOfField, 0, "line of field")
	flag.Int(flagLineOfData, 0, "line of data")
	return cmd
}
