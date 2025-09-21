package main

import (
	"fmt"
)

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45
	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%X %X\n", a, b)
	fmt.Printf("%#x %x\n", a, b)
	fmt.Printf("%f %.2f\n", c, d)

	fmt.Println()

	fmt.Printf("|%6d|%6d|\n", a, b) //use 6 columns
	fmt.Printf("|%06d|%06d|\n", a, b)
	fmt.Printf("|%-6d|%-6d|\n", a, b)
	fmt.Printf("|%6f|%6.2f|\n", c, d) //use at least 6 columns

}
