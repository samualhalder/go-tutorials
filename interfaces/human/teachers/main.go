package main

import (
	"fmt"

	"github.com/samualhalder/go-tutorials/interfaces/human"
)

type Teacher struct {
	Age    int
	Name   string
	Salary int
}

func (t *Teacher) GetAge() int {
	return t.Age
}
func (t *Teacher) GetName() string {
	return t.Name
}
func (t *Teacher) IsRich() bool {
	return t.Salary > 50000
}

type createrTeacher struct {
	NewTeacher human.Human
}

func main() {
	teacher := &Teacher{Age: 24, Name: "Samual", Salary: 30000}
	newTeacher := createrTeacher{
		NewTeacher: teacher,
	}
	fmt.Print(newTeacher.NewTeacher.GetAge())
}
