package util

import (
	"testing"
)

func TestChecker_1(t *testing.T) {
	testMessage1 := "test message1"
	testMessage2 := "test message2"
	testMessage3 := "test message3"
	checker := NewChecker()
	checker.Add(true, testMessage1)
	checker.Add(true, testMessage2)
	checker.Add(true, testMessage3)

	if checker.Result() == false {
		t.Error("result == false")
	}

	errors := checker.Errors()

	if len(errors) > 0 {
		t.Error("errors len > 0")
	}
}

func TestChecker_2(t *testing.T) {
	testMessage1 := "test message1"
	testMessage2 := "test message2"
	testMessage3 := "test message3"
	checker := NewChecker()
	checker.Add(false, testMessage1)
	checker.Add(false, testMessage2)
	checker.Add(false, testMessage3)

	if checker.Result() != false {
		t.Error("result != false")
	}

	errors := checker.Errors()

	if len(errors) != 3 {
		t.Error("errors len != 3")
	}

	if errors[0] != testMessage1 {
		t.Error("errors[0] invalid")
	}

	if errors[1] != testMessage2 {
		t.Error("errors[1] invalid")
	}

	if errors[2] != testMessage3 {
		t.Error("errors[2] invalid")
	}
}

func TestChecker_3(t *testing.T) {
	testMessage1 := "test message1"
	testMessage2 := "test message2"
	testMessage3 := "test message3"
	checker := NewChecker()
	checker.Add(false, testMessage1)
	checker.Add(true, testMessage2)
	checker.Add(false, testMessage3)

	if checker.Result() != false {
		t.Error("result != false")
	}

	errors := checker.Errors()

	if len(errors) != 2 {
		t.Error("errors len != 2")
	}

	if errors[0] != testMessage1 {
		t.Error("errors[0] invalid")
	}

	if errors[1] != testMessage3 {
		t.Error("errors[1] invalid")
	}
}

func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	if checker == nil {
		t.Error("new checker failed")
	}
}
