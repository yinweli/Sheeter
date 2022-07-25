package util

import (
	"sync"

	"github.com/vbauerster/mpb/v7"
)

const barWidth = 40 // 進度條寬度

// NewMpb 建立多行進度條
func NewMpb(signaler *sync.WaitGroup) *mpb.Progress {
	return mpb.New(mpb.WithWidth(barWidth), mpb.WithWaitGroup(signaler))
}
