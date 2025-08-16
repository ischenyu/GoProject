
package main

import (
	"fmt"
)

func main(){
	// 1. func new([type]) *type 分配内存, new函数的返回值是对应类型的指针
	num := new(int)
	*num = 25
	fmt.Printf("Type: %T, num: %v, Address: %v\n", num, *num, num)
}
