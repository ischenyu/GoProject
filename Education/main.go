package main

import (
	"fmt"
)

func main() {
	var num1 int = 10
	var num2 int = 20
	result := sum(num1, num2)
	fmt.Println(result)
	sum1(num1, num2)
}

func sum(num1, num2 int) int {
	sumedNum := num1 + num2
	return sumedNum
}

func sum1(num1, num2 int) {
	sumedNum := num1 + num2
	fmt.Println("sum1's result:", sumedNum)
}
