package build

import (
	"fmt"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal/builds"
	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "build sheet",
		Long:  "generate struct, reader, json data from excel & sheet",
		Run:   execute,
	}
	builds.SetFlags(cmd)
	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, _ []string) {
	startTime := time.Now()

	if utils.ShellExist("gofmt") == false {
		cmd.Println(fmt.Errorf("build failed: `gofmt` not installed"))
		return
	} // if

	if err := tmpls.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("build failed: %w", err))
		return
	} // if

	config := &builds.Config{}

	if err := config.Initialize(cmd); err != nil {
		cmd.Println(fmt.Errorf("build failed: %w", err))
		return
	} // if

	if err := config.Check(); err != nil {
		cmd.Println(fmt.Errorf("build failed: %w", err))
		return
	} // if

	runtime := &builds.Runtime{}

	if errs := builds.Initialize(runtime, config); len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(fmt.Errorf("build failed: %w", itor))
		} // for

		return
	} // id

	if errs := builds.Generate(runtime, config); len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(fmt.Errorf("build failed: %w", itor))
		} // for

		return
	} // if

	if errs := builds.Encoding(runtime, config); len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(fmt.Errorf("build failed: %w", itor))
		} // for

		return
	} // if

	if err := builds.Poststep(runtime, config); err != nil {
		cmd.Println(fmt.Errorf("build failed: %w", err))
		return
	} // if

	for _, itor := range runtime.Sector {
		itor.CloseExcel()
	} // for

	cmd.Printf("usage time=%s\n", durafmt.Parse(time.Since(startTime)))
}
