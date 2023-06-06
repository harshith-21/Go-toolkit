package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}

func sum(ints ...int) int {
	var sum int
	for x := range ints {
		sum += x
	}
	return sum
}
