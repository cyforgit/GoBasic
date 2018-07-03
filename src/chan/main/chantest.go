package main

import (
	"fmt"
	"strconv"
)

func fibonacci(c1, c2 chan int) {
	var a, b int
	for {
		select {
		case a = <-c1:
			fmt.Println("c1 recv " + strconv.Itoa(a))
		case b = <-c2:
			fmt.Println("c2 recv " + strconv.Itoa(b))
		default:
			fmt.Println("run default")
		}
	}

}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()
	c2 <- 0
	fibonacci(c1, c2)
	//	go func() {
	//		for i := 19; i < 99; i++ {
	//			c2 <- i
	//		}
	//	}()
}

//func say(s string) {
//	for i := 0; i < 6; i++ {
//		fmt.Println(s)
//		runtime.Gosched()
//	}
//}
//func main() {
//	go say("hello")
//	go say("world")
//	go say("go")
//	say("main")
//}

//func main() {
//	cin := make(chan int)
//	go test(cin)
//	for i := 0; i < 10; i++ {
//		cin <- i
//	}
//}
//
//func test(c chan int) {
//	for temp := range c {
//		fmt.Println(temp)
//	}
//}
