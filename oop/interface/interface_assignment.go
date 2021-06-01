package main

import "fmt"

type Integer int

func (a *Integer) Add(b Integer) {
	*a = (*a) + b
}
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

func (a *Integer) GetValue() Integer {
	return *a
}

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

func main() {
	var x Integer
	var y Integer

	x = 10
	y = 10

	var m, m2 Math

	m = &x
	m2 = &y

	_ = m
	_ = m2

	m.Add(*(m2.(*Integer)))

	fmt.Println(m.Multiply(*(m2.(*Integer))))

	x.Add(y)

	fmt.Println(x)
}
