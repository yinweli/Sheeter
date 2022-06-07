package main

import (
	"Sheeter/internal/command/build"

	"github.com/spf13/cobra"
)

const title = "sheeter" // 程式名稱
const version = "1.0.0" // 版本字串

func main() {
	rootCommand := cobra.Command{
		Use:     title,
		Long:    "Sheeter used to convert excel file to json file, and generate code of data structure",
		Version: version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
