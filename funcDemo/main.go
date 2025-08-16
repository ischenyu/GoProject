package main

import (
	"fmt"
)

func init(){
	fmt.Println("Init function is progressed.")
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
	// 在函数执行完毕后, 从栈中取出语句, 开始执行(先进后出)
	// 先输出num2, 后输出num1
	defer fmt.Println("num1: ", num1)
	defer fmt.Println("num2: ", num2)
	// 这里即使对num1 num2的值作了变动, 但上面两个语句不发生改变, 因为在上面两个阶段中num1 num2的值就是30 & 60
	num1 += 90
	num2 += 60
	fmt.Println("sum: ", num1 + num2)
	return num1 + num2
	// defer应用场景: 
	// 想关闭某个使用资源, 在使用的时候随手defer, 因为defer有延迟执行机制, (函数执行完毕再执行defer压入栈的语句)
	//    所以用完随手写了关闭, 比较省心 省事(底层会帮忙处理)
}

func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(1))

	fmt.Println(add(30, 60))

}
