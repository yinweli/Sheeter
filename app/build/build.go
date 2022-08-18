package build

import (
	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build the sheet in the params",
		Run: func(cmd *cobra.Command, args []string) {
			duration, errs := execute(args[0])

			for _, itor := range errs {
				cmd.Println(itor)
			} // for

			cmd.Printf("usage time=%s\n", durafmt.Parse(duration))
		},
	}
	// p *string, name, shorthand string, value string, usage string
	cmd.Flags().StringP()
	return cmd
}

// Params 編譯參數
type Params struct {
	Path        string // 來源excel路徑
	Bom         bool   // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int    // 欄位行號(1為起始行)
	LineOfLayer int    // 階層行號(1為起始行)
	LineOfNote  int    // 註解行號(1為起始行)
	LineOfData  int    // 資料起始行號(1為起始行)
	Excel       string // excel檔案名稱
	Sheet       string // excel表單名稱
}

// TODO: 要記得做build命令
