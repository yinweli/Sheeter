package build

import (
	"fmt"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/builds"
	"github.com/yinweli/Sheeter/internal/codes"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "build sheet",
		Long:  "generate json, schema, struct code, reader code from excel & sheet",
		Run:   execute,
	}
	builds.SetFlags(cmd)
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	startTime := time.Now()

	if utils.ShellExist("gofmt") == false {
		cmd.Println(fmt.Errorf("build failed, `gofmt` not installed"))
		return
	} // if

	if utils.ShellExist("quicktype") == false {
		cmd.Println(fmt.Errorf("build failed, `quicktype` not installed"))
		return
	} // if

	if err := codes.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("build failed, code initialize failed: %w", err))
		return
	} // if

	config := &builds.Config{}

	if err := config.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("build failed, config initialize failed: %w", err))
		return
	} // if

	if err := config.Check(); err != nil {
		cmd.Println(fmt.Errorf("build failed, config check failed: %w", err))
		return
	} // if

	if errs := builds.BuildSector(config); len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(itor)
		} // for

		return
	} // if

	cmd.Printf("usage time=%s\n", durafmt.Parse(time.Since(startTime)))
}
