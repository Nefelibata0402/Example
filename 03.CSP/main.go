package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // 模拟任务处理
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动三个goroutines作为工作协程
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送5个任务到通道
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs) // 关闭任务通道，表示没有更多任务

	// 收集结果
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Printf("Received result: %d\n", result)
	}
}
