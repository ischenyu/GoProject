package main

import (
	"fmt"
)

/*
数组的注意事项:
1. 长度属于数组的一部分 (见main.go 17:6)
2. Go中数组属值类型, 在默认情况下是值传递, 因此会进行拷贝    [main.go 51-68]
3. 如果想在其他函数中修改原来的数组, 可以使用值传递(指针方式) [main.go 51-68]
*/

func main() {
	// 定义数组
	var arr [3]int
	fmt.Printf("现在数组为: %v, 数组的类型为: %T\n", arr, arr) //[0 0 0] 声明后的数组默认为0
	fmt.Printf("数组的长度为: %v\n", len(arr))

	// 打印证明arr中存储的是地址值
	fmt.Printf("数组的地址为: %p\n", &arr)

	// arr存储的地址值是指向数组中下标为0的地址
	// fmt.Printf("数组索引为0的地址为: %p\n", &arr[0])
	for i := 0; i < len(arr); i++ {
		// 数组每个空间占用字节取决于数组类型
		fmt.Printf("数组索引为%v的地址为: %p\n", i, &arr[i])
	}

	// 赋值
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	for i := 0; i < len(arr); i++ {
		fmt.Printf("请录入下标为%v的数值: ", i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println("更新后的数组: ", arr)

	// 遍历数组:
	// <法一>
	for i := 0; i < len(arr); i++ {
		// 数组每个空间占用字节取决于数组类型
		fmt.Printf("数组索引为%v的值为: %v\n", i, arr[i])
	}
	// <法二> for-range
	for key, val := range arr {
		fmt.Printf("数组索引为%v的值为: %v\n", key, val)
	}

	arrFunc(arr)
	fmt.Printf("原数组: %v\n", arr)

	arrFuncPtr(&arr)
	fmt.Printf("原数组: %v\n", arr)
}

func arrFunc(arr [3]int) {
	// 这里的arr是外部拷贝过来的, 修改这个不会影响外部数组, 如要修改, 需使用指针
	arr[0] = 60
	fmt.Printf("arrFunc更改后的数组: %v\n", arr)
}

func arrFuncPtr(arr *[3]int) {
	// 这里的arr是外部拷贝过来的, 修改这个不会影响外部数组, 如要修改, 需使用指针
	arr[0] = 60
	fmt.Printf("arrFuncPtr更改后的数组: %v\n", arr)
}
