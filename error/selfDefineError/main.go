package main

import (
	"fmt"
	"errors"
)

func main(){

	num1 := 15
	num2 := 00

	sum, err := addNum(num1, num2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(sum)
	}

}

func addNum(num1, num2 int) (int, error){
	if num2 == 0 {	
		return -1, errors.New("ERROR")
	}else {
		return num1 + num2, nil
	}
}
