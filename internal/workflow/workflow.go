package workflow

import (
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewWorkflow 建立工作流, 如果count為0, 會引發無法結束的錯誤
func NewWorkflow(name string, count int) *Workflow {
	workflow := &Workflow{}
	workflow.signaler = utils.NewWaitGroup(count)
	workflow.progress = mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(workflow.signaler))
	workflow.progressbar = workflow.progress.AddBar(
		int64(count),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name(name),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)
	workflow.errors = make(chan error, count)
	workflow.count = count
	return workflow
}

// Workflow 工作流
type Workflow struct {
	signaler    *sync.WaitGroup // waitGroup
	progress    *mpb.Progress   // 進度條管理器
	progressbar *mpb.Bar        // 進度條物件
	errors      chan error      // 錯誤通道
	count       int             // 工作數量
}

// Increment 增加工作進度
func (this *Workflow) Increment() {
	this.progressbar.Increment()
	this.signaler.Done()
	this.count--
}

// Abort 放棄工作
func (this *Workflow) Abort() {
	this.progressbar.Abort(false)
	this.signaler.Add(-this.count)
	this.count = 0
}

// Error 增加錯誤
func (this *Workflow) Error(err error) {
	this.errors <- err
}

// End 結束工作流
func (this *Workflow) End() (errs []error) {
	this.progress.Wait()
	close(this.errors) // 先關閉錯誤通道, 免得下面的for變成無限迴圈

	for itor := range this.errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
