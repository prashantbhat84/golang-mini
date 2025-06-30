package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a int32
	var b bool
	var c uint32
	var result, remainder int
	d := '%'
	const myName = "Prashant Bhat"

	fmt.Println((a))
	// fmt.Println("Size of int:", unsafe.Sizeof(a))
	printMessage(fmt.Sprintf("Size of int: %d", unsafe.Sizeof(a)))
	printMessage(fmt.Sprintf("Size of bool: %d", unsafe.Sizeof(b)))
	printMessage(fmt.Sprintf("Size of uint32: %d", unsafe.Sizeof(c)))
	utf8Rune := utf8.RuneCountInString(string(d))
	printMessage(fmt.Sprintf("Size of rune: %d", utf8Rune))
	printMessage(fmt.Sprintf("My name is: %s", myName))
	printMessage(fmt.Sprintf("Size of string: %d", unsafe.Sizeof(myName)))
	result, remainder, err := intDivision(11, 0)
	if err != nil {
		printMessage(err.Error())

	}
	printMessage(fmt.Sprintf("Division Result: %d", result))
	printMessage(fmt.Sprintf("Remainder: %d", remainder))
	printMessage("End of program")

}
func printMessage(value string) {
	fmt.Println(value)
}
func intDivision(a, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero is invalid")
		return 0, 0, err
	}
	result := a / b
	remainder := a % b
	return result, remainder, err
}
