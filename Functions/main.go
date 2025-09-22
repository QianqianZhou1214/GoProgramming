package main

import (
	"fmt"
)

// map is reference type
func do(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int) //reassigned to a new empty map
	m1[4] = 4              //local variable
	fmt.Println("m1", m1)
}
func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	do(m)

	fmt.Println("m", m)
}
