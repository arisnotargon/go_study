package main

import (
	. "./animal"
	"fmt"
)

func main() {
	animal := NewAnimal("country dog")
	pet := NewPet("pet dog")
	dog := NewDog(&animal, pet)
	fmt.Println(dog.GetAnimal().Call())
}
