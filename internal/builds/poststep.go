package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// Poststep 後製
func Poststep(runtime *Runtime, config *Config) error {
	tasks := []func(*Runtime) error{}

	if config.Global.ExportJson {
		tasks = append(
			tasks,
			poststepJsonCsDepot,
			poststepJsonGoDepot,
		)
	} // if

	if config.Global.ExportProto {
		tasks = append(
			tasks,
			poststepProtoCsDepot,
			poststepProtoGoDepot,
			poststepProtoCsBat,
			poststepProtoCsSh,
			poststepProtoGoBat,
			poststepProtoGoSh,
		)
	} // if

	totalCount := len(tasks)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth))
	progressbar := progress.AddBar(
		int64(totalCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("poststep "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range tasks {
		if err := itor(runtime); err != nil {
			return fmt.Errorf("poststep failed: %w", err)
		} // if

		progressbar.Increment()
	} // for

	return nil
}
