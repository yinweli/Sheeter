package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ShellRun 執行shell命令
func ShellRun(name string, args ...string) error {
	buffer := &bytes.Buffer{}
	cmd := exec.Command(name, args...)
	cmd.Stderr = buffer

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("shell run failed: %w", err)
	} // if

	return nil
}

// ShellExist shell命令是否存在
func ShellExist(name string) bool {
	if _, err := exec.LookPath(name); err != nil {
		return false
	} // if

	return true
}
