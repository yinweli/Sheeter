package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	sheeterJson "github.com/yinweli/Sheeter/support/verifygo/target/json/codeGo"
)

func verifyJsonFrom() {
	loader := newJsonFileLoader()
	depot := sheeterJson.NewDepot(loader)

	assertJson(depot.FromData())

	actual1, ok := depot.VerifyData1.Get(1)
	assertJson(ok)
	assertJson(actual1.Key == 1)
	assertJson(actual1.Hide == false)
	assertJson(actual1.Enable == true)
	assertJson(actual1.Name == "名稱1")
	assertJson(actual1.Reward.Desc == "獎勵說明1")
	assertJson(actual1.Reward.Gold == 100)
	assertJson(actual1.Reward.Diamond == 10)
	assertJson(actual1.Reward.Crystal == 199)
	assertJson(actual1.Reward.FelIron == 5)
	assertJson(actual1.Reward.Atium == 1)
	assertJson(len(actual1.Reward.Item) == 3)
	assertJson(actual1.Reward.Item[0].ItemID == 10001)
	assertJson(actual1.Reward.Item[0].Type == 1)
	assertJson(actual1.Reward.Item[0].Count == 1)
	assertJson(actual1.Reward.Item[1].ItemID == 10002)
	assertJson(actual1.Reward.Item[1].Type == 2)
	assertJson(actual1.Reward.Item[1].Count == 2)
	assertJson(actual1.Reward.Item[2].ItemID == 0)
	assertJson(actual1.Reward.Item[2].Type == 0)
	assertJson(actual1.Reward.Item[2].Count == 0)

	actual1, ok = depot.VerifyData1.Get(2)
	assertJson(ok)
	assertJson(actual1.Key == 2)
	assertJson(actual1.Hide == false)
	assertJson(actual1.Enable == false)
	assertJson(actual1.Name == "名稱2")
	assertJson(actual1.Reward.Desc == "獎勵說明2")
	assertJson(actual1.Reward.Gold == 200)
	assertJson(actual1.Reward.Diamond == 20)
	assertJson(actual1.Reward.Crystal == 299)
	assertJson(actual1.Reward.FelIron == 10)
	assertJson(actual1.Reward.Atium == 2)
	assertJson(len(actual1.Reward.Item) == 3)
	assertJson(actual1.Reward.Item[0].ItemID == 10001)
	assertJson(actual1.Reward.Item[0].Type == 1)
	assertJson(actual1.Reward.Item[0].Count == 1)
	assertJson(actual1.Reward.Item[1].ItemID == 10002)
	assertJson(actual1.Reward.Item[1].Type == 2)
	assertJson(actual1.Reward.Item[1].Count == 2)
	assertJson(actual1.Reward.Item[2].ItemID == 10003)
	assertJson(actual1.Reward.Item[2].Type == 3)
	assertJson(actual1.Reward.Item[2].Count == 3)

	actual1, ok = depot.VerifyData1.Get(3)
	assertJson(ok == false)

	actual2, ok := depot.VerifyData2.Get(1)
	assertJson(ok)
	assertJson(actual2.Key == 1)
	assertJson(actual2.Hide == false)
	assertJson(actual2.Enable == true)
	assertJson(actual2.Name == "名稱1")
	assertJson(actual2.Reward.Desc == "獎勵說明1")
	assertJson(actual2.Reward.Gold == 100)
	assertJson(actual2.Reward.Diamond == 10)
	assertJson(actual2.Reward.Crystal == 0)
	assertJson(actual2.Reward.FelIron == 0)
	assertJson(actual2.Reward.Atium == 0)
	assertJson(len(actual2.Reward.Item) == 3)
	assertJson(actual2.Reward.Item[0].ItemID == 10001)
	assertJson(actual2.Reward.Item[0].Type == 1)
	assertJson(actual2.Reward.Item[0].Count == 1)
	assertJson(actual2.Reward.Item[1].ItemID == 10002)
	assertJson(actual2.Reward.Item[1].Type == 2)
	assertJson(actual2.Reward.Item[1].Count == 2)
	assertJson(actual2.Reward.Item[2].ItemID == 0)
	assertJson(actual2.Reward.Item[2].Type == 0)
	assertJson(actual2.Reward.Item[2].Count == 0)

	actual2, ok = depot.VerifyData2.Get(2)
	assertJson(ok)
	assertJson(actual2.Key == 2)
	assertJson(actual2.Hide == false)
	assertJson(actual2.Enable == false)
	assertJson(actual2.Name == "名稱2")
	assertJson(actual2.Reward.Desc == "獎勵說明2")
	assertJson(actual2.Reward.Gold == 200)
	assertJson(actual2.Reward.Diamond == 20)
	assertJson(actual2.Reward.Crystal == 0)
	assertJson(actual2.Reward.FelIron == 0)
	assertJson(actual2.Reward.Atium == 0)
	assertJson(len(actual2.Reward.Item) == 3)
	assertJson(actual2.Reward.Item[0].ItemID == 10001)
	assertJson(actual2.Reward.Item[0].Type == 1)
	assertJson(actual2.Reward.Item[0].Count == 1)
	assertJson(actual2.Reward.Item[1].ItemID == 10002)
	assertJson(actual2.Reward.Item[1].Type == 2)
	assertJson(actual2.Reward.Item[1].Count == 2)
	assertJson(actual2.Reward.Item[2].ItemID == 10003)
	assertJson(actual2.Reward.Item[2].Type == 3)
	assertJson(actual2.Reward.Item[2].Count == 3)

	actual2, ok = depot.VerifyData2.Get(3)
	assertJson(ok == false)

	fmt.Println("verify json from: success")
}

func verifyJsonMerge() {
	loader := newJsonFileLoader()
	depot := sheeterJson.NewDepot(loader)

	assertJson(depot.MergeData())

	actual1, ok := depot.VerifyData1.Get(1)
	assertJson(ok)
	assertJson(actual1.Key == 1)
	assertJson(actual1.Hide == false)
	assertJson(actual1.Enable == true)
	assertJson(actual1.Name == "名稱1")
	assertJson(actual1.Reward.Desc == "獎勵說明1")
	assertJson(actual1.Reward.Gold == 100)
	assertJson(actual1.Reward.Diamond == 10)
	assertJson(actual1.Reward.Crystal == 199)
	assertJson(actual1.Reward.FelIron == 5)
	assertJson(actual1.Reward.Atium == 1)
	assertJson(len(actual1.Reward.Item) == 3)
	assertJson(actual1.Reward.Item[0].ItemID == 10001)
	assertJson(actual1.Reward.Item[0].Type == 1)
	assertJson(actual1.Reward.Item[0].Count == 1)
	assertJson(actual1.Reward.Item[1].ItemID == 10002)
	assertJson(actual1.Reward.Item[1].Type == 2)
	assertJson(actual1.Reward.Item[1].Count == 2)
	assertJson(actual1.Reward.Item[2].ItemID == 0)
	assertJson(actual1.Reward.Item[2].Type == 0)
	assertJson(actual1.Reward.Item[2].Count == 0)

	actual1, ok = depot.VerifyData1.Get(2)
	assertJson(ok)
	assertJson(actual1.Key == 2)
	assertJson(actual1.Hide == false)
	assertJson(actual1.Enable == false)
	assertJson(actual1.Name == "名稱2")
	assertJson(actual1.Reward.Desc == "獎勵說明2")
	assertJson(actual1.Reward.Gold == 200)
	assertJson(actual1.Reward.Diamond == 20)
	assertJson(actual1.Reward.Crystal == 299)
	assertJson(actual1.Reward.FelIron == 10)
	assertJson(actual1.Reward.Atium == 2)
	assertJson(len(actual1.Reward.Item) == 3)
	assertJson(actual1.Reward.Item[0].ItemID == 10001)
	assertJson(actual1.Reward.Item[0].Type == 1)
	assertJson(actual1.Reward.Item[0].Count == 1)
	assertJson(actual1.Reward.Item[1].ItemID == 10002)
	assertJson(actual1.Reward.Item[1].Type == 2)
	assertJson(actual1.Reward.Item[1].Count == 2)
	assertJson(actual1.Reward.Item[2].ItemID == 10003)
	assertJson(actual1.Reward.Item[2].Type == 3)
	assertJson(actual1.Reward.Item[2].Count == 3)

	actual1, ok = depot.VerifyData1.Get(3)
	assertJson(ok == false)

	actual2, ok := depot.VerifyData2.Get(1)
	assertJson(ok)
	assertJson(actual2.Key == 1)
	assertJson(actual2.Hide == false)
	assertJson(actual2.Enable == true)
	assertJson(actual2.Name == "名稱1")
	assertJson(actual2.Reward.Desc == "獎勵說明1")
	assertJson(actual2.Reward.Gold == 100)
	assertJson(actual2.Reward.Diamond == 10)
	assertJson(actual2.Reward.Crystal == 0)
	assertJson(actual2.Reward.FelIron == 0)
	assertJson(actual2.Reward.Atium == 0)
	assertJson(len(actual2.Reward.Item) == 3)
	assertJson(actual2.Reward.Item[0].ItemID == 10001)
	assertJson(actual2.Reward.Item[0].Type == 1)
	assertJson(actual2.Reward.Item[0].Count == 1)
	assertJson(actual2.Reward.Item[1].ItemID == 10002)
	assertJson(actual2.Reward.Item[1].Type == 2)
	assertJson(actual2.Reward.Item[1].Count == 2)
	assertJson(actual2.Reward.Item[2].ItemID == 0)
	assertJson(actual2.Reward.Item[2].Type == 0)
	assertJson(actual2.Reward.Item[2].Count == 0)

	actual2, ok = depot.VerifyData2.Get(2)
	assertJson(ok)
	assertJson(actual2.Key == 2)
	assertJson(actual2.Hide == false)
	assertJson(actual2.Enable == false)
	assertJson(actual2.Name == "名稱2")
	assertJson(actual2.Reward.Desc == "獎勵說明2")
	assertJson(actual2.Reward.Gold == 200)
	assertJson(actual2.Reward.Diamond == 20)
	assertJson(actual2.Reward.Crystal == 0)
	assertJson(actual2.Reward.FelIron == 0)
	assertJson(actual2.Reward.Atium == 0)
	assertJson(len(actual2.Reward.Item) == 3)
	assertJson(actual2.Reward.Item[0].ItemID == 10001)
	assertJson(actual2.Reward.Item[0].Type == 1)
	assertJson(actual2.Reward.Item[0].Count == 1)
	assertJson(actual2.Reward.Item[1].ItemID == 10002)
	assertJson(actual2.Reward.Item[1].Type == 2)
	assertJson(actual2.Reward.Item[1].Count == 2)
	assertJson(actual2.Reward.Item[2].ItemID == 10003)
	assertJson(actual2.Reward.Item[2].Type == 3)
	assertJson(actual2.Reward.Item[2].Count == 3)

	actual2, ok = depot.VerifyData2.Get(3)
	assertJson(ok == false)

	fmt.Println("verify json merge: success")
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

func (this *jsonFileLoader) Load(name, ext, fullname string) []byte {
	path := filepath.Join(this.path, fullname)
	data, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("%s: json file load failed: %w", name, err))
	}

	return data
}

func assertJson(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify json: verify failed"))
	} // if
}
