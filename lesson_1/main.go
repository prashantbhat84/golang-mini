package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a int32
	var b bool
	var c uint32
	d := '%'
	const myName = "Prashant Bhat"

	fmt.Println((a))
	fmt.Println("Size of int:", unsafe.Sizeof(a))
	fmt.Println("Size of bool:", unsafe.Sizeof(b))
	fmt.Println("Size of uint:", unsafe.Sizeof(c))
	utf8Rune := utf8.RuneCountInString(string(d))
	fmt.Println("Size of rune:", utf8Rune)
	fmt.Println("Size of rune:", unsafe.Sizeof(d))
	fmt.Println("Size of string:", unsafe.Sizeof(myName))
	fmt.Println("Size of string:", len(myName))

}
