package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// 模拟一个长时间运行的goroutine
		for {
			select {
			case <-ctx.Done():
				log.Println("收到退出信号，优雅退出")
				return
			default:
				log.Println("正在运行...")
				time.Sleep(time.Second)
			}
		}
	}()

	// 监听退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 等待退出信号
	<-sigCh

	// 发送取消信号
	cancel()

	// 执行一些清理操作
	log.Println("执行清理操作...")

	// 等待一段时间，以确保所有goroutine都退出
	time.Sleep(time.Second * 2)
	log.Println("退出程序")
}
