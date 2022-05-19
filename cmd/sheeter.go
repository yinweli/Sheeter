package main

import (
	"Sheeter/internal"
	"Sheeter/internal/command"
	"Sheeter/internal/logger"

	"github.com/spf13/cobra"
)

func main() {
	logInitialize()
	defer logFinalize()

	rootCommand := &cobra.Command{
		Use:  internal.Title,
		Long: "Sheeter used to convert excel file to json file, and generate code of data structure",
	}

	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	rootCommand.AddCommand(command.Version())
	rootCommand.AddCommand(command.Build())

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	} // if
}

// logInitialize 初始化日誌
func logInitialize() {
	err := logger.Initialize()

	if err != nil {
		panic(err)
	} // if
}

// logFinalize 結束日誌
func logFinalize() {
	err := logger.Finalize()

	if err != nil {
		panic(err)
	} // if
}
