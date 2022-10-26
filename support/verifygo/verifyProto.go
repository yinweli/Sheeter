package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	sheeterProto "github.com/yinweli/Sheeter/support/verifygo/target/proto/codeGo"
)

func verifyProtoFrom(threads int) {
	loader := newProtoFileLoader()
	depot := sheeterProto.NewDepot(loader)

	assertProto(depot.FromData())
	verifyProto(depot, threads)

	fmt.Println("verify proto from: success")
}

func verifyProtoMerge(threads int) {
	loader := newProtoFileLoader()
	depot := sheeterProto.NewDepot(loader)

	assertProto(depot.MergeData())
	verifyProto(depot, threads)

	fmt.Println("verify proto merge: success")
}

func verifyProto(depot *sheeterProto.Depot, threads int) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			for i := int64(1); i <= 100; i++ {
				actual1, ok := depot.VerifyData1.Get(i)
				assertJson(ok)
				assertJson(actual1.GetKey() == i)
				assertJson(actual1.GetHide() == false)
				assertJson(actual1.GetEnable() == (i%2 == 1))
				assertJson(actual1.GetName() == fmt.Sprintf("名稱%d", i))
				assertJson(actual1.GetReward().GetDesc() == fmt.Sprintf("獎勵說明%d", i))
				assertJson(actual1.GetReward().GetGold() == i*2)
				assertJson(actual1.GetReward().GetDiamond() == i*3)
				assertJson(actual1.GetReward().GetCrystal() == i*4)
				assertJson(actual1.GetReward().GetFelIron() == i*5)
				assertJson(actual1.GetReward().GetAtium() == i*6)
				assertJson(len(actual1.GetReward().GetItem()) == 3)
				assertJson(actual1.GetReward().GetItem()[0].GetItemID() == 1000+i)
				assertJson(actual1.GetReward().GetItem()[0].GetType() == 0)
				assertJson(actual1.GetReward().GetItem()[0].GetCount() == i)
				assertJson(actual1.GetReward().GetItem()[1].GetItemID() == 10000+i)
				assertJson(actual1.GetReward().GetItem()[1].GetType() == 1)
				assertJson(actual1.GetReward().GetItem()[1].GetCount() == i)
				assertJson(actual1.GetReward().GetItem()[2].GetItemID() == 100000+i)
				assertJson(actual1.GetReward().GetItem()[2].GetType() == 2)
				assertJson(actual1.GetReward().GetItem()[2].GetCount() == i)

				actual2, ok := depot.VerifyData2.Get(i)
				assertJson(ok)
				assertJson(actual2.GetKey() == i)
				assertJson(actual2.GetHide() == false)
				assertJson(actual2.GetEnable() == (i%2 == 1))
				assertJson(actual2.GetName() == fmt.Sprintf("名稱%d", i))
				assertJson(actual2.GetReward().GetDesc() == fmt.Sprintf("獎勵說明%d", i))
				assertJson(actual2.GetReward().GetGold() == i*2)
				assertJson(actual2.GetReward().GetDiamond() == i*3)
				assertJson(actual2.GetReward().GetCrystal() == 0)
				assertJson(actual2.GetReward().GetFelIron() == 0)
				assertJson(actual2.GetReward().GetAtium() == 0)
				assertJson(len(actual2.GetReward().GetItem()) == 3)
				assertJson(actual2.GetReward().GetItem()[0].GetItemID() == 1000+i)
				assertJson(actual2.GetReward().GetItem()[0].GetType() == 0)
				assertJson(actual2.GetReward().GetItem()[0].GetCount() == i)
				assertJson(actual2.GetReward().GetItem()[1].GetItemID() == 10000+i)
				assertJson(actual2.GetReward().GetItem()[1].GetType() == 1)
				assertJson(actual2.GetReward().GetItem()[1].GetCount() == i)
				assertJson(actual2.GetReward().GetItem()[2].GetItemID() == 100000+i)
				assertJson(actual2.GetReward().GetItem()[2].GetType() == 2)
				assertJson(actual2.GetReward().GetItem()[2].GetCount() == i)
			} // for

			_, ok := depot.VerifyData1.Get(101)
			assertJson(ok == false)

			_, ok = depot.VerifyData2.Get(101)
			assertJson(ok == false)

			waitGroup.Done()
		}()
	} // for

	waitGroup.Wait()
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
