package command

import (
	"io/ioutil"

	"Sheeter/internal/config"
	"Sheeter/internal/logger"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
)

// Build 建立表格命令
func Build() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "build",
		Short: "Build sheet",
		Long:  "Build all the sheet written in the configuration file",
		Args:  cobra.MinimumNArgs(1),
		Run:   build,
	}

	return command
}

// command 建立表格命令
func build(cmd *cobra.Command, args []string) {
	yamlFile, err := ioutil.ReadFile(args[0])

	if err != nil {
		logger.Error("read configuration failed")
		return
	}

	var configuration config.Config

	err = yaml.Unmarshal(yamlFile, configuration)

	if err != nil {
		logger.Error("read configuration failed")
		return
	}
}
