package main

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/cmd/sheeter/build"
	"github.com/yinweli/Sheeter/cmd/sheeter/code"
	"github.com/yinweli/Sheeter/cmd/sheeter/version"
	"github.com/yinweli/Sheeter/internal"
)

func main() {
	rootCommand := cobra.Command{
		Use:     internal.AppName,
		Long:    "Sheeter used to convert excel file to json file, and generate code of data structure",
		Version: internal.Version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.AddCommand(code.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
