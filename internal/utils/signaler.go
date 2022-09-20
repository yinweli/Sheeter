package utils

import (
	"sync"
)

// NewWaitGroup 建立waitGroup
func NewWaitGroup(delta int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(delta)
	return wg
}
