package verifydata

type Struct struct {
	Reward Reward `json:"Reward"`
	Enable bool   `json:"enable"`
	Key    int64  `json:"key"`
	Name   string `json:"name"`
}

type Reward struct {
	Item    []Item `json:"Item"`
	Atium   int64  `json:"atium"`
	Crystal int64  `json:"crystal"`
	Diamond int64  `json:"diamond"`
	FelIron int64  `json:"felIron"`
	Gold    int64  `json:"gold"`
}

type Item struct {
	Count  int64 `json:"count"`
	ItemID int64 `json:"itemID"`
	Type   int64 `json:"type"`
}
