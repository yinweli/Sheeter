package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, root, _, ok := runtime.Caller(0)

	if ok == false {
		panic(fmt.Errorf("verify go: get root path failed"))
	} // if

	rootPath := filepath.Dir(root)
	testJsonGo(rootPath)
}
