package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
	sheeterProto "github.com/yinweli/Sheeter/verify/protodata/code"
)

func main() {
	name01 := int64(1)
	name02 := int64(2)
	name03 := int64(3)
	name11 := true
	name12 := false
	name13 := true
	name21 := int64(1)
	name22 := int64(2)
	name23 := int64(3)
	name31 := "a"
	name32 := "b"
	name33 := "c"

	storer := &sheeterProto.RealDataStorer{Datas: map[int64]*sheeterProto.RealData{}}
	storer.Datas[1] = &sheeterProto.RealData{
		Name0: &name01,
		S: &sheeterProto.S{
			Name1: &name11,
			A: []*sheeterProto.A{
				{Name2: &name21, Name3: &name31},
				{Name2: &name21, Name3: &name31},
				{Name2: &name21, Name3: &name31},
			},
		},
	}
	storer.Datas[2] = &sheeterProto.RealData{
		Name0: &name02,
		S: &sheeterProto.S{
			Name1: &name12,
			A: []*sheeterProto.A{
				{Name2: &name22, Name3: &name32},
				{Name2: &name22, Name3: &name32},
				{Name2: &name22, Name3: &name32},
			},
		},
	}
	storer.Datas[3] = &sheeterProto.RealData{
		Name0: &name03,
		S: &sheeterProto.S{
			Name1: &name13,
			A: []*sheeterProto.A{
				{Name2: &name23, Name3: &name33},
				{Name2: &name23, Name3: &name33},
				{Name2: &name23, Name3: &name33},
			},
		},
	}

	data, err := proto.Marshal(storer)

	if err != nil {
		fmt.Println(err)
	} // if

	if err := utils.WriteFile("real."+internal.ExtProtoData, data); err != nil {
		fmt.Println(err)
	} // if

	fmt.Println("build testdata success")
}
