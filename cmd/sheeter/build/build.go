package build

import (
	"fmt"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/builds"
	"github.com/yinweli/Sheeter/internal/codes"
	"github.com/yinweli/Sheeter/internal/utils"
)

const taskSector = 3 // 區段建置的工作數量
const taskEntire = 4 // 全域建置的工作數量

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

	sectors, errs := buildSector(config)

	if len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(fmt.Errorf("build failed, buildSector failed: %w", itor))
		} // for

		return
	} // if

	_, errs = buildEntire(sectors)

	if len(errs) > 0 {
		for _, itor := range errs {
			cmd.Println(fmt.Errorf("build failed, buildEntire failed: %w", itor))
		} // for

		return
	} // if

	cmd.Printf("usage time=%s\n", durafmt.Parse(time.Since(startTime)))
}

// buildSector 區段建置
func buildSector(config *builds.Config) (sectors []*builds.Sector, errs []error) {
	for _, itor := range config.Elements {
		sectors = append(sectors, &builds.Sector{
			Global:  config.Global,
			Element: itor,
		})
	} // for

	count := len(sectors)
	errors := make(chan error) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(&signaler))
	progressbar := progress.AddBar(
		int64(count*taskSector),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("build sector "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	signaler.Add(count)

	for _, itor := range sectors {
		sector := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()
			defer sector.Close()

			if err := builds.SectorInit(sector); err != nil {
				errors <- err
				return
			} // if

			progressbar.Increment()

			if err := builds.SectorJson(sector); err != nil {
				errors <- err
				return
			} // if

			progressbar.Increment()

			if err := builds.SectorJsonSchema(sector); err != nil {
				errors <- err
				return
			} // if

			progressbar.Increment()
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return sectors, errs
}

// buildEntire 全域建置
func buildEntire(sectors []*builds.Sector) (entires []*builds.Entire, errs []error) {
	layoutType, err := builds.MergeSectorLayoutType(sectors)

	if err != nil {
		return entires, []error{err}
	} // if

	for _, itor := range layoutType.TypeNames() {
		if type_ := layoutType.Types(itor); type_ != nil {
			entires = append(entires, &builds.Entire{Type: type_})
		} // if
	} // for

	count := len(entires)
	errors := make(chan error) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(&signaler))
	progressbar := progress.AddBar(
		int64(count*taskEntire),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("build entire "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	signaler.Add(len(entires))

	for _, itor := range entires {
		entire := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()

			if err := builds.EntireJsonCsStruct(entire); err != nil {
				errors <- err
			} // if

			progressbar.Increment()

			if err := builds.EntireJsonCsReader(entire); err != nil {
				errors <- err
			} // if

			progressbar.Increment()

			if err := builds.EntireJsonGoStruct(entire); err != nil {
				errors <- err
			} // if

			progressbar.Increment()

			if err := builds.EntireJsonGoReader(entire); err != nil {
				errors <- err
			} // if

			progressbar.Increment()
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return entires, errs
}
