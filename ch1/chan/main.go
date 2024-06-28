package main

import (
	"fmt"
	"sync"
	"time"
)

// 数据生成器，将无限数据发送到数据channel
func dataGenerator(dataCh chan<- int) {
	i := 0
	for {
		dataCh <- i
		i++
		time.Sleep(10 * time.Millisecond) // 模拟生成数据的间隔
	}
}

// worker函数，从数据channel读取数据并处理，然后将结果发送到结果channel
func worker(id int, dataCh <-chan int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		// 模拟数据处理
		result := data * 2 // 处理逻辑：将数据乘以2
		fmt.Printf("Worker %d processed data %d into result %d\n", id, data, result)
		resultCh <- result
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
	}
}

func main() {
	const numWorkers = 128

	dataCh := make(chan int, numWorkers)   // 带缓冲的channel，容量为128
	resultCh := make(chan int, numWorkers) // 结果channel，容量为128
	var wg sync.WaitGroup

	// 启动数据生成器
	go dataGenerator(dataCh)

	// 启动固定数量的worker
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataCh, resultCh, &wg)
	}

	// 启动一个goroutine来等待所有worker完成
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	// 在主协程中处理结果
	for result := range resultCh {
		fmt.Printf("Main thread received result: %d\n", result)
	}

	fmt.Println("All workers have completed their tasks")
}
