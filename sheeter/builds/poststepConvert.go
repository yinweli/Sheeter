package builds

import (
	"fmt"
	"os"

	"github.com/yinweli/Sheeter/sheeter/utils"
)

// poststepConvert 後製轉換資料
type poststepConvert struct {
	include  string // 引入路徑
	outputCs string // 輸出cs路徑
	outputGo string // 輸出go路徑
	source   string // 來源檔案
}

// PoststepConvertCs 後製轉換cs
func PoststepConvertCs(material any, _ chan any) error {
	data, ok := material.(*poststepConvert)

	if ok == false {
		return nil
	} // if

	_ = os.MkdirAll(data.outputCs, os.ModePerm)
	include := fmt.Sprintf("--proto_path=./%s", data.include)
	output := fmt.Sprintf("--csharp_out=./%s", data.outputCs)
	source := fmt.Sprintf("./%s", data.source)

	if err := utils.ShellRun("protoc", "--experimental_allow_proto3_optional", include, output, source); err != nil {
		return fmt.Errorf("%s: poststep convert cs failed: %w", data.source, err)
	} // if

	return nil
}

// PoststepConvertGo 後製轉換go
func PoststepConvertGo(material any, _ chan any) error {
	data, ok := material.(*poststepConvert)

	if ok == false {
		return nil
	} // if

	_ = os.MkdirAll(data.outputGo, os.ModePerm)
	include := fmt.Sprintf("--proto_path=./%s", data.include)
	output := fmt.Sprintf("--go_out=./%s", data.outputGo)
	source := fmt.Sprintf("./%s", data.source)

	if err := utils.ShellRun("protoc", "--experimental_allow_proto3_optional", include, output, source); err != nil {
		return fmt.Errorf("%s: poststep convert go failed: %w", data.source, err)
	} // if

	return nil
}
