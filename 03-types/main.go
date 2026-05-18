package main

import (
	"fmt"
)

func main() {
	x := new(int)
	// fmt.println()
	*x = 5
	print(*x)
	s := fmt.Sprintf("%T", x)

	if s == "*int" {
		print("true")
	}
}
