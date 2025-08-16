package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	// 1.统计字符串的长度, 按字节进行统计(一个中文占三个字节, 会把三个字节拆开)
	// [内置函数] func len(str string) int
	str := "Hello Golang"
	fmt.Println(len(str))

	// 2.字符串遍历
	// <法一> for-range键值循环遍历
	for i, value := range str{
		fmt.Printf("索引为: %d, 具体的值为: %c\n", i, value)
	}
	// <法二>r := rune(str) [切片]
	r:= []rune(str)
	for i := 0; i < len(r); i++{
		fmt.Printf("%c", r[i])
		
	}

	fmt.Println()

	// 3. 字符串转整数
	// [strconv.Atoi] func Atoi(s string) (i int, err error)
	var stringconv string = "114514"
	convednum, err := strconv.Atoi(stringconv)
	if err != nil{
		return
	}
	fmt.Println(convednum)

	// 4. 整数转字符串
	//    [strconv.Itoa] func Atoi(i, int) string

	// 5. 统计一个字符串有几个指定的子串
	//    [strings.Count] func Count(s, step string) int
	countedNum := strings.Count(str, "l")
	fmt.Println(countedNum)

	// 6. 查找子串是否在指定字符串中
	//    [strings.Contains] func ContainsAny(s, chars string) bool

	// 7. 不区分大小写字符串的比较
	//    [strings.EqualFold] func ContainsAny(s, t string) bool

	// 8. 子串sep在字符串s中第一次出现的位置, 不存在则返回-1
	//    [strings.Index] func Index(s, sep string) int

}
