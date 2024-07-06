package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    // 存储键值对
    m.Store("foo", "bar")
    m.Store("baz", 42)

    // 加载键值对
    if value, ok := m.Load("foo"); ok {
        fmt.Println("Loaded value:", value)
    }

    // 加载或存储键值对
    if actual, loaded := m.LoadOrStore("baz", 100); loaded {
        fmt.Println("Key 'baz' already exists with value:", actual)
    } else {
        fmt.Println("Stored new value:", actual)
    }

    // 遍历键值对
    m.Range(func(key, value interface{}) bool {
        fmt.Println("Key:", key, "Value:", value)
        return true
    })

    // 删除键值对
    m.Delete("foo")
    if _, ok := m.Load("foo"); !ok {
        fmt.Println("Key 'foo' deleted")
    }
}