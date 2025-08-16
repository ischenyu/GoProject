package main

import "fmt"

// 定义结构体
type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t1 Teacher // 未赋值时默认值: { 0 }
	fmt.Println(t1)
	t1.Name = "ZhangSan"
	t1.Age = 18
	t1.School = "TEST SCHOOL"
}
