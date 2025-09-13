package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义任务结构体
type Task struct {
	id     int
	value  int
	result int
}

func main() {
	// 创建一个带缓冲的通道，用于在生产者和消费者之间传递任务
	// 缓冲区大小为5，允许生产者和消费者在一定程度上解耦
	taskQueue := make(chan *Task, 5)
	
	// 创建一个等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup
	
	// 设置随机数种子，用于生成随机任务值
	rand.Seed(time.Now().UnixNano())
	
	// 启动生产者goroutine
	wg.Add(1) // 增加等待组计数
	go producer(taskQueue, &wg)
	
	// 启动3个消费者goroutine，展示多个消费者并行处理
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 为每个消费者增加等待组计数
		go consumer(i, taskQueue, &wg)
	}
	
	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("所有任务处理完成!")
}

// 生产者函数：生成任务并发送到任务队列
func producer(taskQueue chan<- *Task, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时减少等待组计数
	
	// 生成10个任务
	for i := 1; i <= 10; i++ {
		// 创建新任务
		task := &Task{
			id:    i,
			value: rand.Intn(100), // 生成0-99的随机数
		}
		
		fmt.Printf("生产者: 发送任务 %d (值: %d)\n", task.id, task.value)
		
		// 将任务发送到通道（如果通道已满，这里会阻塞）
		taskQueue <- task
		
		// 模拟生产所需时间
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	
	// 所有任务生产完成后关闭通道
	// 这是重要的步骤，它告诉消费者没有更多任务了
	close(taskQueue)
	fmt.Println("生产者: 所有任务已发送，通道已关闭")
}

// 消费者函数：从任务队列获取任务并处理
func consumer(id int, taskQueue <-chan *Task, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时减少等待组计数
	
	// 循环从通道接收任务，直到通道关闭
	for task := range taskQueue {
		fmt.Printf("消费者 %d: 接收任务 %d (值: %d)\n", id, task.id, task.value)
		
		// 处理任务：计算值的平方
		task.result = task.value * task.value
		
		// 模拟处理任务所需时间
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		
		fmt.Printf("消费者 %d: 任务 %d 完成 (结果: %d)\n", id, task.id, task.result)
	}
	
	fmt.Printf("消费者 %d: 检测到通道已关闭，退出处理循环\n", id)
}
