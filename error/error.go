package main

import (
	"errors"
	"fmt"
)

func add(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
}
func main() {
	x, y := 1, 2
	z, err := add(&x, &y)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(z)
		return
	}
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
