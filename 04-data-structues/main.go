package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("hello from data structures")
	// nf, _ := os.OpenFile("cash.gofile", os.O_RDWR, os.ModeAppend)
	// nf.Seek(20, io.SeekStart)
	// for i := range 1000 {
	// 	fmt.Fprintf(nf, "%-[2]*[1]d\n", i, 3)

	// }

	s := "reda mhmd"

	str := "H世界"

	for i, v := range s {
		fmt.Printf("%d, %c \n", s[i], v)
	}
	for _, v := range str {
		fmt.Printf("%v \n", v)
		// fmt.Print(v)
		fmt.Println(utf8.RuneLen(v))
	}
	println(len(str))

}
func init() {
	fmt.Println("reda")
}
