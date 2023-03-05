// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: exampleData.proto

package sheeterProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward *Reward `protobuf:"bytes,1,opt,name=Reward,proto3,oneof" json:"Reward,omitempty"`  //
	Wallet *Wallet `protobuf:"bytes,2,opt,name=Wallet,proto3,oneof" json:"Wallet,omitempty"`  //
	Ignore *int32  `protobuf:"varint,3,opt,name=Ignore,proto3,oneof" json:"Ignore,omitempty"` // 忽略
	Key    *int32  `protobuf:"varint,4,opt,name=Key,proto3,oneof" json:"Key,omitempty"`       // 索引
	Name   *string `protobuf:"bytes,5,opt,name=Name,proto3,oneof" json:"Name,omitempty"`      // 名稱
}

func (x *ExampleData) Reset() {
	*x = ExampleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampleData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleData) ProtoMessage() {}

func (x *ExampleData) ProtoReflect() protoreflect.Message {
	mi := &file_exampleData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleData.ProtoReflect.Descriptor instead.
func (*ExampleData) Descriptor() ([]byte, []int) {
	return file_exampleData_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleData) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *ExampleData) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

func (x *ExampleData) GetIgnore() int32 {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return 0
}

func (x *ExampleData) GetKey() int32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *ExampleData) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type ExampleDataStorer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[int32]*ExampleData `protobuf:"bytes,1,rep,name=Datas,proto3" json:"Datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleDataStorer) Reset() {
	*x = ExampleDataStorer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampleData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleDataStorer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleDataStorer) ProtoMessage() {}

func (x *ExampleDataStorer) ProtoReflect() protoreflect.Message {
	mi := &file_exampleData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleDataStorer.ProtoReflect.Descriptor instead.
func (*ExampleDataStorer) Descriptor() ([]byte, []int) {
	return file_exampleData_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleDataStorer) GetDatas() map[int32]*ExampleData {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_exampleData_proto protoreflect.FileDescriptor

var file_exampleData_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01,
	0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a,
	0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x31, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x48, 0x01, 0x52, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52,
	0x03, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x4b, 0x65, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x44, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x53, 0x0a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exampleData_proto_rawDescOnce sync.Once
	file_exampleData_proto_rawDescData = file_exampleData_proto_rawDesc
)

func file_exampleData_proto_rawDescGZIP() []byte {
	file_exampleData_proto_rawDescOnce.Do(func() {
		file_exampleData_proto_rawDescData = protoimpl.X.CompressGZIP(file_exampleData_proto_rawDescData)
	})
	return file_exampleData_proto_rawDescData
}

var file_exampleData_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_exampleData_proto_goTypes = []interface{}{
	(*ExampleData)(nil),       // 0: sheeterProto.ExampleData
	(*ExampleDataStorer)(nil), // 1: sheeterProto.ExampleDataStorer
	nil,                       // 2: sheeterProto.ExampleDataStorer.DatasEntry
	(*Reward)(nil),            // 3: sheeterProto.Reward
	(*Wallet)(nil),            // 4: sheeterProto.Wallet
}
var file_exampleData_proto_depIdxs = []int32{
	3, // 0: sheeterProto.ExampleData.Reward:type_name -> sheeterProto.Reward
	4, // 1: sheeterProto.ExampleData.Wallet:type_name -> sheeterProto.Wallet
	2, // 2: sheeterProto.ExampleDataStorer.Datas:type_name -> sheeterProto.ExampleDataStorer.DatasEntry
	0, // 3: sheeterProto.ExampleDataStorer.DatasEntry.value:type_name -> sheeterProto.ExampleData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_exampleData_proto_init() }
func file_exampleData_proto_init() {
	if File_exampleData_proto != nil {
		return
	}
	file_reward_proto_init()
	file_wallet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exampleData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exampleData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleDataStorer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_exampleData_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exampleData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exampleData_proto_goTypes,
		DependencyIndexes: file_exampleData_proto_depIdxs,
		MessageInfos:      file_exampleData_proto_msgTypes,
	}.Build()
	File_exampleData_proto = out.File
	file_exampleData_proto_rawDesc = nil
	file_exampleData_proto_goTypes = nil
	file_exampleData_proto_depIdxs = nil
}
