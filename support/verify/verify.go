package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/yinweli/Sheeter/v3/support/verify/codeGo"
)

func main() {
	sheet := sheeter.NewSheeter(newFileLoader())
	assert(sheet.FromData())
	check(sheet, 1000)
	fmt.Println("verify success")
}

func check(sheet *sheeter.Sheeter, threads int) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			actual := sheet.VerifyData.Get(1)
			assert(actual != nil)
			assert(actual.Name1 == 1)
			assert(actual.Name2 == 10)
			assert(actual.Name3 == 11)
			assert(actual.Name4 == 12)
			assert(actual.Name5 == 13)
			assert(actual.Float == 0.01)

			ratio, err := sheeter.RunParse[sheeter.Ratio](actual.Ratio)
			assert(err == nil)
			assert(ratio.Float32() == 0.01)

			duration, err := sheeter.RunParse[sheeter.Duration](actual.Duration)
			assert(err == nil)
			assert(duration.Interval() == time.Hour*26+time.Minute*3+time.Second*4+time.Millisecond*5)

			actual = sheet.VerifyData.Get(2)
			assert(actual != nil)
			assert(actual.Name1 == 2)
			assert(actual.Name2 == 20)
			assert(actual.Name3 == 21)
			assert(actual.Name4 == 22)
			assert(actual.Name5 == 23)
			assert(actual.Float == 0.0001)

			ratio, err = sheeter.RunParse[sheeter.Ratio](actual.Ratio)
			assert(err == nil)
			assert(ratio.Float32() == 0.0001)

			duration, err = sheeter.RunParse[sheeter.Duration](actual.Duration)
			assert(err == nil)
			assert(duration.Interval() == time.Hour*26+time.Minute*3+time.Second*4+time.Millisecond*5)

			actual = sheet.VerifyData.Get(3)
			assert(actual == nil)

			actual = sheet.VerifyData.Get(4)
			assert(actual != nil)
			assert(actual.Name1 == 4)
			assert(actual.Name2 == 40)
			assert(actual.Name3 == 41)
			assert(actual.Name4 == 42)
			assert(actual.Name5 == 43)

			actual = sheet.VerifyData.Get(5)
			assert(actual != nil)
			assert(actual.Name1 == 5)
			assert(actual.Name2 == 50)
			assert(actual.Name3 == 51)
			assert(actual.Name4 == 52)
			assert(actual.Name5 == 53)

			actual = sheet.MergeData.Get(1)
			assert(actual != nil)
			assert(actual.Name1 == 1)
			assert(actual.Name2 == 10)
			assert(actual.Name3 == 11)
			assert(actual.Name4 == 12)
			assert(actual.Name5 == 13)

			actual = sheet.MergeData.Get(2)
			assert(actual != nil)
			assert(actual.Name1 == 2)
			assert(actual.Name2 == 20)
			assert(actual.Name3 == 21)
			assert(actual.Name4 == 22)
			assert(actual.Name5 == 23)

			actual = sheet.MergeData.Get(3)
			assert(actual == nil)

			actual = sheet.MergeData.Get(4)
			assert(actual != nil)
			assert(actual.Name1 == 4)
			assert(actual.Name2 == 40)
			assert(actual.Name3 == 41)
			assert(actual.Name4 == 42)
			assert(actual.Name5 == 43)

			actual = sheet.MergeData.Get(5)
			assert(actual != nil)
			assert(actual.Name1 == 5)
			assert(actual.Name2 == 50)
			assert(actual.Name3 == 51)
			assert(actual.Name4 == 52)
			assert(actual.Name5 == 53)

			waitGroup.Done()
		}()
	} // for

	waitGroup.Wait()
}

func assert(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify failed"))
	} // if
}

func newFileLoader() *fileLoader {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("root path failed"))
	} // if

	return &fileLoader{
		path: filepath.Join(filepath.Dir(root), "json"),
	}
}

type fileLoader struct {
	path string
}

func (this *fileLoader) Load(filename sheeter.FileName) []byte {
	path := filepath.Join(this.path, filename.File())
	data, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("file loader: %w", err))
	}

	return data
}

func (this *fileLoader) Error(name string, err error) {
	panic(fmt.Errorf("file loader: %w", err))
}
