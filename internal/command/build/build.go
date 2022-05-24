package build

import (
	"io"
	"io/ioutil"

	"Sheeter/internal/command/build/config"
	"Sheeter/internal/command/build/field"

	"github.com/schollz/progressbar/v3"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewCommand 建立命令
func NewCommand() *cobra.Command {
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
	buildConfig := config.Config{}
	err := readConfig(args[0], &buildConfig)

	if err != nil {
		cmd.Println(err)
		return
	} // if
}

// readConfig 讀取編譯設置
func readConfig(filename string, buildConfig *config.Config) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	} // if

	err = yaml.Unmarshal(file, &buildConfig)

	if err != nil {
		return err
	} // if

	return nil
}

// meta 元資料
type meta struct {
	output   io.Writer                // 輸出物件
	global   *config.Global           // 全域設定
	element  *config.Element          // 項目設定
	progress *progressbar.ProgressBar // 進度條
	columns  []column                 // 欄位列表
}

// column 欄位資料
type column struct {
	note  string       // 欄位註解
	name  string       // 欄位名稱
	field *field.Field // 欄位類型
	datas []string     // 資料列表
}
