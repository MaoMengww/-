package main

import "fmt"

func hello() {
	fmt.Println("hello, MaoMeng")
}
func area() {
	pi := 3.14
	r := 5
	area := pi * float64(r * r)
	fmt.Println(area)
}

func add() {
	result := 0
	for i := 0; i <= 1000; i++{
		result = result + i
	}
	fmt.Println(result)
}
func main() {
	hello()
	area()
	add()
}