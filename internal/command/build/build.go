package build

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Build 命令物件
var Build = &cobra.Command{
	Use:   "build [config]",
	Short: "build sheet",
	Long:  "build all the sheet in the configuration",
	Args:  cobra.ExactArgs(1),
	Run:   execute,
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	build := build{
		cmd:  cmd,
		args: args,
		stop: false,
	}
	build.readConfig()
}

// build 編譯命令
type build struct {
	cmd      *cobra.Command           // 命令物件
	args     []string                 // 參數列表
	stop     bool                     // 停止旗標
	config   *config.Config           // 設定物件
	progress *progressbar.ProgressBar // 進度條
}

// readConfig 讀取設定
func (this *build) readConfig() {
	bytes, err := ioutil.ReadFile(this.args[0])

	if err != nil {
		this.failed(err.Error())
		return
	} // if

	this.config = &config.Config{}
	err = yaml.Unmarshal(bytes, this.config)

	if err != nil {
		this.failed(err.Error())
		return
	} // if

	err = this.config.Check()

	if err != nil {
		this.failed(err.Error())
		return
	} // if
}

// failed 命令失敗
func (this *build) failed(message string) {
	this.cmd.Println(message)
	this.stop = true
}
