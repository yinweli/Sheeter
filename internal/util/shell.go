package util

import (
	"bytes"
	"os/exec"
)

// ShellRun 執行shell命令
func ShellRun(name string, args ...string) (err error, detail string) {
	buffer := &bytes.Buffer{}
	cmd := exec.Command(name, args...)
	cmd.Stderr = buffer

	return cmd.Run(), buffer.String()
}

// ShellExist shell命令是否存在
func ShellExist(name string) error {
	_, err := exec.LookPath(name)

	return err
}
