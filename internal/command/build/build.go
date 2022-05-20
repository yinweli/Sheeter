package build

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

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
	err := readBuildConfig(args[0], &buildConfig)

	if err != nil {
		cmd.Println(err)
		return
	} // if
}

// readBuildConfig 讀取編譯設置
func readBuildConfig(filename string, buildConfig *config.Config) error {
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
