// generated by sheeter, DO NOT EDIT.

package sheeterJson

type Benchmark09Data struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type Benchmark09DataStorer struct {
	Datas map[int64]Benchmark09Data
}
