package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间 func Now() Time
	//    type Time
	now := time.Now()
	fmt.Printf("%v, \n类型: %T\n", now, now)
	fmt.Printf("年: %v 月: %v 日: %v\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("时: %v 分: %v 秒: %v\n", now.Hour(), now.Minute(), now.Second())

	// 2. 日期格式化
	// <No.1>
	datestr := fmt.Sprintf("年月日: %d-%d-%d	时分秒: %d:%d:%d", now.Year(), int(now.Month()),
	now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)
	// <No.2>
	// "2006/01/02 15/04/05"必须这样写, 数字固定
	datestr2 := now.Format("2006/01/02 15:04:05")
	fmt.Println(datestr2)
}
