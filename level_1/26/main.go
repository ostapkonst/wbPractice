// 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"unicode"
)

func isAllSymbolsUnic(s string) bool {
	m := make(map[rune]struct{})

	for _, r := range s {
		rs := unicode.ToUpper(r)
		if _, ok := m[rs]; ok {
			return false
		}
		m[rs] = struct{}{}
	}

	return true
}

func main() {
	a := "abcd"
	b := "abCdefAaf "

	fmt.Printf("str: %v, unic: %v\n", a, isAllSymbolsUnic(a))
	fmt.Printf("str: %v, unic: %v", b, isAllSymbolsUnic(b))
}
