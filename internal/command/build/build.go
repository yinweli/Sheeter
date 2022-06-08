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
	core.NewJobs().Flag(cmd)

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

		err := core.ReadSheet(cargo)

		if err != nil {
			return err
		} // if

		jobs := core.NewJobs()
		sheetSize := cargo.Sheets.Size()
		progressValue := jobs.Calc(cmd, sheetSize) + sheetSize // + sheetSize是把讀取表格也算進進度中
		cargo.Progress = core.NewProgress(progressValue, cargo.LogName(), cmd.OutOrStdout())

		err = core.ReadContent(cargo)

		if err != nil {
			return err
		} // if

		err = jobs.Execute(cmd, cargo)

		if err != nil {
			return err
		} // if

		cargo.Progress.Finish()
	} // for

	return nil
}
