package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

type User struct {
	name    string
	age     int
	address Address
}

type Address struct {
	Street string
}

func main() {
	// 1.声明和初始化第一种
	var p1 Person
	p1.name = "<UNK>"
	p1.age = 18
	p1.sex = "Male"
	fmt.Println(p1)

	// 2.初始化第二种
	var p2 = new(Person)
	p2.name = "<UNK>"
	p2.age = 18
	p2.sex = "Male"

	// 3.初始化第三种
	var p3 = Person{
		Name: "<UNK>",
		Age:  18,
		Sex:  "Male",
	}
	fmt.Println(p3)
}
