package main

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	sheeterProto "github.com/yinweli/Sheeter/verify/verifygo/target/proto/codeGo"
)

func verifyProto(rootPath string) {
	path := filepath.Join(rootPath, "target", internal.PathProto, internal.PathData)
	verifyProto1(path)
	verifyProto2(path)
}

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyProto1(path string) {
	reader := sheeterProto.VerifyData1Reader{}

	if err := reader.FromPathHalf(path); err != nil {
		panic(fmt.Errorf("verify proto: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertProto(ok)
	assertProto(actual.GetKey() == 1)
	assertProto(actual.GetEnable() == true)
	assertProto(actual.GetName() == "名稱1")
	assertProto(actual.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual.GetReward().GetGold() == 100)
	assertProto(actual.GetReward().GetDiamond() == 10)
	assertProto(actual.GetReward().GetCrystal() == 199)
	assertProto(actual.GetReward().GetFelIron() == 5)
	assertProto(actual.GetReward().GetAtium() == 1)
	assertProto(len(actual.GetReward().GetItem()) == 3)
	assertProto(actual.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual.GetReward().GetItem()[2].GetCount() == 0)

	actual, ok = reader.Datas[2]
	assertProto(ok)
	assertProto(actual.GetKey() == 2)
	assertProto(actual.GetEnable() == false)
	assertProto(actual.GetName() == "名稱2")
	assertProto(actual.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual.GetReward().GetGold() == 200)
	assertProto(actual.GetReward().GetDiamond() == 20)
	assertProto(actual.GetReward().GetCrystal() == 299)
	assertProto(actual.GetReward().GetFelIron() == 10)
	assertProto(actual.GetReward().GetAtium() == 2)
	assertProto(len(actual.GetReward().GetItem()) == 3)
	assertProto(actual.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual.GetReward().GetItem()[2].GetCount() == 3)

	actual, ok = reader.Datas[3]
	assertProto(ok == false)

	fmt.Println("verify proto: success")
}

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyProto2(path string) {
	reader := sheeterProto.VerifyData2Reader{}

	if err := reader.FromPathHalf(path); err != nil {
		panic(fmt.Errorf("verify proto: %w", err))
	} // if

	actual, ok := reader.Datas[1]
	assertProto(ok)
	assertProto(actual.GetKey() == 1)
	assertProto(actual.GetEnable() == true)
	assertProto(actual.GetName() == "名稱1")
	assertProto(actual.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual.GetReward().GetGold() == 100)
	assertProto(actual.GetReward().GetDiamond() == 10)
	assertProto(actual.GetReward().GetCrystal() == 0)
	assertProto(actual.GetReward().GetFelIron() == 0)
	assertProto(actual.GetReward().GetAtium() == 0)
	assertProto(len(actual.GetReward().GetItem()) == 3)
	assertProto(actual.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual.GetReward().GetItem()[2].GetCount() == 0)

	actual, ok = reader.Datas[2]
	assertProto(ok)
	assertProto(actual.GetKey() == 2)
	assertProto(actual.GetEnable() == false)
	assertProto(actual.GetName() == "名稱2")
	assertProto(actual.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual.GetReward().GetGold() == 200)
	assertProto(actual.GetReward().GetDiamond() == 20)
	assertProto(actual.GetReward().GetCrystal() == 0)
	assertProto(actual.GetReward().GetFelIron() == 0)
	assertProto(actual.GetReward().GetAtium() == 0)
	assertProto(len(actual.GetReward().GetItem()) == 3)
	assertProto(actual.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual.GetReward().GetItem()[2].GetCount() == 3)

	actual, ok = reader.Datas[3]
	assertProto(ok == false)

	fmt.Println("verify proto: success")
}

func assertProto(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify proto: verify failed"))
	} // if
}
