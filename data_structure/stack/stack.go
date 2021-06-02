package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// 初始化空栈
var size = 0
var stack = new(Node)

func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size++
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop() (v int, b bool) {
	if size == 0 {
		// 空栈
		b = false
		return
	}

	if size == 1 {
		// 仅一个节点
		v = stack.Value
		stack = nil
		b = true
		size--
		return
	}
	v = stack.Value
	stack = stack.Next
	b = true
	size--
	return
}

func traverse() {
	if size == 0 {
		fmt.Println("空栈!")
		return
	}
	t := stack
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println("===end===")
}

func main() {
	stack = nil
	// 读取空栈
	v, b := Pop()
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() 失败!")
	}
	// 进栈
	Push(100)
	// 遍历栈
	traverse()
	// 再次进栈
	Push(200)
	// 再次遍历栈
	traverse()
	// 批量进栈
	for i := 0; i < 10; i++ {
		Push(i)
	}
	// 批量出栈
	for i := 0; i < 15; i++ {
		v, b := Pop()
		if b {
			fmt.Print(v, " ")
		} else {
			// 如果已经是空栈，则退出循环
			break
		}
	}
	fmt.Println()
	// 再次遍历栈
	traverse()
}
