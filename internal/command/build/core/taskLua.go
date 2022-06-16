package core

// executeLua 輸出lua
func (this *Task) executeLua() error {
	if this.bar != nil {
		this.bar.IncrBy(taskProgressL)
	} // if

	return nil
}
