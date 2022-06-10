package build

import (
	"Sheeter/internal/command/build/core"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		RunE:  execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) error {
	config, err := core.ReadConfig(args[0])

	if err != nil {
		return err
	} // if

	// TODO: 測試多執行緒版本

	for _, itor := range config.Elements {
		err := core.NewTask(&config.Global, &itor).Execute()

		if err != nil {
			return err
		} // if
	} // for

	return nil
}
