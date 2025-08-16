package main

import (
	"fmt"
)

func test(args ...int) {
	fmt.Printf("%v\n", args)
	fmt.Println(len(args))
}

// 闭包
// 返回一个func(int) int类型的值, 类型为函数
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

// defer
func add(num1 ,num2 int) int{
	// defer将其后的语句压入栈中,继续执行后面的语句
	// 栈的特点: 先进后出
	// 在函数执行完毕后, 从栈中取出语句, 开始执行
	defer fmt.Println("num1: ", num1)
	defer fmt.Println("num2: ", num2)
	fmt.Println("sum: ", num1 + num2)
	return num1 + num2
}

func main() {
	var str string = "Hello Golang你好"

	for i, value := range str {
		fmt.Printf("%d, %c\n", i, value)
	}
	test(0, 1, 2, 3, 4, 5)
    fmt.Println(&str)

	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(1))

	fmt.Println(add(30, 60))

}
