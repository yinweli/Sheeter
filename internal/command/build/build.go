package build

import (
	"Sheeter/internal/command/build/builder"
	"Sheeter/internal/command/build/config"

	"github.com/spf13/cobra"
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
	err := builder.ReadConfig(args[0], &buildConfig)

	if err != nil {
		cmd.Println(err)
		return
	} // if

	ok, errs := buildConfig.Check()

	if ok == false {
		for _, itor := range errs {
			cmd.Println(itor)
		} // for

		return
	} // if

	for _, itor := range buildConfig.Elements {
		cargo := builder.Cargo{
			Output:  cmd.OutOrStdout(),
			Global:  &buildConfig.Global,
			Element: &itor,
		}
		err := builder.ReadExcel(&cargo)

		if err != nil {
			cmd.Println(err)
			continue
		} // if
	} // for
}
