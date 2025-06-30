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
	// intArray := [8]int{1, 2, 3, 4, 5}
	var intSlice = []int{1, 2, 3, 4}
	var intSlice2 = []int{5, 6}
	printMessage("Start of program")
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
	fmt.Println("before append")
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice2)
	fmt.Println((intSlice))

	fmt.Println("after append")
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice)
	fmt.Println(intSlice2)
	//using the make fn to define the slice
	intSlice3 := make([]int, 5, 10)
	intSlice3[0] = 1
	fmt.Println("intSlice3 before append")
	fmt.Println(intSlice3)
	fmt.Println("MAPS")
	var mymap = make(map[string]int)
	var mymap2 = map[string]int{"count": 1, "population": 2, "deleteobject": 3}
	mymap["a"] = 1
	fmt.Println(mymap)
	fmt.Println(mymap2)
	// delete an key from  map
	delete(mymap2, "deleteobject")
	mymap2["count"] = 3
	fmt.Println(mymap2)
	fmt.Println("Accessing map values and checking existence")
	var count2, ok = mymap2["count2"]
	if ok {
		fmt.Println("Count2 exists:", count2)
	} else {
		fmt.Println("Count2 does not exist")
	}
	// iterating over  a map and slice
	for key, value := range mymap2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	for i, value := range intSlice {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
	// while loop variation using for as go does not have a while loop
	var i = 0
	for i < 5 {
		fmt.Println("without 3 parts")
		printMessage(fmt.Sprintf("Value of i: %d", i))
		i++
	}
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("with 3 parts")
		printMessage(fmt.Sprintf("Value of i: %d", i))
	}

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
