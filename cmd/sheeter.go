package main

import (
	"Sheeter/internal"
	"Sheeter/internal/command/build"
	"Sheeter/internal/command/version"

	"github.com/spf13/cobra"
)

func main() {
	rootCommand := cobra.Command{
		Use:  internal.Title,
		Long: "Sheeter used to convert excel file to json file, and generate code of data structure",
	}
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
