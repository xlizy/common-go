package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "L"
	n := 123
	desiredLength := 8
	paddedString := fmt.Sprintf("%0[1]*s", desiredLength, strconv.Itoa(n))
	s = s + paddedString
	fmt.Println(s)
}
