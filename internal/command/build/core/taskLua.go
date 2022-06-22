package core

// runLua 輸出lua
func (this *Task) runLua() error {
	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
