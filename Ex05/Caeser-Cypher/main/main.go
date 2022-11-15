package main

import (
	"fmt"
	"unicode"
)

func rotateBase(char rune, base, delta int) rune {
	tmp := int(char) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

//11
// middle-Outz
// 2

func main() {
	var len int
	var str string
	var spin int
	fmt.Scanf("%d\n", &len)
	fmt.Scanf("%s\n", &str)
	fmt.Scanf("%d\n", &spin)
	for i, c := range str {
		if unicode.IsUpper(c) {
			str = str[:i] + string(rotateBase(c, 'A', spin)) + str[i+1:]
		} else if unicode.IsLower(c) {
			str = str[:i] + string(rotateBase(c, 'a', spin)) + str[i+1:]
		}
	}
	fmt.Printf("%s\n", str)
}
