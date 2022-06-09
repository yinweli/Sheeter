package util

import "testing"

func TestSilentClose(t *testing.T) {
	SilentClose(&mockCloser{})
	SilentClose(nil)
}

type mockCloser struct {
}

func (this *mockCloser) Close() error {
	return nil
}
