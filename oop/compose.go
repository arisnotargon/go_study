package main

import (
	. "./type_assert"
	"fmt"
	"reflect"
)

func main() {
	//animal := NewAnimal("country dog")
	//pet := NewPet("pet dog")
	//dog := NewDog(&animal, pet)
	//fmt.Println(dog.GetAnimal().Call())

	var animal = NewAnimal("中华田园犬")
	var pet = NewPet("泰迪")
	var ianimal IAnimal = NewDog(&animal, pet)
	var ianimal2 IAnimal = NewAnimal("aname")

	// 类型断言
	if dog, ok := ianimal.(Animal); ok {
		fmt.Println(dog.GetName())
		fmt.Println(dog.Call())
		fmt.Println(dog.FavorFood())
	} else {
		fmt.Println("false =====")
	}

	if dog, ok := ianimal2.(Animal); ok {
		fmt.Println(dog.GetName())
		fmt.Println(dog.Call())
		fmt.Println(dog.FavorFood())
	} else {
		fmt.Println("false =====")
	}

	// 通过反射进行类型断言
	fmt.Println(reflect.TypeOf(ianimal))
	fmt.Println(reflect.TypeOf(ianimal2))
}
