package main

import (
    "bytes"
    "fmt"
    "sync"
)

var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func main() {
    // 获取缓冲区
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset() // 确保缓冲区是干净的

    // 使用缓冲区
    buf.WriteString("Hello, Pool!")
    fmt.Println(buf.String())

    // 将缓冲区放回池中
    bufferPool.Put(buf)
}