package build

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

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
	}

	build.readConfig()

	if build.err != nil {
		cmd.Println(build.err)
		return
	} // if
}

// build 編譯命令
type build struct {
	cmd    *cobra.Command // 命令物件
	args   []string       // 參數列表
	err    error          // 錯誤物件
	config *config.Config // 設定物件
}

// readConfig 讀取設定
func (this *build) readConfig() {
	bytes, err := ioutil.ReadFile(this.args[0])

	if err != nil {
		this.err = err
		return
	} // if

	this.config = &config.Config{}
	err = yaml.Unmarshal(bytes, this.config)

	if err != nil {
		this.err = err
		return
	} // if

	err = this.config.Check()

	if err != nil {
		this.err = err
		return
	} // if
}
