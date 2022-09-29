package main

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	sheeterJson "github.com/yinweli/Sheeter/verify/verifygo/target/json/codeGo"
)

func verifyJson(rootPath string) {
	path := filepath.Join(rootPath, "target", internal.PathJson, internal.PathData)
	verifyJson1(path)
	verifyJson2(path)
}

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyJson1(path string) {
	reader := sheeterJson.VerifyData1Reader{}

	if err := reader.FromPathHalf(path); err != nil {
		panic(fmt.Errorf("verify json: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertJson(ok)
	assertJson(actual.Key == 1)
	assertJson(actual.Enable == true)
	assertJson(actual.Name == "名稱1")
	assertJson(actual.Reward.Desc == "獎勵說明1")
	assertJson(actual.Reward.Gold == 100)
	assertJson(actual.Reward.Diamond == 10)
	assertJson(actual.Reward.Crystal == 199)
	assertJson(actual.Reward.FelIron == 5)
	assertJson(actual.Reward.Atium == 1)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 1)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 2)
	assertJson(actual.Reward.Item[1].Count == 2)
	assertJson(actual.Reward.Item[2].ItemID == 0)
	assertJson(actual.Reward.Item[2].Type == 0)
	assertJson(actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[2]
	assertJson(ok)
	assertJson(actual.Key == 2)
	assertJson(actual.Enable == false)
	assertJson(actual.Name == "名稱2")
	assertJson(actual.Reward.Desc == "獎勵說明2")
	assertJson(actual.Reward.Gold == 200)
	assertJson(actual.Reward.Diamond == 20)
	assertJson(actual.Reward.Crystal == 299)
	assertJson(actual.Reward.FelIron == 10)
	assertJson(actual.Reward.Atium == 2)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 1)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 2)
	assertJson(actual.Reward.Item[1].Count == 2)
	assertJson(actual.Reward.Item[2].ItemID == 10003)
	assertJson(actual.Reward.Item[2].Type == 3)
	assertJson(actual.Reward.Item[2].Count == 3)

	actual, ok = reader.Datas[3]
	assertJson(ok == false)

	fmt.Println("verify json: success")
}

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyJson2(path string) {
	reader := sheeterJson.VerifyData2Reader{}

	if err := reader.FromPathHalf(path); err != nil {
		panic(fmt.Errorf("verify json: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertJson(ok)
	assertJson(actual.Key == 1)
	assertJson(actual.Enable == true)
	assertJson(actual.Name == "名稱1")
	assertJson(actual.Reward.Desc == "獎勵說明1")
	assertJson(actual.Reward.Gold == 100)
	assertJson(actual.Reward.Diamond == 10)
	assertJson(actual.Reward.Crystal == 0)
	assertJson(actual.Reward.FelIron == 0)
	assertJson(actual.Reward.Atium == 0)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 1)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 2)
	assertJson(actual.Reward.Item[1].Count == 2)
	assertJson(actual.Reward.Item[2].ItemID == 0)
	assertJson(actual.Reward.Item[2].Type == 0)
	assertJson(actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[2]
	assertJson(ok)
	assertJson(actual.Key == 2)
	assertJson(actual.Enable == false)
	assertJson(actual.Name == "名稱2")
	assertJson(actual.Reward.Desc == "獎勵說明2")
	assertJson(actual.Reward.Gold == 200)
	assertJson(actual.Reward.Diamond == 20)
	assertJson(actual.Reward.Crystal == 0)
	assertJson(actual.Reward.FelIron == 0)
	assertJson(actual.Reward.Atium == 0)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 1)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 2)
	assertJson(actual.Reward.Item[1].Count == 2)
	assertJson(actual.Reward.Item[2].ItemID == 10003)
	assertJson(actual.Reward.Item[2].Type == 3)
	assertJson(actual.Reward.Item[2].Count == 3)

	actual, ok = reader.Datas[3]
	assertJson(ok == false)

	fmt.Println("verify json: success")
}

func assertJson(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify json: verify failed"))
	} // if
}
