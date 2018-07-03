package main

import (
	"fmt"
)

type rectangle struct {
	width, heigh float64
}

func main() {
	c1 := circle{10}
	fmt.Println(c1.area())
}
