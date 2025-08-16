package main

import "fmt"

/*
结构体实例的创建方法
*/

// 定义结构体
type Teacher struct {
	Name   string
	Age    int
	School string
}

func func1() {
	var t1 Teacher // 未赋值时默认值: { 0 }
	fmt.Println(t1)
	t1.Name = "ZhangSan"
	t1.Age = 18
	t1.School = "TEST SCHOOL"
	t1.Age++
	fmt.Println(t1)
}

func func2() {
	// 创建老师结构体的实例 对象 变量
	var t Teacher = Teacher{"ZhangSan", 18, "TEST SCHOOL"}
	fmt.Println(t)
}

func func3() {
	// 返回结构体指针
	// 创建老师结构体的实例 对象 变量
	var t *Teacher = new(Teacher)
	// t是指针, 其实指向的就是这个地址, 应该给地址指向的对象的字段赋值:
	(*t).Name = "ZhangSan"
	(*t).Age = 18
	t.School = "TEST SCHOOL" // go编译器对 t.School 进行了转化 --> (*t).School
	fmt.Println(*t)
}

func func4() {
	// 创建老师结构体的实例 对象 变量
	var t *Teacher = &Teacher{}
	// t是指针, 其实指向的就是这个地址, 应该给地址指向的对象的字段赋值:
	(*t).Name = "ZhangSan"
	(*t).Age = 18
	t.School = "TEST SCHOOL" // go编译器对 t.School 进行了转化 --> (*t).School
	fmt.Println(*t)
}

func main() {
	fmt.Println("==========First==========")
	func1()
	fmt.Println("==========Second==========")
	func2()
	fmt.Println("==========Third==========")
	func3()
}
