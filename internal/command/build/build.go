package build

import (
	"io/ioutil"

	config2 "Sheeter/internal/command/build/config"
	"Sheeter/internal/logger"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewCommand 建立命令
func NewCommand() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "build",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	yamlFile, err := ioutil.ReadFile(args[0])

	if err != nil {
		logger.Error("read config failed")
		return
	} // if

	var config config2.Config

	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		logger.Error("read config failed")
		return
	} // if
}
