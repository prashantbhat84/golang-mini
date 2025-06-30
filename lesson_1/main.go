package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = []rune("résumé")
	var indexed = str[1]
	fmt.Printf("%v , %T\n", indexed, indexed)
	for i, v := range str {
		// fmt.Printf("%d %c\n", i, v)
		fmt.Println(i, v)
	}
	fmt.Println("Length of string:", len(str))
	var str2 = []string{"s", "t", "r", "i", "n", "g"}
	var strBuilder strings.Builder

	for i := range str2 {

		strBuilder.WriteString(str2[i])
	}
	fmt.Println("Concatenated string:", strBuilder.String())

}
