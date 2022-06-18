package core

// executeJsonCs 輸出json/cs
func (this *Task) executeJsonCs() error {
	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
