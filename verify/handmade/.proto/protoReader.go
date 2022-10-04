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
	RewardStorer
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
	if err := proto.Unmarshal(data, &this.RewardStorer); err != nil {
		return fmt.Errorf("RewardReader: from data failed: %w", err)
	}

	return nil
}

// 以下是為了通過編譯的程式碼, 不可使用

type RewardStorer struct {
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

func (x *RewardStorer) GetDatas() map[int64]*int {
	return nil
}
