// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type File1Data struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 忽略
	Ignore int64 `json:"Ignore"`
	// 索引
	Key int64 `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type File1DataStorer struct {
	Datas map[int64]*File1Data
}
