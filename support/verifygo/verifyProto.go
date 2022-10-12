package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	sheeterProto "github.com/yinweli/Sheeter/support/verifygo/target/proto/codeGo"
)

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyProtoFrom() {
	loader := newProtoFileLoader()
	depot := sheeterProto.NewDepot(loader)

	assertProto(depot.FromData())

	actual1, ok := depot.VerifyData1.Get(1)
	assertProto(ok)
	assertProto(actual1.GetKey() == 1)
	assertProto(actual1.GetHide() == false)
	assertProto(actual1.GetEnable() == true)
	assertProto(actual1.GetName() == "名稱1")
	assertProto(actual1.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual1.GetReward().GetGold() == 100)
	assertProto(actual1.GetReward().GetDiamond() == 10)
	assertProto(actual1.GetReward().GetCrystal() == 199)
	assertProto(actual1.GetReward().GetFelIron() == 5)
	assertProto(actual1.GetReward().GetAtium() == 1)
	assertProto(len(actual1.GetReward().GetItem()) == 3)
	assertProto(actual1.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual1.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual1.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual1.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual1.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual1.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual1.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual1.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual1.GetReward().GetItem()[2].GetCount() == 0)

	actual1, ok = depot.VerifyData1.Get(2)
	assertProto(ok)
	assertProto(actual1.GetKey() == 2)
	assertProto(actual1.GetHide() == false)
	assertProto(actual1.GetEnable() == false)
	assertProto(actual1.GetName() == "名稱2")
	assertProto(actual1.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual1.GetReward().GetGold() == 200)
	assertProto(actual1.GetReward().GetDiamond() == 20)
	assertProto(actual1.GetReward().GetCrystal() == 299)
	assertProto(actual1.GetReward().GetFelIron() == 10)
	assertProto(actual1.GetReward().GetAtium() == 2)
	assertProto(len(actual1.GetReward().GetItem()) == 3)
	assertProto(actual1.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual1.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual1.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual1.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual1.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual1.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual1.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual1.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual1.GetReward().GetItem()[2].GetCount() == 3)

	actual1, ok = depot.VerifyData1.Get(3)
	assertProto(ok == false)

	actual2, ok := depot.VerifyData2.Get(1)
	assertProto(ok)
	assertProto(actual2.GetKey() == 1)
	assertProto(actual2.GetHide() == false)
	assertProto(actual2.GetEnable() == true)
	assertProto(actual2.GetName() == "名稱1")
	assertProto(actual2.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual2.GetReward().GetGold() == 100)
	assertProto(actual2.GetReward().GetDiamond() == 10)
	assertProto(actual2.GetReward().GetCrystal() == 0)
	assertProto(actual2.GetReward().GetFelIron() == 0)
	assertProto(actual2.GetReward().GetAtium() == 0)
	assertProto(len(actual2.GetReward().GetItem()) == 3)
	assertProto(actual2.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual2.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual2.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual2.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual2.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual2.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual2.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual2.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual2.GetReward().GetItem()[2].GetCount() == 0)

	actual2, ok = depot.VerifyData2.Get(2)
	assertProto(ok)
	assertProto(actual2.GetKey() == 2)
	assertProto(actual2.GetHide() == false)
	assertProto(actual2.GetEnable() == false)
	assertProto(actual2.GetName() == "名稱2")
	assertProto(actual2.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual2.GetReward().GetGold() == 200)
	assertProto(actual2.GetReward().GetDiamond() == 20)
	assertProto(actual2.GetReward().GetCrystal() == 0)
	assertProto(actual2.GetReward().GetFelIron() == 0)
	assertProto(actual2.GetReward().GetAtium() == 0)
	assertProto(len(actual2.GetReward().GetItem()) == 3)
	assertProto(actual2.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual2.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual2.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual2.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual2.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual2.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual2.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual2.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual2.GetReward().GetItem()[2].GetCount() == 3)

	actual2, ok = depot.VerifyData2.Get(3)
	assertProto(ok == false)

	fmt.Println("verify proto from: success")
}

//nolint // 太多魔術數字了, 所以只好略過lint
func verifyProtoMerge() {
	loader := newProtoFileLoader()
	depot := sheeterProto.NewDepot(loader)

	assertProto(depot.MergeData())

	actual1, ok := depot.VerifyData1.Get(1)
	assertProto(ok)
	assertProto(actual1.GetKey() == 1)
	assertProto(actual1.GetHide() == false)
	assertProto(actual1.GetEnable() == true)
	assertProto(actual1.GetName() == "名稱1")
	assertProto(actual1.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual1.GetReward().GetGold() == 100)
	assertProto(actual1.GetReward().GetDiamond() == 10)
	assertProto(actual1.GetReward().GetCrystal() == 199)
	assertProto(actual1.GetReward().GetFelIron() == 5)
	assertProto(actual1.GetReward().GetAtium() == 1)
	assertProto(len(actual1.GetReward().GetItem()) == 3)
	assertProto(actual1.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual1.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual1.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual1.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual1.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual1.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual1.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual1.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual1.GetReward().GetItem()[2].GetCount() == 0)

	actual1, ok = depot.VerifyData1.Get(2)
	assertProto(ok)
	assertProto(actual1.GetKey() == 2)
	assertProto(actual1.GetHide() == false)
	assertProto(actual1.GetEnable() == false)
	assertProto(actual1.GetName() == "名稱2")
	assertProto(actual1.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual1.GetReward().GetGold() == 200)
	assertProto(actual1.GetReward().GetDiamond() == 20)
	assertProto(actual1.GetReward().GetCrystal() == 299)
	assertProto(actual1.GetReward().GetFelIron() == 10)
	assertProto(actual1.GetReward().GetAtium() == 2)
	assertProto(len(actual1.GetReward().GetItem()) == 3)
	assertProto(actual1.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual1.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual1.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual1.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual1.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual1.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual1.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual1.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual1.GetReward().GetItem()[2].GetCount() == 3)

	actual1, ok = depot.VerifyData1.Get(3)
	assertProto(ok == false)

	actual2, ok := depot.VerifyData2.Get(1)
	assertProto(ok)
	assertProto(actual2.GetKey() == 1)
	assertProto(actual2.GetHide() == false)
	assertProto(actual2.GetEnable() == true)
	assertProto(actual2.GetName() == "名稱1")
	assertProto(actual2.GetReward().GetDesc() == "獎勵說明1")
	assertProto(actual2.GetReward().GetGold() == 100)
	assertProto(actual2.GetReward().GetDiamond() == 10)
	assertProto(actual2.GetReward().GetCrystal() == 0)
	assertProto(actual2.GetReward().GetFelIron() == 0)
	assertProto(actual2.GetReward().GetAtium() == 0)
	assertProto(len(actual2.GetReward().GetItem()) == 3)
	assertProto(actual2.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual2.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual2.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual2.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual2.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual2.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual2.GetReward().GetItem()[2].GetItemID() == 0)
	assertProto(actual2.GetReward().GetItem()[2].GetType() == 0)
	assertProto(actual2.GetReward().GetItem()[2].GetCount() == 0)

	actual2, ok = depot.VerifyData2.Get(2)
	assertProto(ok)
	assertProto(actual2.GetKey() == 2)
	assertProto(actual2.GetHide() == false)
	assertProto(actual2.GetEnable() == false)
	assertProto(actual2.GetName() == "名稱2")
	assertProto(actual2.GetReward().GetDesc() == "獎勵說明2")
	assertProto(actual2.GetReward().GetGold() == 200)
	assertProto(actual2.GetReward().GetDiamond() == 20)
	assertProto(actual2.GetReward().GetCrystal() == 0)
	assertProto(actual2.GetReward().GetFelIron() == 0)
	assertProto(actual2.GetReward().GetAtium() == 0)
	assertProto(len(actual2.GetReward().GetItem()) == 3)
	assertProto(actual2.GetReward().GetItem()[0].GetItemID() == 10001)
	assertProto(actual2.GetReward().GetItem()[0].GetType() == 1)
	assertProto(actual2.GetReward().GetItem()[0].GetCount() == 1)
	assertProto(actual2.GetReward().GetItem()[1].GetItemID() == 10002)
	assertProto(actual2.GetReward().GetItem()[1].GetType() == 2)
	assertProto(actual2.GetReward().GetItem()[1].GetCount() == 2)
	assertProto(actual2.GetReward().GetItem()[2].GetItemID() == 10003)
	assertProto(actual2.GetReward().GetItem()[2].GetType() == 3)
	assertProto(actual2.GetReward().GetItem()[2].GetCount() == 3)

	actual2, ok = depot.VerifyData2.Get(3)
	assertProto(ok == false)

	fmt.Println("verify proto merge: success")
}

type protoFileLoader struct {
	path string
}

func newProtoFileLoader() *protoFileLoader {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("root path failed"))
	} // if

	return &protoFileLoader{
		path: filepath.Join(filepath.Dir(root), "target", "proto", "data"),
	}
}

func (this *protoFileLoader) Error(name string, err error) {
	panic(fmt.Errorf("%s: proto file load failed: %w", name, err))
}

func (this *protoFileLoader) Load(name, ext, fullname string) []byte {
	path := filepath.Join(this.path, fullname)
	data, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("%s: proto file load failed: %w", name, err))
	}

	return data
}

func assertProto(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify proto: verify failed"))
	} // if
}
