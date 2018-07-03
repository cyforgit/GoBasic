package main

import (
	"fmt"
)

var user string = "a"

func main() {
	fmt.Println(defertest())
}
func defertest() int {

	fmt.Println("before defer")
	defer fmt.Println("defer")
	fmt.Println("after defer")
	return 1
}
func init() {
	if user == "" {
		panic("panic ")
	}
}
