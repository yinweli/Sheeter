package core

import (
	"fmt"
	"os/exec"
)

// executeJsonGo 輸出json/go
func (this *Task) executeJsonGo() error {
	err := exec.Command("go", "fmt", this.jsonGoFilePath()).Run()

	if err != nil {
		return fmt.Errorf("format go failed: %s [%s]", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
