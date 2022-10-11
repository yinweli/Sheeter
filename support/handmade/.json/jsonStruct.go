// 以下是模板驗證用程式碼

package _json

type Reward struct {
	//
	Item []Item `json:"Item"`
	// 天金
	Atium int64 `json:"Atium"`
	// 魔晶
	Crystal int64 `json:"Crystal"`
	// 獎勵說明
	Desc string `json:"Desc"`
	// 鑽石
	Diamond int64 `json:"Diamond"`
	// 精鐵
	FelIron int64 `json:"FelIron"`
	// 金幣
	Gold int64 `json:"Gold"`
}

type RewardStorer struct {
	Datas map[int64]*Reward
}

// 以下是為了通過編譯的程式碼, 不可使用

type Item struct {
	// 物品數量
	Count int64 `json:"Count"`
	// 物品編號
	ItemID int64 `json:"ItemID"`
	// 物品類型
	Type int64 `json:"Type"`
}
