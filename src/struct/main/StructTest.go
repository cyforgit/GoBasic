package main

import (
	"fmt"
)

const (
	white = iota
	black
	blue
	red
)

type person struct {
	name string
	age  int
}

func introduce(p person) {
	fmt.Println(p.name)
	fmt.Println(p.age)
	
}

func main() {
	p := person{"peter", 11}
	introduce(p)
}
