package main

import (
	"fmt"
)

func test(cin chan int, cout chan int) {
	for temp := range cin {
		println("cin get ", temp)
		cout <- temp
	}
}

func main() {
	cin := make(chan int)
	cout := make(chan int)
	go test(cin, cout)
	for i := 0; i <= 10; i++ {
		cin <- i
		var a int = <-cout
		fmt.Println(a)
	}
}
