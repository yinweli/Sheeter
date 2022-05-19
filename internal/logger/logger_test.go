package logger

import (
	"os"
	"testing"

	"Sheeter/internal"
	"Sheeter/test"
)

func TestInitializeFinalize(t *testing.T) {
	err := Initialize()

	if err != nil {
		t.Error(err)
	} // if

	err = Finalize()

	if err != nil {
		t.Error(err)
	} // if

	err = os.Remove(internal.LoggerFileName)

	if err != nil {
		t.Error(err)
	} // if
}

func TestInfo(t *testing.T) {
	alterLog := test.NewAlterLog()
	defer alterLog.Finalize()

	Info("test message")

	if len(alterLog.String()) <= 0 {
		t.Error("info message should not empty")
	} // if
}

func TestError(t *testing.T) {
	alterLog := test.NewAlterLog()
	defer alterLog.Finalize()

	Error("test message")

	if len(alterLog.String()) <= 0 {
		t.Error("error message should not empty")
	} // if
}
