package build

import (
	"io/ioutil"

	"Sheeter/internal/command/build/config"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewCommand 建立命令
func NewCommand() (cmd *cobra.Command) {
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
	filename := args[0]
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		cmd.Println(err)
		return
	} // if

	buildConfig := config.Config{}
	err = yaml.Unmarshal(file, &buildConfig)

	if err != nil {
		cmd.Println(err)
		return
	} // if
}
