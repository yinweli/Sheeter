package main

import (
	"fmt"

	sheeterEnum "github.com/yinweli/Sheeter/support/example/enum/codeGo"
)

func verifyEnum() {
	assertEnum(sheeterEnum.ExampleEnum_Name0 == 0)
	assertEnum(sheeterEnum.ExampleEnum_Name1 == 1)
	assertEnum(sheeterEnum.ExampleEnum_Name2 == 2)
	fmt.Println("verify enum from: success")
}

func assertEnum(condition bool) {
	if condition == false {
		panic(fmt.Errorf("verify enum: verify failed"))
	} // if
}
