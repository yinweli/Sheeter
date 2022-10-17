package utils

import (
	"fmt"
)

// Format 格式化程式碼
func Format() error {
	if err := ShellRun("dotnet", "csharpier", "."); err != nil {
		return fmt.Errorf("format cs code failed: %w", err)
	} // if

	if err := ShellRun("gofmt", "-w", "."); err != nil {
		return fmt.Errorf("format go code failed: %w", err)
	} // if

	if err := ShellRun("buf", "format", "-w", "."); err != nil {
		return fmt.Errorf("format proto schema failed: %w", err)
	} // if

	return nil
}
