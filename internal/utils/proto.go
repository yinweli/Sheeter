package utils

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

// ParseProto 解析proto檔案, 獲得proto描述器
func ParseProto(path string) (proto *desc.FileDescriptor, err error) {
	parser := protoparse.Parser{}
	protos, err := parser.ParseFiles(path)

	if err != nil {
		return nil, fmt.Errorf("%s: parse proto failed: %w", path, err)
	} // if

	if len(protos) == 0 {
		return nil, fmt.Errorf("%s: parse proto failed, empty proto", path)
	} // if

	return protos[0], nil
}
