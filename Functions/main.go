package main

import (
	"fmt"
)

func do(b []int) int {
	b[0] = 0
	return b[1]
}
func main() {
	a := []int{1, 2, 3} //slice
	v := do(a)

	fmt.Println(a, v)
}
