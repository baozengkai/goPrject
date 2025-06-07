package main

import (
	"fmt"
)

func main() {
	// 1.变量
	testVar := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", testVar)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	// 2.常量
	const pi = 3.1415926
	fmt.Println(pi)
}
