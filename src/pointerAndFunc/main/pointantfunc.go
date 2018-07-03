package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student struct {
	person
	number int
}

func (p person) setPerson(name string) {
	p.name = name
}

func (p *person) setPointPerson(name string) {
	p.name = name
}

func (p person) getPersonName() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func (p *person) getPointPersonNmae() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}
func main() {
	p1 := person{"bob", 12}
	p1.getPersonName()
	p1.getPointPersonNmae()

	p1.setPerson("personname")
	p1.getPersonName()
	p1.getPointPersonNmae()

	p1.setPointPerson("pointname")
	p1.getPersonName()
	p1.getPointPersonNmae()

	s1 := student{p1, 1223}
	s1.getPersonName()

}
