package main

import (
	"fmt"
	"strings"
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

	// 3.float精度丢失问题
	floa1 := 3.2
	floa2 := 4.8
	fmt.Println(floa2 - floa1)

	// 4.字符串常用方法
	substring := "hello world"
	fmt.Println(len(substring))
	fmt.Println(strings.Index(substring, "o"))

	// 5.数组
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	var arr3 = [...]int{1: 1, 3: 3}
	fmt.Println(arr3)

	// 6多维数组
	var arr4 = [3][2]string{
		{"a", "b"},
		{"b", "c"},
		{"d", "e"},
	}
	fmt.Println(arr4)

	// 7.切片
	var slice1 = []int{1, 2, 3}
	slice2 := slice1
	slice2[2] = 4
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 切片扩容
	slice2 = append(slice2, 10)
	fmt.Println(slice2)

	// 切片copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 3, 3)
	copy(slice4, slice3)
	slice4[2] = 4
	fmt.Println(slice3)
	fmt.Println(slice4)

	// 8.map
	map1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map1)

	map2 := make(map[string]int)
	fmt.Println(map2)

	delete(map1, "a")
	fmt.Println(map1)

	// 9.切片和map结合
	userInfo := make(map[string][]string)
	userInfo["name"] = []string{"a", "b", "c"}
	userInfo["sex"] = []string{"male", "female"}
	fmt.Println(userInfo)
}
