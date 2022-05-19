package util

import (
	"testing"

	"Sheeter/test"
)

func TestChecker_Add(t *testing.T) {
	alterLog := test.NewAlterLog()
	defer alterLog.Finalize()

	checker := NewChecker()
	checker.Add(true, "test message")

	if checker.Result() == false {
		t.Error("add true failed, result should be true")
	}

	if len(alterLog.String()) > 0 {
		t.Error("add true failed, message should be empty")
	}

	checker.Add(false, "test message")

	if checker.Result() == true {
		t.Error("add false failed, result should be false")
	}

	if len(alterLog.String()) <= 0 {
		t.Error("add false failed, message should not be empty")
	}
}

func TestChecker_Result(t *testing.T) {
	checker := NewChecker()
	checker.Add(true, "test message")

	if checker.Result() == false {
		t.Error("result should be true")
	}

	checker.Add(false, "test message")

	if checker.Result() == true {
		t.Error("result should be false")
	}
}

func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	if checker == nil {
		t.Error("new checker failed")
	}
}
