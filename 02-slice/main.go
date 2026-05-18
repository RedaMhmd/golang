package main

import (
	"fmt"
)

func main() {
	fmt.Println("slices")
	x := []int{1, 3, 5}
	fmt.Println(x)
	fmt.Println(cap(x))
	xx := []int(x[0:1])
	fmt.Println(cap(xx))
	xx = append(xx, 100)
	for i := 0; i < 2; i++ {
		fmt.Println("____")
		fmt.Println(xx[i])
		fmt.Println("____")
	}
	fmt.Println(cap(xx))

}
