// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type Benchmark02DataF struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type Benchmark02DataFStorer struct {
	Datas map[int64]Benchmark02DataF
}
