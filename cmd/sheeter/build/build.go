package build

import (
	"fmt"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/v2/sheeter/builds"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "build sheet",
		Long:  "generate reader, sheeter, data from excel & sheet",
		Run:   execute,
	}
	builds.SetFlag(cmd)
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	startTime := time.Now()
	config := &builds.Config{}

	if err := config.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("build: %w", err))
		return
	} // if

	if err := config.Check(); err != nil {
		cmd.Println(fmt.Errorf("build: %w", err))
		return
	} // if

	initializeData, err := builds.Initialize(config)

	if len(err) > 0 {
		for _, itor := range err {
			cmd.Println(fmt.Errorf("build: %w", itor))
		} // for

		return
	} // if

	_, err = builds.Operation(config, initializeData)

	if len(err) > 0 {
		for _, itor := range err {
			cmd.Println(fmt.Errorf("build: %w", itor))
		} // for

		return
	} // if

	_, err = builds.Poststep(config, initializeData)

	if len(err) > 0 {
		for _, itor := range err {
			cmd.Println(fmt.Errorf("build: %w", itor))
		} // for

		return
	} // if

	excels.CloseAll() // 最後關閉所有開啟的excel, sheet
	cmd.Printf("usage time=%v\n", durafmt.Parse(time.Since(startTime)))
}
