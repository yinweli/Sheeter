// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type Benchmark329Data struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type Benchmark329DataStorer struct {
	Datas map[int64]Benchmark329Data
}
