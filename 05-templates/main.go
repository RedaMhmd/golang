package main

import "fmt"

func main() {

	fmt.Println("hello")
	x := 1
	switch x {
	case 10:
		fmt.Println("10")
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("def")
	}

}
