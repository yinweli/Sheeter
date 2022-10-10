// 以下是模板驗證用程式碼

package _proto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect" // 這為了通過編譯的程式碼, 不可使用
)

type RewardReader struct {
	*RewardStorer
}

func (this *RewardReader) DataName() string {
	return "reward"
}

func (this *RewardReader) DataExt() string {
	return "bytes"
}

func (this *RewardReader) DataFile() string {
	return "reward.bytes"
}

func (this *RewardReader) FromData(data []byte) error {
	this.RewardStorer = &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := proto.Unmarshal(data, this.RewardStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *RewardReader) MergeData(data []byte) error {
	tmpl := &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.RewardStorer == nil {
		this.RewardStorer = &RewardStorer{
			Datas: map[int64]*Reward{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.RewardStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.RewardStorer.Datas[k] = v
	}

	return nil
}

// 以下是為了通過編譯的程式碼, 不可使用

type RewardStorer struct {
	Datas map[int64]*Reward
}

func (x *RewardStorer) Reset() {
}

func (x *RewardStorer) String() string {
	return ""
}

func (*RewardStorer) ProtoMessage() {}

func (x *RewardStorer) ProtoReflect() protoreflect.Message {
	return nil
}

func (*RewardStorer) Descriptor() ([]byte, []int) {
	return nil, nil
}

func (x *RewardStorer) GetDatas() map[int64]*Reward {
	return nil
}

type Reward struct {
}
