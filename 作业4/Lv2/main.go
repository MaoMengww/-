package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	count int
}

func (c *Counter) Increment(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(id int, wg *sync.WaitGroup, counter *Counter, jobs <-chan struct{}) {
	for range jobs {
		for range 100  {
			counter.Increment()
		}
		wg.Done()   //一个任务完成
	}
	fmt.Printf("第%v号工人工作完成\n", id)
}

func main(){
	var wg sync.WaitGroup
	var counter = Counter{}
	
	//配置协程池
	const taskNum = 100000   //任务数
	const workerNum = 100   //工人数

	//每个工人一个专属管道
	jobs := make([]chan struct{}, workerNum)   
	for i := range workerNum {
		jobs[i] = make(chan struct{})
	}

	wg.Add(taskNum)
	for i := range 100 {
		go worker(i, &wg, &counter, jobs[i])
	}

	for i := range taskNum {
		jobs[i%workerNum] <- struct{}{}
	}
	for i := range workerNum {
		close(jobs[i])
	}
	fmt.Println("任务分发完成")
	wg.Wait()
	time.Sleep(1*time.Millisecond)
	fmt.Println("任务完成")
	fmt.Println("最终计数：", counter.Value())
}