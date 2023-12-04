package main

import (
	"fmt"
	"sync"
)

// Worker 是协程池的工作协程结构体
type Worker struct {
	ID     int           // 协程的唯一标识符
	TaskCh chan Runnable // 用于接收任务的通道
}

// Runnable 定义任务接口
type Runnable interface {
	Run(w *Worker)
}

// NewWorker 创建一个新的 Worker
func NewWorker(id int, taskCh chan Runnable) *Worker {
	return &Worker{
		ID:     id,
		TaskCh: taskCh,
	}
}

// Start 启动 Worker，监听任务通道并执行任务
func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range w.TaskCh {
		task.Run(w)
	}
}

// Pool 是协程池结构体
type Pool struct {
	workers []*Worker      // 存储所有的工作协程
	taskCh  chan Runnable  // 用于分发任务的通道 这个通道可以接受任何实现了 Runnable 接口的类型的值。
	wg      sync.WaitGroup // 用于等待所有工作协程完成任务
}

// NewPool 创建一个新的协程池
func NewPool(size int) *Pool {
	return &Pool{
		workers: make([]*Worker, size),
		taskCh:  make(chan Runnable),
	}
}

// Start 启动协程池，初始化工作协程并等待它们完成
func (p *Pool) Start() {
	p.wg.Add(len(p.workers))

	for i := 0; i < len(p.workers); i++ {
		p.workers[i] = NewWorker(i, p.taskCh)
		go p.workers[i].Start(&p.wg)
	}

	go func() {
		p.wg.Wait()
		close(p.taskCh)
	}()
}

// Submit 提交任务到协程池
func (p *Pool) Submit(task Runnable) {
	p.taskCh <- task
}

// Shutdown 关闭协程池，等待所有工作协程完成任务
func (p *Pool) Shutdown() {
	close(p.taskCh)
	p.wg.Wait()
}

// ExampleTask 是一个示例任务，实现了 Runnable 接口
type ExampleTask struct {
	ID int
}

// Run 实现 Runnable 接口的 Run 方法
func (t *ExampleTask) Run(w *Worker) {
	fmt.Printf("Goroutine %d execute Task %d is running\n", w.ID, t.ID)
}

func main() {
	// 创建一个协程池，指定协程池大小为3
	pool := NewPool(3)

	// 启动协程池
	pool.Start()

	// 提交一些任务到协程池
	for i := 1; i < 100; i++ {
		task := &ExampleTask{ID: i}
		pool.Submit(task)
	}

	// 关闭协程池，等待所有任务完成
	pool.Shutdown()
}
