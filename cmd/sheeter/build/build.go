package build

import (
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/v3/sheeter/builds"
	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
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
	stdColor := utils.NewStdColor(cmd.OutOrStdout(), cmd.ErrOrStderr())
	startTime := time.Now()
	config := &builds.Config{}

	if err := config.Initialize(cmd); err != nil {
		stdColor.Err("build: %v", err)
		return
	} // if

	if err := config.Check(); err != nil {
		stdColor.Err("build: %v", err)
		return
	} // if

	initializeData, err := builds.Initialize(config)

	if len(err) > 0 {
		for _, itor := range err {
			stdColor.Err("build: %v", itor)
		} // for

		return
	} // if

	_, err = builds.Operation(config, initializeData)

	if len(err) > 0 {
		for _, itor := range err {
			stdColor.Err("build: %v", itor)
		} // for

		return
	} // if

	_, err = builds.Poststep(config, initializeData)

	if len(err) > 0 {
		for _, itor := range err {
			stdColor.Err("build: %v", itor)
		} // for

		return
	} // if

	excels.CloseAll() // 最後關閉所有開啟的excel, sheet
	stdColor.Out("Done! usage time=%v", durafmt.Parse(time.Since(startTime)))
}
