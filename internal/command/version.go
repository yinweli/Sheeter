package command

import (
	"fmt"

	"Sheeter/internal"

	"github.com/spf13/cobra"
)

// Version 顯示版本命令
func Version() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  "Show version",
		Run:   version,
	}

	return command
}

// version 顯示版本命令
func version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s %s", internal.Title, internal.Version)
}
