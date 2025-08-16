package main

import "fmt"

/*
结构体之间的转换
1. 结构体是用户单独定义的类型， 和其他类型进行转换时需要完全相同的字段(名字 个数 类型)
2. 结构体进行type重新定义(相当于取别名, 见15-17行), Golang会认为是新的数据类型, 但它们相互间可以强转
*/

type Student struct {
	Age int
}

type Person struct {
	Age int
}

type Stu Student

func main() {
	var s Student = Student{10}
	var p Person = Person{20}
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)

	var s1 Student = Student{10}
	var s2 Stu = Stu{19}
	s1 = Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
