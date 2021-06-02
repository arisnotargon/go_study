package main

import "fmt"

//自己实现一个面向对象的单链表

type Node struct {
	value int
	next  *Node
}

func NewLinkedList(v int) Node {
	return Node{v, nil}
}

// Add 在链表末端添加元素
func (t *Node) Add(v int) int {
	if t.next == nil {
		// 不存在下个节点时
		next := &Node{v, nil}
		t.next = next
		return 0
	} else {
		// 下个节点不为空时
		return t.next.Add(v)
	}
}

// Insert 在当前节点后面插入一个节点
func (t *Node) Insert(v int) {
	next := t.next
	t.next = &Node{v, next}
}

// Traverse 遍历打印链表之后的全部
func (t *Node) Traverse() {
	fmt.Println("====开始遍历====")
	for t != nil {
		fmt.Println("%d=>", t.value)
		t = t.next
	}

	fmt.Println("====遍历完成====")
}

// Size 获取链表长度
func (t *Node) Size() int {
	size := 0
	for t != nil {
		size++
		t = t.next
	}
	return size
}

// LookUp 判断链表中是否存在值
func (t *Node) LookUp(v int) bool {
	for t != nil {
		if v == t.value {
			return true
		}
		t = t.next
	}
	return false
}

// Del 删除从0开始的第i个元素
func (t *Node) Del(i int) bool {
	size := t.Size()
	if i >= size {
		return false
	}
	if i == 0 {
		*t = *t.next
		return true
	}

	for j := 0; j < i-1; j++ {
		t = t.next
	}
	t.next = t.next.next
	return true
}

func main() {
	head := NewLinkedList(1)

	head.Add(2)
	head.Add(3)
	head.Add(4)

	head.Traverse()
	fmt.Println(head.Size())
	fmt.Println(head.LookUp(2))
	head.Del(3)
	head.Traverse()
}
