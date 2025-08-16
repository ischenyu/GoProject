package main

import "fmt"

func main() {
	// 定义数组
	arr := [6]int{1, 2, 3, 4, 5, 6}
	// 切片构建在数组之上
	// [1:3]切片 - 切除的一段片段 - 索引从1开始, 到3结束(不包含3) -> 切割区间前闭后开 [1, 3)
	slice := arr[1:3]
	fmt.Println("slice:" ,slice)
	fmt.Println("slice长度: ", len(slice))
	fmt.Println("slice容量: ", cap(slice))

	// 定义切片
	slice1 := make([]int, 4, 20) // 定义一个int类型的切片, 长度为4, 容量为20
	fmt.Println(slice1)
	fmt.Println("slice1长度:", len(slice1))
	fmt.Println("slice1容量:", cap(slice1))

	// 向切片中追加元素
	// 1.底层追加元素的时候对数组进行扩容, 老数组扩容为新数组
	// 2.创建一个新数组, 将老数组中的2, 3复制到新数组中, 在新数组中追加88, 50
	// 3.底层数组指向的是新数组, 老数组不发生改变
	slice2 := append(slice, 88, 50)
	fmt.Println(slice2)
	// 4.如果只想给原数组追加, 只需
	slice = append(slice, 88, 50)
	fmt.Println(slice)
	slice = append(slice, slice1...)
	fmt.Println(slice)

	// 切片拷贝
	var a []int = []int{1,4,7,2,5,8}
	b := make([]int,10)
	// 将a的元素复制给b
	copy(b,a)
	fmt.Println(b) // result: [1 4 7 2 5 8 0 0 0 0]
}
