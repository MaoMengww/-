package main

import (
	"fmt"
)

func main() {
	var sum, count int
	var input int
	for {
		fmt.Println("请输入整数(输入0结束): ")
		fmt.Printf("请输入第%v个数:\n ", count + 1)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("非法输入")
			continue
		}
		if input == 0 {
			break
		}
		sum = sum + input
		count++
	}
	avg := float64(sum) / float64(count)
	if avg >= 60 {
		fmt.Printf("平均成绩为 %.2f，成绩合格\n", avg)
	} else {
		fmt.Printf("平均成绩为 %.2f，成绩不合格\n", avg)
	}
}