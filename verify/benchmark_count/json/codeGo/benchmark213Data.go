// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type Benchmark213Data struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type Benchmark213DataStorer struct {
	Datas map[int64]Benchmark213Data
}