package main

import "fmt"

func sumFunc1(x int, y int) int {

	sum := x + y
	return sum
}

// 可变参数函数
func sumFunc2(x ...int) {
	fmt.Printf("%v", x)
}

func sumFunc3(x string, y ...int) {
	fmt.Printf("%v", x)
}

// 返回多个值
func sumFunc4(x int, y int) (int, int) {
	sum1 := x + y
	def1 := x - y
	return sum1, def1
}

type calcuType func(int, int) int

func clacuFunc(x int, y int, calcu calcuType) int {
	return calcu(x, y)
}

func funcPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("抛出一个异常")
}

func main() {
	// 1.普通函数
	sum1 := sumFunc1(1, 2)
	fmt.Println(sum1)

	// 2.可变参数函数
	sumFunc2(1, 2, 3, 4)

	// 3.可变参数函数
	sumFunc3("test", 2, 3, 4)

	// 4.返回多个值
	fmt.Println(sumFunc4(2, 10))

	// 5.自定义函数并将函数当做参数传递
	fmt.Println(clacuFunc(1, 2, sumFunc1))

	// 6.匿名函数
	func() {
		fmt.Println("匿名函数")
	}()

	// 7.panic/revocer异常处理
	funcPanic()

}
