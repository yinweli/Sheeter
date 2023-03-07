package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	sheeterJson "github.com/yinweli/Sheeter/support/verifygo/target/json/codeGo"
)

func verifyJsonFrom(threads int) {
	loader := newJsonFileLoader()
	depot := sheeterJson.NewDepot(loader)

	assertJson(depot.FromData())
	verifyJson(depot, threads)

	fmt.Println("verify json from: success")
}

func verifyJsonMerge(threads int) {
	loader := newJsonFileLoader()
	depot := sheeterJson.NewDepot(loader)

	assertJson(depot.MergeData())
	verifyJson(depot, threads)

	fmt.Println("verify json merge: success")
}

func verifyJson(depot *sheeterJson.Depot, threads int) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			for i := int32(1); i <= 100; i++ {
				actual1, ok := depot.VerifyData1.Get(i)
				assertJson(ok)
				assertJson(actual1.Key == i)
				assertJson(actual1.Hide == false)
				assertJson(actual1.Enable == (i%2 == 1))
				assertJson(actual1.Name == fmt.Sprintf("名稱%d", i))
				assertJson(actual1.Reward.Desc == fmt.Sprintf("獎勵說明%d", i))
				assertJson(actual1.Reward.Gold == i*2)
				assertJson(actual1.Reward.Diamond == i*3)
				assertJson(actual1.Reward.Crystal == i*4)
				assertJson(actual1.Reward.FelIron == i*5)
				assertJson(actual1.Reward.Atium == i*6)
				assertJson(len(actual1.Reward.Item) == 3)
				assertJson(actual1.Reward.Item[0].ItemID == 1000+i)
				assertJson(actual1.Reward.Item[0].Type == 0)
				assertJson(actual1.Reward.Item[0].Count == i)
				assertJson(actual1.Reward.Item[1].ItemID == 10000+i)
				assertJson(actual1.Reward.Item[1].Type == 1)
				assertJson(actual1.Reward.Item[1].Count == i)
				assertJson(actual1.Reward.Item[2].ItemID == 100000+i)
				assertJson(actual1.Reward.Item[2].Type == 2)
				assertJson(actual1.Reward.Item[2].Count == i)

				actual2, ok := depot.VerifyData2.Get(fmt.Sprint(i))
				assertJson(ok)
				assertJson(actual2.Key == fmt.Sprint(i))
				assertJson(actual2.Hide == false)
				assertJson(actual2.Enable == (i%2 == 1))
				assertJson(actual2.Name == fmt.Sprintf("名稱%d", i))
				assertJson(actual2.Reward.Desc == fmt.Sprintf("獎勵說明%d", i))
				assertJson(actual2.Reward.Gold == i*2)
				assertJson(actual2.Reward.Diamond == i*3)
				assertJson(actual2.Reward.Crystal == 0)
				assertJson(actual2.Reward.FelIron == 0)
				assertJson(actual2.Reward.Atium == 0)
				assertJson(len(actual2.Reward.Item) == 3)
				assertJson(actual2.Reward.Item[0].ItemID == 1000+i)
				assertJson(actual2.Reward.Item[0].Type == 0)
				assertJson(actual2.Reward.Item[0].Count == i)
				assertJson(actual2.Reward.Item[1].ItemID == 10000+i)
				assertJson(actual2.Reward.Item[1].Type == 1)
				assertJson(actual2.Reward.Item[1].Count == i)
				assertJson(actual2.Reward.Item[2].ItemID == 100000+i)
				assertJson(actual2.Reward.Item[2].Type == 2)
				assertJson(actual2.Reward.Item[2].Count == i)
			} // for

			_, ok := depot.VerifyData1.Get(101)
			assertJson(ok == false)

			_, ok = depot.VerifyData2.Get(fmt.Sprint(101))
			assertJson(ok == false)

			waitGroup.Done()
		}()
	} // for

	waitGroup.Wait()
}

func assertJson(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify json: verify failed"))
	} // if
}

type jsonFileLoader struct {
	path string
}

func newJsonFileLoader() *jsonFileLoader {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("root path failed"))
	} // if

	return &jsonFileLoader{
		path: filepath.Join(filepath.Dir(root), "target", "json", "data"),
	}
}

func (this *jsonFileLoader) Error(name string, err error) {
	panic(fmt.Errorf("%s: json file load failed: %w", name, err))
}

func (this *jsonFileLoader) Load(filename sheeterJson.FileName) []byte {
	path := filepath.Join(this.path, filename.File())
	data, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("%s: json file load failed: %w", filename.Name(), err))
	}

	return data
}
