package main

import (
	"Sheeter/internal/command/build"
	"Sheeter/internal/command/version"

	"github.com/spf13/cobra"
)

const title = "sheeter" // 程式名稱

func main() {
	rootCommand := cobra.Command{
		Use:  title,
		Long: "Sheeter used to convert excel file to json file, and generate code of data structure",
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
