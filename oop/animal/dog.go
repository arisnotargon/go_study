package animal

type Dog struct {
	animal *Animal
	pet    Pet
}

func NewDog(animal *Animal, pet Pet) Dog {
	return Dog{animal: animal, pet: pet}
}

func (d Dog) FavorFood() string {
	return "bone"
}

func (d Dog) Call() string {
	return "dog:wangwangwang"
}

func (d Dog) GetAnimal() Animal {
	return *d.animal
}
