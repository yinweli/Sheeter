package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecker_AllSuccess(t *testing.T) {
	checker := NewChecker()
	checker.Add(true, message1)
	checker.Add(true, message2)
	checker.Add(true, message3)
	errors := checker.Errors()

	assert.Equal(t, true, checker.Result(), "result failed")
	assert.Equal(t, 0, len(errors), "errors len failed")
}

func TestChecker_AllFailed(t *testing.T) {
	checker := NewChecker()
	checker.Add(false, message1)
	checker.Add(false, message2)
	checker.Add(false, message3)
	errors := checker.Errors()

	assert.Equal(t, false, checker.Result(), "result failed")
	assert.Equal(t, 3, len(errors), "errors len failed")
	assert.Equal(t, message1, errors[0], "errors[0] failed")
	assert.Equal(t, message2, errors[1], "errors[1] failed")
	assert.Equal(t, message3, errors[2], "errors[2] failed")
}

func TestChecker_Mix(t *testing.T) {
	checker := NewChecker()
	checker.Add(false, message1)
	checker.Add(true, message2)
	checker.Add(false, message3)
	errors := checker.Errors()

	assert.Equal(t, false, checker.Result(), "result failed")
	assert.Equal(t, 2, len(errors), "errors len failed")
	assert.Equal(t, message1, errors[0], "errors[0] failed")
	assert.Equal(t, message3, errors[1], "errors[1] failed")
}

func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	assert.NotNil(t, checker, "checker nil")
}

var message1 string = "message1" // 測試訊息1
var message2 string = "message2" // 測試訊息2
var message3 string = "message3" // 測試訊息3
