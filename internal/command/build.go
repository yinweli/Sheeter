package command

import (
	"io/ioutil"

	"Sheeter/internal/logger"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewCommandBuild 建立編譯表格命令
func NewCommandBuild() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "build",
		Short: "NewCommandBuild sheet",
		Long:  "NewCommandBuild all the sheet written in the configuration file",
		Args:  cobra.MinimumNArgs(1),
		Run:   build,
	}

	return command
}

// build 編譯表格命令
func build(cmd *cobra.Command, args []string) {
	yamlFile, err := ioutil.ReadFile(args[0])

	if err != nil {
		logger.Error("read config failed")
		return
	} // if

	var config BuildConfig

	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		logger.Error("read config failed")
		return
	} // if
}
