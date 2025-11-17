//
//示例由ai生成
//
//
package main

import (
	"fmt"
	"sync"

)

// Counter 结构体从 Lv2/main.go 复制而来
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var counter = Counter{}

	// 配置协程池
	const taskNum = 100000 // 任务数 (与 Lv2/main.go 一致)
	const workerNum = 100  // 工人数 (与 Lv2/main.go 一致)
	const queueSize = 100  // 每个 worker 的队列缓冲

	// 1. 使用 gopool.New 创建协程池
	pool := New(workerNum, uint(queueSize))

	// 2. 启动协程池的 worker
	pool.Run()

	// 3. 定义我们希望协程池执行的“任务”
	//    (这就是原来 Lv2/main.go 中 worker 的核心逻辑)
	task := func() {
		for range 100 {
			counter.Increment()
		}
	}

	// 4. 将所有任务添加到协程池
	fmt.Println("开始分发任务...")
	for i := 0; i < taskNum; i++ {
		// gopool 的 Add 方法会自动处理 WaitGroup 和轮询分发
		pool.Add(task)
	}
	fmt.Println("任务分发完成")

	// 5. 等待 gopool 中所有的任务执行完毕
	pool.Wait()
	fmt.Println("任务完成")

	// 6. 关闭协程池
	pool.Shutdown()

	// 打印最终结果 (应为 100000 * 100 = 10,000,000)
	fmt.Println("最终计数：", counter.Value())
}