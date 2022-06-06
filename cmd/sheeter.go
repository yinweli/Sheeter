package main

import (
	"Sheeter/internal/command/build"

	"github.com/spf13/cobra"
)

const title = "sheeter" // 程式名稱
const version = "1.0.0" // 版本字串

func main() {
	rootCommand := cobra.Command{
		Use:     title,
		Long:    "Sheeter used to convert excel file to json file, and generate code of data structure",
		Version: version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}

// TODO: 浮點數轉換由strconv.ParseFloat改用strconv.FormatFloat
// TODO: 浮點數精度改由設定檔指定或是在欄位上指定
// TODO: Jobs機制放棄介面, 改用結構加上函式指標, 應該會比較單純
// TODO: 加上writeProto格式
// TODO: 事前檢查機制(可能要加到Jobs中), 要檢查是否有安裝Go/Protoc
// TODO: 讓BOM機制有效
