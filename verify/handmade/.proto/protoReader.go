// 以下是模板驗證用程式碼

package _proto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RewardReader struct {
	*RewardStorer
}

func (this *RewardReader) FileName() string {
	return "reward.pbd"
}

func (this *RewardReader) FromPath(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("RewardReader: from path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *RewardReader) FromData(data []byte) error {
	this.RewardStorer = &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := proto.Unmarshal(data, this.RewardStorer); err != nil {
		return fmt.Errorf("RewardReader: from data failed: %w", err)
	}

	return nil
}

func (this *RewardReader) MergePath(path ...string) (repeats []int64) {
	for _, itor := range path {
		if data, err := os.ReadFile(filepath.Join(itor, this.FileName())); err == nil {
			repeats = append(repeats, this.MergeData(data)...)
		}
	}

	return repeats
}

func (this *RewardReader) MergeData(data []byte) (repeats []int64) {
	tmpl := &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return repeats
	}

	if this.RewardStorer == nil {
		this.RewardStorer = &RewardStorer{
			Datas: map[int64]*Reward{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.RewardStorer.Datas[k]; ok == false {
			this.RewardStorer.Datas[k] = v
		} else {
			repeats = append(repeats, k)
		}
	}

	return repeats
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
