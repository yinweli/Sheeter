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
	executor = core.NewExecutor(cmd, []core.ExecData{
		{LongName: "json", ShortName: "j", Note: "generate json file", ExecFunc: core.WriteJson, Progress: true},
		{LongName: "cpp", ShortName: "c", Note: "generate cpp file", ExecFunc: core.WriteCpp, Progress: false},
		{LongName: "cs", ShortName: "s", Note: "generate cs file", ExecFunc: core.WriteCs, Progress: false},
		{LongName: "go", ShortName: "g", Note: "generate go file", ExecFunc: core.WriteGo, Progress: false},
	})

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) error {
	config, err := core.ReadConfig(args[0])

	if err != nil {
		return err
	} // if

	for _, itor := range config.Elements {
		cargo := &core.Cargo{
			Global:  &config.Global,
			Element: &itor,
		}

		err := core.ReadSheet(cargo, executor.Progress()+1, cmd.OutOrStdout()) // +1是把讀取表格也算進進度中

		if err != nil {
			return err
		} // if

		err = core.ReadContent(cargo)

		if err != nil {
			return err
		} // if

		err = executor.Run(cargo)

		if err != nil {
			return err
		} // if

		cargo.Progress.Finish()
	} // for

	return nil
}

var executor *core.Executor // 執行者物件
