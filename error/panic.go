package main

import (
	"fmt"
)

// recover相当于catch,接收前一个panic的报错

func divide() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()
	var i = 1
	var j = 0
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)
}
func main() {
	divide()
	fmt.Println("divide 方法调用完毕，回到 main 函数")
}
