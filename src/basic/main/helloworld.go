package main

import (
	"fmt"
)

var array = [5]int{1, 2, 3}
var aslice []int

func main() {
	fmt.Println("helloworld")
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 10:
		fmt.Println("i is equal to 10")
		fallthrough
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")

	default:
		fmt.Println("All I know is that i is an integer")
	}
	fmt.Println(max(10, 23))
}

func max(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}
