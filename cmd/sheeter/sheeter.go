package main

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/v3/cmd/sheeter/build"
	"github.com/yinweli/Sheeter/v3/sheeter"
)

func main() {
	rootCommand := &cobra.Command{
		Use:   sheeter.Application,
		Short: "Generate reader and data from Excel",
		Long: `Sheeter is a tool for generating reader code, sheeter logic, and structured data
directly from Excel. It helps automate the workflow of turning
spreadsheet definitions into usable program assets, reducing repetitive tasks and
keeping data consistent across projects. With Sheeter, you can quickly build and
maintain readers, schema definitions, and data exports in a reproducible way`,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.SilenceUsage = true                       // 遇到錯誤不要把整個 usage 洩出來
	rootCommand.SilenceErrors = true                      // 遇到錯誤不要把整個 usage 洩出來
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
