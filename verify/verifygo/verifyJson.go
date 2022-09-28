package main

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	sheeterJson "github.com/yinweli/Sheeter/verify/verifygo/target/json/codeGo"
)

func verifyJson(rootPath string) {
	reader := sheeterJson.VerifyData1Reader{}

	if err := reader.FromPathHalf(filepath.Join(rootPath, "target", internal.PathJson, internal.PathData)); err != nil {
		panic(fmt.Errorf("verify json: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertJson(ok)
	assertJson(actual.Key == 1)
	assertJson(actual.Name == "名稱1")
	assertJson(actual.Enable == true)
	assertJson(actual.Reward.Atium == 2)
	assertJson(actual.Reward.Crystal == 120)
	assertJson(actual.Reward.Diamond == 10)
	assertJson(actual.Reward.FelIron == 6)
	assertJson(actual.Reward.Gold == 500)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 10)
	assertJson(actual.Reward.Item[1].ItemID == 0)
	assertJson(actual.Reward.Item[1].Type == 0)
	assertJson(actual.Reward.Item[1].Count == 0)
	assertJson(actual.Reward.Item[2].ItemID == 0)
	assertJson(actual.Reward.Item[2].Type == 0)
	assertJson(actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[2]
	assertJson(ok)
	assertJson(actual.Key == 2)
	assertJson(actual.Name == "名稱2")
	assertJson(actual.Enable == true)
	assertJson(actual.Reward.Atium == 2)
	assertJson(actual.Reward.Crystal == 135)
	assertJson(actual.Reward.Diamond == 12)
	assertJson(actual.Reward.FelIron == 8)
	assertJson(actual.Reward.Gold == 550)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 10)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 1)
	assertJson(actual.Reward.Item[1].Count == 5)
	assertJson(actual.Reward.Item[2].ItemID == 0)
	assertJson(actual.Reward.Item[2].Type == 0)
	assertJson(actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[3]
	assertJson(ok)
	assertJson(actual.Key == 3)
	assertJson(actual.Name == "名稱3")
	assertJson(actual.Enable == false)
	assertJson(actual.Reward.Atium == 3)
	assertJson(actual.Reward.Crystal == 150)
	assertJson(actual.Reward.Diamond == 14)
	assertJson(actual.Reward.FelIron == 10)
	assertJson(actual.Reward.Gold == 600)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 10)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 1)
	assertJson(actual.Reward.Item[1].Count == 5)
	assertJson(actual.Reward.Item[2].ItemID == 10003)
	assertJson(actual.Reward.Item[2].Type == 1)
	assertJson(actual.Reward.Item[2].Count == 2)

	actual, ok = reader.Datas[4]
	assertJson(ok)
	assertJson(actual.Key == 4)
	assertJson(actual.Name == "名稱4")
	assertJson(actual.Enable == false)
	assertJson(actual.Reward.Atium == 3)
	assertJson(actual.Reward.Crystal == 165)
	assertJson(actual.Reward.Diamond == 16)
	assertJson(actual.Reward.FelIron == 12)
	assertJson(actual.Reward.Gold == 650)
	assertJson(len(actual.Reward.Item) == 3)
	assertJson(actual.Reward.Item[0].ItemID == 10001)
	assertJson(actual.Reward.Item[0].Type == 1)
	assertJson(actual.Reward.Item[0].Count == 10)
	assertJson(actual.Reward.Item[1].ItemID == 10002)
	assertJson(actual.Reward.Item[1].Type == 1)
	assertJson(actual.Reward.Item[1].Count == 5)
	assertJson(actual.Reward.Item[2].ItemID == 10003)
	assertJson(actual.Reward.Item[2].Type == 1)
	assertJson(actual.Reward.Item[2].Count == 3)

	actual, ok = reader.Datas[5]
	assertJson(ok == false)

	fmt.Println("verify json: success")
}

func assertJson(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify json: verify failed"))
	} // if
}
