package main

import (
	"fmt"
	"testing"
)

func TestSscanf(t *testing.T) {
	// 示例字符串
	str := "2006 Jan 02 15:04:05"

	t.Helper()
	var month, day, year int
	// 解析字符串
	// %s 表示字符串，%d 表示整数
	n, err := fmt.Sscanf(str, "%d %s %d", &month, &day, &year)
	if err != nil {
		t.Logf("解析错误: %#v\n", err)
	} else {
		t.Logf("数量: %d, 月份: %d, 日期: %d, 年份: %d\n", n, month, day, year)
	}

	// 另一个示例，解析包含浮点数的字符串
	valueStr := "3.14159"
	var value float64
	_, err = fmt.Sscanf(valueStr, "%f", &value)
	if err != nil {
		t.Logf("解析错误: %s\n", err)
	} else {
		t.Logf("解析的浮点数: %f\n", value)
	}
}
