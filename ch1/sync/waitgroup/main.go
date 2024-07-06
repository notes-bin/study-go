package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保在函数结束时调用Done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

func test0() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有worker完成
	fmt.Println("All workers completed")
}

func process(data int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing data: %d\n", data)
}

func test1() {
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5}

	for _, d := range data {
		wg.Add(1)
		go process(d, &wg)
	}

	wg.Wait()
	fmt.Println("All data processed")
}

func stage1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Stage 1 completed")
}

func stage2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Stage 2 completed")
}

func test2() {
	var wg sync.WaitGroup

	wg.Add(2)
	go stage1(&wg)
	go stage2(&wg)

	wg.Wait()
	fmt.Println("All stages completed")
}

func main() {
	test0()
	test1()
	test2()
}
