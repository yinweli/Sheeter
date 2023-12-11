package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/yinweli/Sheeter/v2/support/verify/codeGo"
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
			assert(actual.Pkey == 1)
			assert(actual.Name1 == 10)
			assert(actual.Name2 == 11)
			assert(actual.Name3 == 12)
			assert(actual.Name4 == 13)

			actual = sheet.VerifyData.Get(2)
			assert(actual != nil)
			assert(actual.Pkey == 2)
			assert(actual.Name1 == 20)
			assert(actual.Name2 == 21)
			assert(actual.Name3 == 22)
			assert(actual.Name4 == 23)

			actual = sheet.VerifyData.Get(3)
			assert(actual == nil)

			actual = sheet.VerifyData.Get(4)
			assert(actual != nil)
			assert(actual.Pkey == 4)
			assert(actual.Name1 == 40)
			assert(actual.Name2 == 41)
			assert(actual.Name3 == 42)
			assert(actual.Name4 == 43)

			actual = sheet.VerifyData.Get(5)
			assert(actual != nil)
			assert(actual.Pkey == 5)
			assert(actual.Name1 == 50)
			assert(actual.Name2 == 51)
			assert(actual.Name3 == 52)
			assert(actual.Name4 == 53)

			actual = sheet.MergeData.Get(1)
			assert(actual != nil)
			assert(actual.Pkey == 1)
			assert(actual.Name1 == 10)
			assert(actual.Name2 == 11)
			assert(actual.Name3 == 12)
			assert(actual.Name4 == 13)

			actual = sheet.MergeData.Get(2)
			assert(actual != nil)
			assert(actual.Pkey == 2)
			assert(actual.Name1 == 20)
			assert(actual.Name2 == 21)
			assert(actual.Name3 == 22)
			assert(actual.Name4 == 23)

			actual = sheet.MergeData.Get(3)
			assert(actual == nil)

			actual = sheet.MergeData.Get(4)
			assert(actual != nil)
			assert(actual.Pkey == 4)
			assert(actual.Name1 == 40)
			assert(actual.Name2 == 41)
			assert(actual.Name3 == 42)
			assert(actual.Name4 == 43)

			actual = sheet.MergeData.Get(5)
			assert(actual != nil)
			assert(actual.Pkey == 5)
			assert(actual.Name1 == 50)
			assert(actual.Name2 == 51)
			assert(actual.Name3 == 52)
			assert(actual.Name4 == 53)

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
