# Go Routine 示例：并发处理与线程间通信

下面是一个完整的 Go 程序示例，展示了如何使用 goroutine 以及通过通道（channel）进行线程间通信。程序模拟了一个简单的生产者-消费者场景，并包含详细的注释说明。

```go
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
```

## 代码说明

### 1. 基本结构
- 定义了一个 `Task` 结构体，包含任务ID、输入值和结果字段
- 使用 `sync.WaitGroup` 来等待所有goroutine完成

### 2. 生产者函数
- 生成10个带有随机值的任务
- 将任务发送到带缓冲的通道中
- 生产完成后关闭通道，通知消费者不再有新的任务

### 3. 消费者函数
- 启动3个消费者goroutine并行处理任务
- 使用 `range` 从通道读取任务，直到通道关闭
- 每个消费者处理任务（计算平方值）并输出结果

### 4. 通道通信
- 使用带缓冲的通道实现生产者和消费者之间的解耦
- 生产者关闭通道作为"任务已完成"的信号
- 多个消费者竞争从通道获取任务，展示并发处理

## 运行结果示例

运行此程序会产生类似以下的输出（每次运行结果可能不同）：

```
生产者: 发送任务 1 (值: 87)
消费者 1: 接收任务 1 (值: 87)
生产者: 发送任务 2 (值: 59)
消费者 2: 接收任务 2 (值: 59)
生产者: 发送任务 3 (值: 18)
消费者 3: 接收任务 3 (值: 18)
消费者 1: 任务 1 完成 (结果: 7569)
生产者: 发送任务 4 (值: 40)
消费者 1: 接收任务 4 (值: 40)
消费者 2: 任务 2 完成 (结果: 3481)
生产者: 发送任务 5 (值: 0)
消费者 2: 接收任务 5 (值: 0)
...
消费者 3: 任务 10 完成 (结果: 16)
消费者 3: 检测到通道已关闭，退出处理循环
消费者 1: 检测到通道已关闭，退出处理循环
消费者 2: 检测到通道已关闭，退出处理循环
所有任务处理完成!
```

这个示例展示了Go语言中goroutine的基本用法、通道的创建和使用、以及如何使用WaitGroup同步goroutine。通过注释，详细解释了每个关键部分的作用和工作原理。
