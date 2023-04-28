package main

import (
	"fmt"
	"github.com/xlizy/common-go/crypto"
)

func main() {
	ss := crypto.AesEncryptECB(fmt.Sprintf("%v", 193446002458628097), "jpEUQ7wPbfvAd08z3o2QWZ4P1Doa9ayx")
	fmt.Println("======")
	fmt.Println(ss)
	fmt.Println("======")
}
