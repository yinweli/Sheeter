package thirdparty

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// Check 檢查是否安裝第三方軟體
func Check() error {
	if util.ShellExist("go") == false {
		return fmt.Errorf("`go` not installed")
	} // if

	if util.ShellExist("quicktype") == false {
		return fmt.Errorf("`quicktype` not installed")
	} // if

	return nil
}
