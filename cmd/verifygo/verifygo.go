package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"

	verifydata "github.com/yinweli/Sheeter/cmd/verifygo/json-go"
)

func main() {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("verify go: get root path failed"))
	} // if

	reader, err := verifydata.FromJsonFile(filepath.Join(filepath.Dir(root), verifydata.Json))

	if err != nil {
		panic(fmt.Errorf("verify go: %w", err))
	} // if

	expects := []verifydata.VerifyData{
		{
			Reward: verifydata.Reward{
				Item: []verifydata.Item{
					{Count: 10, ItemID: 10001, Type: 1},
					{Count: 0, ItemID: 0, Type: 0},
					{Count: 0, ItemID: 0, Type: 0},
				},
				Atium:   2,
				Crystal: 120,
				Diamond: 10,
				FelIron: 6,
				Gold:    500,
			},
			Enable: true,
			Key:    1,
			Name:   "名稱1",
		},
		{
			Reward: verifydata.Reward{
				Item: []verifydata.Item{
					{Count: 10, ItemID: 10001, Type: 1},
					{Count: 5, ItemID: 10002, Type: 1},
					{Count: 0, ItemID: 0, Type: 0},
				},
				Atium:   2,
				Crystal: 135,
				Diamond: 12,
				FelIron: 8,
				Gold:    550,
			},
			Enable: true,
			Key:    2,
			Name:   "名稱2",
		},
		{
			Reward: verifydata.Reward{
				Item: []verifydata.Item{
					{Count: 10, ItemID: 10001, Type: 1},
					{Count: 5, ItemID: 10002, Type: 1},
					{Count: 2, ItemID: 10003, Type: 1},
				},
				Atium:   3,
				Crystal: 150,
				Diamond: 14,
				FelIron: 10,
				Gold:    600,
			},
			Enable: false,
			Key:    3,
			Name:   "名稱3",
		},
		{
			Reward: verifydata.Reward{
				Item: []verifydata.Item{
					{Count: 10, ItemID: 10001, Type: 1},
					{Count: 5, ItemID: 10002, Type: 1},
					{Count: 3, ItemID: 10003, Type: 1},
				},
				Atium:   3,
				Crystal: 165,
				Diamond: 16,
				FelIron: 12,
				Gold:    650,
			},
			Enable: false,
			Key:    4,
			Name:   "名稱4",
		},
	}

	for _, itor := range expects {
		if actual, ok := reader[itor.Key]; ok == false || reflect.DeepEqual(actual, itor) == false {
			panic(fmt.Errorf("verify go: compare failed"))
		} // if
	} // for

	fmt.Println("verify go: success")
}
