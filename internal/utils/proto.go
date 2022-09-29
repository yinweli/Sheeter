package utils

import (
	"fmt"

	//nolint:staticcheck // 由於使用的proto反射庫沒有升級到proto.V2 所以必須使用舊的方式來轉換
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

// JsonToProto json轉為proto資料
func JsonToProto(filename, message string, importPaths []string, json []byte) (result []byte, err error) {
	pd, err := parseProto(filename, importPaths)

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
		return nil, fmt.Errorf("json to proto failed: %w", err)
	} // if

	data, err := proto.Marshal(obj)

	if err != nil {
		return nil, fmt.Errorf("json to proto failed: %w", err)
	} // if

	return data, nil
}

// ProtoToJson proto轉為json資料
func ProtoToJson(filename, message string, importPaths []string, data []byte) (result []byte, err error) {
	pd, err := parseProto(filename, importPaths)

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

	if err := proto.Unmarshal(data, obj); err != nil {
		return nil, fmt.Errorf("proto to json failed: %w", err)
	} // if

	json, err := obj.MarshalJSON()

	if err != nil {
		return nil, fmt.Errorf("proto to json failed: %w", err)
	} // if

	return json, nil
}

// ParseProto 解析proto檔案, 獲得proto描述器
func parseProto(filename string, importPaths []string) (result *desc.FileDescriptor, err error) {
	parser := protoparse.Parser{ImportPaths: importPaths}
	pds, err := parser.ParseFiles(filename)

	if err != nil {
		return nil, fmt.Errorf("%s: parse proto failed: %w", filename, err)
	} // if

	if len(pds) == 0 {
		return nil, fmt.Errorf("%s: parse proto failed: descriptor empty", filename)
	} // if

	return pds[0], nil
}
