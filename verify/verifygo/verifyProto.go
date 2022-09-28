package main

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	sheeterProto "github.com/yinweli/Sheeter/verify/verifygo/target/proto/codeGo"
)

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyProto(rootPath string) {
	reader := sheeterProto.VerifyData1Reader{}

	if err := reader.FromPathHalf(filepath.Join(rootPath, "target", internal.PathProto, internal.PathData)); err != nil {
		panic(fmt.Errorf("verify proto: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertProto(ok)
	assertProto(*actual.Key == 1)
	assertProto(*actual.Name == "名稱1")
	assertProto(*actual.Enable == true)
	assertProto(*actual.Reward.Atium == 2)
	assertProto(*actual.Reward.Crystal == 120)
	assertProto(*actual.Reward.Diamond == 10)
	assertProto(*actual.Reward.FelIron == 6)
	assertProto(*actual.Reward.Gold == 500)
	assertProto(len(actual.Reward.Item) == 3)
	assertProto(*actual.Reward.Item[0].ItemID == 10001)
	assertProto(*actual.Reward.Item[0].Type == 1)
	assertProto(*actual.Reward.Item[0].Count == 10)
	assertProto(*actual.Reward.Item[1].ItemID == 0)
	assertProto(*actual.Reward.Item[1].Type == 0)
	assertProto(*actual.Reward.Item[1].Count == 0)
	assertProto(*actual.Reward.Item[2].ItemID == 0)
	assertProto(*actual.Reward.Item[2].Type == 0)
	assertProto(*actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[2]
	assertJson(ok)
	assertProto(*actual.Key == 2)
	assertProto(*actual.Name == "名稱2")
	assertProto(*actual.Enable == true)
	assertProto(*actual.Reward.Atium == 2)
	assertProto(*actual.Reward.Crystal == 135)
	assertProto(*actual.Reward.Diamond == 12)
	assertProto(*actual.Reward.FelIron == 8)
	assertProto(*actual.Reward.Gold == 550)
	assertProto(len(actual.Reward.Item) == 3)
	assertProto(*actual.Reward.Item[0].ItemID == 10001)
	assertProto(*actual.Reward.Item[0].Type == 1)
	assertProto(*actual.Reward.Item[0].Count == 10)
	assertProto(*actual.Reward.Item[1].ItemID == 10002)
	assertProto(*actual.Reward.Item[1].Type == 1)
	assertProto(*actual.Reward.Item[1].Count == 5)
	assertProto(*actual.Reward.Item[2].ItemID == 0)
	assertProto(*actual.Reward.Item[2].Type == 0)
	assertProto(*actual.Reward.Item[2].Count == 0)

	actual, ok = reader.Datas[3]
	assertJson(ok)
	assertProto(*actual.Key == 3)
	assertProto(*actual.Name == "名稱3")
	assertProto(*actual.Enable == false)
	assertProto(*actual.Reward.Atium == 3)
	assertProto(*actual.Reward.Crystal == 150)
	assertProto(*actual.Reward.Diamond == 14)
	assertProto(*actual.Reward.FelIron == 10)
	assertProto(*actual.Reward.Gold == 600)
	assertJson(len(actual.Reward.Item) == 3)
	assertProto(*actual.Reward.Item[0].ItemID == 10001)
	assertProto(*actual.Reward.Item[0].Type == 1)
	assertProto(*actual.Reward.Item[0].Count == 10)
	assertProto(*actual.Reward.Item[1].ItemID == 10002)
	assertProto(*actual.Reward.Item[1].Type == 1)
	assertProto(*actual.Reward.Item[1].Count == 5)
	assertProto(*actual.Reward.Item[2].ItemID == 10003)
	assertProto(*actual.Reward.Item[2].Type == 1)
	assertProto(*actual.Reward.Item[2].Count == 2)

	actual, ok = reader.Datas[4]
	assertJson(ok)
	assertProto(*actual.Key == 4)
	assertProto(*actual.Name == "名稱4")
	assertProto(*actual.Enable == false)
	assertProto(*actual.Reward.Atium == 3)
	assertProto(*actual.Reward.Crystal == 165)
	assertProto(*actual.Reward.Diamond == 16)
	assertProto(*actual.Reward.FelIron == 12)
	assertProto(*actual.Reward.Gold == 650)
	assertJson(len(actual.Reward.Item) == 3)
	assertProto(*actual.Reward.Item[0].ItemID == 10001)
	assertProto(*actual.Reward.Item[0].Type == 1)
	assertProto(*actual.Reward.Item[0].Count == 10)
	assertProto(*actual.Reward.Item[1].ItemID == 10002)
	assertProto(*actual.Reward.Item[1].Type == 1)
	assertProto(*actual.Reward.Item[1].Count == 5)
	assertProto(*actual.Reward.Item[2].ItemID == 10003)
	assertProto(*actual.Reward.Item[2].Type == 1)
	assertProto(*actual.Reward.Item[2].Count == 3)

	actual, ok = reader.Datas[5]
	assertJson(ok == false)

	fmt.Println("verify proto: success")
}

func assertProto(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify proto: verify failed"))
	} // if
}
