package main

import (
	"github.com/yinweli/Sheeter/app/buildall"
	"github.com/yinweli/Sheeter/app/version"
	"github.com/yinweli/Sheeter/internal"

	"github.com/spf13/cobra"
)

func main() {
	rootCommand := cobra.Command{
		Use:     internal.Title,
		Long:    "Sheeter used to convert excel file to json file, and generate code of data structure",
		Version: internal.Version,
	}
	rootCommand.AddCommand(buildall.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
