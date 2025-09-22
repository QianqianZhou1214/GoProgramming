package main

import (
	"fmt"
)

func do(b []int) int {
	b[0] = 0
	fmt.Printf("b@ %p\n", b)
	return b[1]
}
func main() {
	a := []int{1, 2, 3} //slice
	fmt.Printf("a@ %p\n", a)
	v := do(a)

	fmt.Println(a, v)
}
