// generated by sheeter, DO NOT EDIT.

package sheeter

type Benchmark06Data struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type Benchmark06DataStorer struct {
	Datas map[int64]Benchmark06Data
}
