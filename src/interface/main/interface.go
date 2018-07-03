package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
	food string
}

type Men interface {
	Say()
	Eat()
	Sleep()
}

func (h1 Human) Say() {
	fmt.Println(h1.name)
}

func (h1 Human) Eat() {
	fmt.Println("eating " + h1.food)
}

func main() {
	h1 := Human{"bob", 22, "fruit"}
	h1.Say()
	h1.Eat()
}
