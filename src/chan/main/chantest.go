package main

import (
	"fmt"
	"html"
	"log"
)

func main() {
	cin := make(chan int)
	go test(cin)
	for i := 0; i < 10; i++ {
		cin <- i
	}
}

func test(c chan int) {
	for temp := range c {
		fmt.Println(temp)
	}
}
