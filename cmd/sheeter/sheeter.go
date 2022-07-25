package main

import (
	"github.com/spf13/cobra"
	"github.com/yinweli/Sheeter/internal/app/sheeter/build"
	"github.com/yinweli/Sheeter/internal/app/sheeter/version"
	"github.com/yinweli/Sheeter/internal/pkg"
)

func main() {
	rootCommand := cobra.Command{
		Use:     pkg.Title,
		Long:    "Sheeter used to convert excel file to json file, and generate code of data structure",
		Version: pkg.Version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
