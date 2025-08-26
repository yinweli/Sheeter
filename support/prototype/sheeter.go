package sheeter

import (
	"sync"
	"sync/atomic"
)

// NewSheeter 建立表格資料
func NewSheeter(loader Loader) *Sheeter {
	sheeter := &Sheeter{}
	sheeter.loader = loader
	sheeter.progress = NewProgress()
	return sheeter
}

// Sheeter 表格資料
type Sheeter struct {
	loader   Loader         // 裝載器物件
	progress *Progress      // 進度物件
	Alone0   HandmadeReader // 獨立表格說明
	Alone1   HandmadeReader // 獨立表格說明
	Merge0   HandmadeReader // 合併表格說明
	Merge1   HandmadeReader // 合併表格說明
}

// FromData 讀取資料處理
func (this *Sheeter) FromData() bool {
	this.progress.Reset()

	if this.loader == nil {
		return false
	} // if

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(4)
	result := atomic.Bool{}
	result.Store(true)

	for _, itor := range []Reader{
		&this.Alone0,
		&this.Alone1,
	} {
		tmpl := itor

		go func() {
			defer waitGroup.Done()
			filename := tmpl.FileName()
			data := this.loader.Load(filename)

			if len(data) == 0 {
				return
			} // if

			if err := tmpl.FromData(data, true, this.progress); err != nil {
				this.loader.Error(filename.File(), err)
				result.Store(false)
			} // if
		}()
	} // for

	go func() {
		defer waitGroup.Done()

		for i, itor := range []Reader{
			&this.Alone0,
			&this.Alone1,
		} {
			filename := itor.FileName()
			data := this.loader.Load(filename)

			if len(data) == 0 {
				continue
			} // if

			if err := this.Merge0.FromData(data, i == 0, this.progress); err != nil {
				this.loader.Error(filename.File(), err)
				result.Store(false)
			} // if
		} // for
	}()

	go func() {
		defer waitGroup.Done()

		for i, itor := range []Reader{
			&this.Alone0,
			&this.Alone1,
		} {
			filename := itor.FileName()
			data := this.loader.Load(filename)

			if len(data) == 0 {
				continue
			} // if

			if err := this.Merge1.FromData(data, i == 0, this.progress); err != nil {
				this.loader.Error(filename.File(), err)
				result.Store(false)
			} // if
		} // for
	}()

	waitGroup.Wait()
	this.progress.Complete()
	return result.Load()
}

// Clear 清除資料
func (this *Sheeter) Clear() {
	this.progress.Reset()
	this.Alone0.Clear()
	this.Alone1.Clear()
	this.Merge0.Clear()
	this.Merge1.Clear()
}

// Progress 取得進度值
func (this *Sheeter) Progress() float32 {
	return this.progress.Get()
}
