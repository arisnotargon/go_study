package animal

type Animal struct {
	// 小写开头进行私有化
	name string
}

func NewAnimal(name string) Animal {
	return Animal{
		name: name,
	}
}

func (a Animal) Call() string {
	return "animal call..."
}

func (a Animal) FavorFood() string {
	return "animal favor food"
}

func (a Animal) GetName() string {
	return a.name
}
