package main

import (
    "fmt"
    "sync"
)

var once sync.Once

// initFunc是需要只执行一次的初始化函数
func initFunc() {
    fmt.Println("Initialization function executed")
}

func worker(wg *sync.WaitGroup) {
    defer wg.Done()
    once.Do(initFunc) // 确保initFunc只执行一次
    fmt.Println("Worker executed")
}

func main() {
    var wg sync.WaitGroup

    // 启动多个协程
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(&wg)
    }

    wg.Wait()
    fmt.Println("All workers completed")
}