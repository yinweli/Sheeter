// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: file2Enum.proto

package sheeterEnum

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

type File2Enum int32

const (
	File2Enum_Name0 File2Enum = 0 // 第0個列舉
	File2Enum_Name1 File2Enum = 1 // 第1個列舉
	File2Enum_Name2 File2Enum = 2 // 第2個列舉
)

// Enum value maps for File2Enum.
var (
	File2Enum_name = map[int32]string{
		0: "Name0",
		1: "Name1",
		2: "Name2",
	}
	File2Enum_value = map[string]int32{
		"Name0": 0,
		"Name1": 1,
		"Name2": 2,
	}
)

func (x File2Enum) Enum() *File2Enum {
	p := new(File2Enum)
	*p = x
	return p
}

func (x File2Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (File2Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_file2Enum_proto_enumTypes[0].Descriptor()
}

func (File2Enum) Type() protoreflect.EnumType {
	return &file_file2Enum_proto_enumTypes[0]
}

func (x File2Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use File2Enum.Descriptor instead.
func (File2Enum) EnumDescriptor() ([]byte, []int) {
	return file_file2Enum_proto_rawDescGZIP(), []int{0}
}

var File_file2Enum_proto protoreflect.FileDescriptor

var file_file2Enum_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2a, 0x2c,
	0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x32, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x4e,
	0x61, 0x6d, 0x65, 0x30, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x31, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x3b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file2Enum_proto_rawDescOnce sync.Once
	file_file2Enum_proto_rawDescData = file_file2Enum_proto_rawDesc
)

func file_file2Enum_proto_rawDescGZIP() []byte {
	file_file2Enum_proto_rawDescOnce.Do(func() {
		file_file2Enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_file2Enum_proto_rawDescData)
	})
	return file_file2Enum_proto_rawDescData
}

var file_file2Enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file2Enum_proto_goTypes = []interface{}{
	(File2Enum)(0), // 0: sheeterEnum.File2Enum
}
var file_file2Enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file2Enum_proto_init() }
func file_file2Enum_proto_init() {
	if File_file2Enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file2Enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file2Enum_proto_goTypes,
		DependencyIndexes: file_file2Enum_proto_depIdxs,
		EnumInfos:         file_file2Enum_proto_enumTypes,
	}.Build()
	File_file2Enum_proto = out.File
	file_file2Enum_proto_rawDesc = nil
	file_file2Enum_proto_goTypes = nil
	file_file2Enum_proto_depIdxs = nil
}
