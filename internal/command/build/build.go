package build

import (
	"Sheeter/internal/command/build/core"
	"Sheeter/internal/util"

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
		{LongName: "json", ShortName: "j", Note: "generate json file", ExecFunc: core.WriteJson},
		{LongName: "cpp", ShortName: "c", Note: "generate cpp file", ExecFunc: core.WriteCpp},
		{LongName: "cs", ShortName: "s", Note: "generate cs file", ExecFunc: core.WriteCs},
		{LongName: "go", ShortName: "g", Note: "generate go file", ExecFunc: core.WriteGo},
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
		progress := util.NewProgressBar(itor.GetFullName(), cmd.OutOrStdout())
		cargo := &core.Cargo{
			Progress: progress,
			Global:   &config.Global,
			Element:  &itor,
		}
		task := executor.Count() + 1 // + 1 是包括讀取設定與表格這項工作
		err = core.ReadSheet(cargo, task)

		if err != nil {
			return err
		} // if

		err = executor.Run(cargo)

		if err != nil {
			return err
		} // if

		_ = progress.Finish()
	} // for

	return nil
}

var executor *core.Executor // 執行者物件
