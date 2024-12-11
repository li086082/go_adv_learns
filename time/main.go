package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now.Format(time.DateTime))      // 2024-12-09 17:01:43
	fmt.Println(now.Format(time.DateOnly))      // 2024-12-09
	fmt.Println(now.Format(time.TimeOnly))      // 17:01:43
	fmt.Println(now.Year())                     // 年
	fmt.Println(fmt.Sprintf("%d", now.Month())) // 月
	fmt.Println(now.Day())                      // 日
	fmt.Println(now.Hour())                     // 时
	fmt.Println(now.Minute())                   // 分
	fmt.Println(now.Second())                   // 秒
	y, m, d := now.Date()
	fmt.Println(y, m, d)
	h, i, s := now.Clock()
	fmt.Println(h, i, s)

	fmt.Println("============")

	fmt.Println(now.Unix())      // 秒
	fmt.Println(now.UnixMilli()) // 毫秒
	fmt.Println(now.UnixMicro()) // 微秒
	fmt.Println(now.UnixNano())  // 纳秒

	fmt.Println("============")
	// 时间戳转时间
	var a1 int64 = 1733812652
	fmt.Println(time.Unix(a1, 0).Format(time.DateTime))

	fmt.Println("============")
	// 时间转时间戳
	var a2 string = "2024-12-10 14:37:32"
	a22, _ := time.Parse(time.DateTime, a2)
	fmt.Println(a22.Unix())
	fmt.Println(a22.UnixMilli())

	a3 := time.Date(2024, 12, 10, 14, 37, 32, 0, time.Local)
	fmt.Println(a3.Unix())

	fmt.Println("============")

	// 计算2个时间戳相差天数
	var a int64 = 1733812652 // 2024-12-10 14:37:32
	var b int64 = 1732982400 //2024-12-01 00:00:00
	as := time.Unix(a, 0)
	bs := time.Unix(b, 0)
	cs := as.Sub(bs)
	fmt.Println(cs.String()) //230h37m32s
	fmt.Println(cs.Hours() / 24)
}
