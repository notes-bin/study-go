package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("微客鸟窝")
	fmt.Println("我是无尘啊")
	time.Sleep(time.Second) //等待一秒，使goroutine 执行完毕
}
