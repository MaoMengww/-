package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type task struct { //任务
	filePath string
	keyWord  string
}

type result struct {
	filePath string
	lineNum  int
	content  string
}

// 搜索文件，分发任务
func searchFile(filePath, word string, jobs []chan task) {
	defer func() {
		for _, jobChan := range jobs {
			close(jobChan)
		}
	}()
	i := 0 //用于区分员工管道
	err := filepath.WalkDir(filePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accessing %s: %v\n", path, err)
			return nil //跳过
		}
		if !d.IsDir() {
			wg.Add(1)
			task := task{
				filePath: path,
				keyWord:  word,
			}
			jobs[i%len(jobs)] <- task
			i++
		}
		return nil
	})
	if err != nil {
		log.Printf("无法打开文件")
		return
	}
}

// 在文件中搜寻keyword
func searchInFile(filePath, keyWord string, resultChan chan<- result) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("无法打开文件")
		return
	}
	defer file.Close()

	lineNum := 0
	scanner := bufio.NewScanner(file) //扫描器，逐行读取文件，成功返回true，失败或读取结束返回false
	for scanner.Scan() {              //成功返回true，失败或读取结束返回false
		lineNum++
		lineContent := scanner.Text()
		if strings.Contains(lineContent, keyWord) { //是否包含
			resultChan <- result{
				filePath: filePath,
				lineNum:  lineNum,
				content:  lineContent,
			}
		}
	}
	//scanner读取中途出错
	if err := scanner.Err(); err != nil {
		log.Printf("读取文件 %v 时出错: %v", filePath, err)
	}
}

// 工人
func worker(jobs <-chan task, reresults chan<- result) {
	for job := range jobs {
		searchInFile(job.filePath, job.keyWord, reresults)
	}
}

// 打印结果
func printResults(reresults <-chan result, done chan<- bool) {
	for res := range reresults {
		fmt.Printf("%s:%d:%s\n", res.filePath, res.lineNum, res.content)
	}
	done <- true
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("参数错误")
	}
	filePath := os.Args[1]
	keyWord := os.Args[2]

	//配置协程池
	const workerNum = 100 //工人数
	//每个工人一个专属通道
	jobs := make([]chan task, workerNum)
	reresults := make(chan result, 100)
	printDone := make(chan bool) // 用于等待打印完成

	for i := range workerNum {
		jobs[i] = make(chan task, 100)
	}

	//启动打印协程
	go printResults(reresults, printDone)

	//启动员工
	for i := range 100 {
		go worker(jobs[i], reresults)
	}

	//分发任务
	searchFile(filePath, keyWord, jobs)
	wg.Wait()
	close(reresults)
	fmt.Println("所有文件搜索完成，正在等待打印...")

	<-printDone

	fmt.Println("搜索和打印均已完成")
}
