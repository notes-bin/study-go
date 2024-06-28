package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前时间，注意此时返回的是 time.Time 类型
	now := time.Now()
	fmt.Println(now)
	// 当前时间戳
	fmt.Println(now.Unix())
	// 纳秒级时间戳
	fmt.Println(now.UnixNano())
	// 时间戳小数部分 单位：纳秒
	fmt.Println(now.Nanosecond())

	// 返回日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())

	// 时分秒
	hour, minute, second := now.Clock()
	fmt.Printf("hour:%d, minute:%d, second:%d\n", hour, minute, second)
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())

	// 返回星期
	fmt.Println(now.Weekday())
	// 返回一年中对应的第几天
	fmt.Println(now.YearDay())
	// 返回时区
	fmt.Println(now.Location())

	// 返回一年中第几天
	fmt.Println(now.YearDay())

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	layout := "2006-01-02 15:04:05"
	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
	fmt.Println(t.Format(layout))

	layout1 := "2006-01-02 15:04:05"

	// 根据指定时间返回 time.Time 类型
	// 分别指定年，月，日，时，分，秒，纳秒，时区
	t1 := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, now.Location())
	fmt.Println(t1.Format(layout1))

	// 输出 2021-01-10 17:28:50 +0800 CST
	// time.Local 指定本地时间
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t2)

}
