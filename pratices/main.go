package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello acen!\n")

	res, sub := add(7, 8)
	fmt.Printf("%d %d \n", res, sub)
	res2, sub2 := add2(7, 8)
	fmt.Printf("%d %d \n", res2, sub2)
}

func add(x, y int) (result, sub int) {
	result = x + y
	sub = x - y
	return
}

func add2(x, y int) (int, int) {
	result := x + y
	sub := x - y
	return result, sub
}
