package utils

import (
	"fmt"

	protos "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

// JsonToProto json轉為proto資料
func JsonToProto(proto, message string, json []byte) (results []byte, err error) {
	pd, err := parseProto(proto)

	if err != nil {
		return nil, fmt.Errorf("json to proto failed: %w", err)
	} // if

	md := pd.FindMessage(message)

	if md == nil {
		return nil, fmt.Errorf("json to proto failed: message not found")
	} // if

	obj := dynamic.NewMessage(md)

	if obj == nil {
		return nil, fmt.Errorf("json to proto failed: new obj failed")
	} // if

	if err := obj.UnmarshalJSON(json); err != nil {
		return nil, fmt.Errorf("json to proto failed: json to storer failed")
	} // if

	data, err := ptypes.MarshalAny(obj) // 由於使用的proto反射庫沒有升級到proto.V2, 所以必須使用舊的方式來轉換

	if err != nil {
		return nil, fmt.Errorf("json to proto failed: %w", err)
	} // if

	return data.Value, nil
}

// ProtoToJson proto轉為json資料
func ProtoToJson(proto, message string, data []byte) (results []byte, err error) {
	pd, err := parseProto(proto)

	if err != nil {
		return nil, fmt.Errorf("proto to json failed: %w", err)
	} // if

	md := pd.FindMessage(message)

	if md == nil {
		return nil, fmt.Errorf("proto to json failed: message not found")
	} // if

	obj := dynamic.NewMessage(md)

	if obj == nil {
		return nil, fmt.Errorf("proto to json failed: new obj failed")
	} // if

	if err := protos.Unmarshal(data, obj); err != nil {
		return nil, fmt.Errorf("proto to json failed: %w", err)
	} // if

	json, err := obj.MarshalJSON()

	if err != nil {
		return nil, fmt.Errorf("proto to json failed: %w", err)
	} // if

	return json, nil
}

// ParseProto 解析proto檔案, 獲得proto描述器
func parseProto(path string) (proto *desc.FileDescriptor, err error) {
	parser := protoparse.Parser{}
	pds, err := parser.ParseFiles(path)

	if err != nil {
		return nil, fmt.Errorf("%s: parse proto failed: %w", path, err)
	} // if

	if len(pds) == 0 {
		return nil, fmt.Errorf("%s: parse proto failed: descriptor empty", path)
	} // if

	return pds[0], nil
}
