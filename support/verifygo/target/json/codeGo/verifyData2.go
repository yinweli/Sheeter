// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type VerifyData2 struct {
	//
	Reward Reward `json:"Reward"`
	// 是否啟用
	Enable bool `json:"Enable"`
	// 隱藏
	Hide bool `json:"Hide"`
	// 索引
	Key string `json:"Key"`
	// 名稱
	Name string `json:"Name"`
}

type VerifyData2Storer struct {
	Datas map[string]*VerifyData2
}
