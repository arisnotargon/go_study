package main

import (
	. "../type_assert"
	"fmt"
	"reflect"
)

//空接口可以指向任意变量类型

var v1 interface{} = 1       // 指向int
var v2 interface{} = "こんにちは" // 指向string
var v3 interface{} = true    // 指向bool

var v4 interface{} = &v2            // 指向指针类型
var v5 interface{} = []int{1, 2, 3} // 指向切片
var v6 interface{} = struct {
	id   int
	name string
}{
	1,
	"asdf",
} // 指向匿名构造体

func main() {
	var animal = NewAnimal("countrydog")
	var pet = NewPet("petdog")
	var any interface{} = NewDog(&animal, pet)
	// 利用空接口实现更加灵活的类型断言
	if dog, ok := any.(Dog); ok {
		fmt.Println(dog.GetName())
		fmt.Println("typeof any===>", reflect.TypeOf(dog))
		// 反射获取结构体信息，并且动态调用其成员方法,需要先获取对应的 reflect.Value 类型值：
		dogValue := reflect.ValueOf(&dog).Elem()
		fmt.Println(dogValue)
		// 根据类型值获取该类型的属性和方法
		// 遍历属性
		for i := 0; i < dogValue.NumField(); i++ {
			// 获取属性名
			fmt.Println("name:", dogValue.Type().Field(i).Name)
			// 获取属性类型
			fmt.Println("type:", dogValue.Type().Field(i).Type)
			// 获取属性值
			fmt.Println("value:", dogValue.Field(i))
		}
		// 遍历获取方法
		for i := 0; i < dogValue.NumField(); i++ {
			// 获取方法名
			fmt.Println("name:", dogValue.Type().Method(i).Name)
			// 获取方法详情
			fmt.Println("type:", dogValue.Type().Method(i).Type)
			// 获取方法值
			fmt.Println("value:", dogValue.Method(i))

			// call
			fmt.Println("call====>", dogValue.Method(i).Call([]reflect.Value{})[0])
		}
	}

}
