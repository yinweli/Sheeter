// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type Reward struct {
	//
	Item []Item `json:"Item"`
	// 天金
	Atium int64 `json:"Atium"`
	// 魔晶
	Crystal int64 `json:"Crystal"`
	// 鑽石
	Diamond int64 `json:"Diamond"`
	// 精鐵
	FelIron int64 `json:"FelIron"`
	// 金幣
	Gold int64 `json:"Gold"`
}

type RewardStorer struct {
	Datas map[int64]Reward
}