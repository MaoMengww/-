package main

import (
	"sync"
	"sync/atomic"
)

type Task func()

// 工作池
type Pool struct {
	workerNum   int            // worker 数量
	jobChans    []chan Task    // worker专属channel
	index       atomic.Int64   //任务索引
	queenSize   uint           //channel缓冲
	taskWg      sync.WaitGroup // 任务
	workerWg    sync.WaitGroup // woker
	stopCh      chan struct{}  // 通知所有 worker 停止
}

//创建新池
func New(workerNum int, queenSize uint) *Pool {
	jobs := make([]chan Task, workerNum)
	for i := range workerNum {
		jobs[i] = make(chan Task, queenSize)
	}
	index := atomic.Int64{}
	index.Store(0)
	return &Pool{
		workerNum: workerNum,
		queenSize: queenSize,
		jobChans:  jobs,
		index:     index,
		stopCh:    make(chan struct{}),
	}
}

//增加任务
func (p *Pool) Add(task Task) {
	if task == nil {
		return
	}
	p.taskWg.Add(1)
	indexsum := p.index.Add(1)
	index := (indexsum - 1) % int64(p.workerNum)
	select {
	case <-p.stopCh:
		p.taskWg.Done()
		p.index.Add(-1)
		return
	case p.jobChans[index] <- task:
	}
}

// Run 启动协程池，分发任务
func (p *Pool) Run() {
	p.workerWg.Add(p.workerNum)
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
}

// 工人工作
func (p *Pool) worker(id int) {
	defer p.workerWg.Done()
	// 获取自己对应的 channel
	jobChan := p.jobChans[id]

	for {
		select {
		case <-p.stopCh:
			return

		case task, ok := <-jobChan:
			if !ok {
				return
			}
			if task != nil {
				task()
			}
			p.taskWg.Done()
		}
	}
}

// 等待所有任务完成
func (p *Pool) Wait() {
	p.taskWg.Wait()
}

//关闭
func (p *Pool) Shutdown() {
	close(p.stopCh)
	p.workerWg.Wait()
	for _, ch := range p.jobChans {
		close(ch)
	}
}


