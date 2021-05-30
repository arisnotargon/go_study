package main

import "fmt"

type Integer int

func (a Integer) Equal(b Integer) bool {
	// 给新类型Integer添加一个判断是否相等的方法
	return a == b
}

func (a Integer) Add(b Integer) Integer {
	return a + b
}
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}

func CheckMath(m Math) bool {
	fmt.Println(m)
	return true
}

func main() {
	var x, y Integer
	x, y = 10, 10
	fmt.Println(x.Equal(y))
	fmt.Println(x.Add(y))
	CheckMath(x)
}
