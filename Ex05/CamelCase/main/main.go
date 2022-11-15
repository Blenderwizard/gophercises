package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	i := 1
	for _, c := range str {
		if unicode.IsUpper(c) {
			i++
		}
	}
	fmt.Printf("%d\n", i)
}
