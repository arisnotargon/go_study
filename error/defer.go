package main

import "fmt"

// defer 兜底逻辑
func printError() {
	fmt.Println("兜底执行")
}
func main() {
	defer printError()
	defer func() {
		fmt.Println("除数不能是0!")
	}()
	var i = 1
	var j = 1
	var k = i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)
}
