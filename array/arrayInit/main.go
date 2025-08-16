package main

import (
	"fmt"
)

func main() {
	// 数组初始化的四种方式
	// 1.
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 2.
	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 3.
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// 4.
	var arr4 = [5]int{2:66, 0:33, 1:99, 3:88} // 对指定索引进行赋值 格式: 索引:值
	fmt.Println(arr4)
}
