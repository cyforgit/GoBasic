//package main
//
//
//func test(c chan int) {
//	for i := 0; i < 10; i++ {
//		println("send ", i)
//		c <- i
//	}
//}
//
//func main() {
//	ch := make(chan int ,0)
//	go test(ch)
//	for j := 0; j < 10; j++ {
//		println("get ", <-ch)
//	}
//}
