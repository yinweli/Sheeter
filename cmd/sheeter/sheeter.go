package main

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/cmd/sheeter/build"
	"github.com/yinweli/Sheeter/cmd/sheeter/tmpl"
	"github.com/yinweli/Sheeter/cmd/sheeter/version"
	"github.com/yinweli/Sheeter/internal"
)

func main() {
	rootCommand := cobra.Command{
		Use:     internal.AppName,
		Long:    "Sheeter used to generate struct, reader, json data from excel & sheet",
		Version: internal.Version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.AddCommand(tmpl.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
