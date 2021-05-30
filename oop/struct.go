package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

func (s *Student) GetName() string {
	fmt.Println("in get name===>", s)
	return s.name
}

func (s *Student) setName(newName string) {
	s.name = newName
}

func (s *Student) String() string {
	return fmt.Sprintf("id->%d,name->%s,male->%t,score->%f", s.id, s.name, s.male, s.score)
}

func main() {
	s := NewStudent(1, "ho", 10)

	fmt.Println(s.GetName())
	s.setName("hyahyahya")
	fmt.Println(s.GetName())
	fmt.Println(s)
}
