// generated by sheeter, DO NOT EDIT.

package sheeter

type VerifyData2 struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type VerifyData2Storer struct {
	Datas map[int64]VerifyData2
}
