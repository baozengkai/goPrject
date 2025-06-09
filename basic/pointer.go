package main

import "fmt"

func main() {
	// 1.new初始化1
	var a = new(int)
	fmt.Println(*a)

	// 2.new初始化2
	var b *int
	b = new(int)
	*b = 100
	fmt.Println(*b)
}
