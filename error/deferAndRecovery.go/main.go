
package main

import (
	"fmt"
)

func main(){

	// defer + recovery捕获 处理错误
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
			return
		}
	}()
	
	num1 := 15
	num2 := 0

	sum := num1/num2
	fmt.Println(sum)
}
