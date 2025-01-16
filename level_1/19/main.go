// 19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import "strings"

func reverse(s string) string {
	// наивная реализация
	result := ""

	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func reverseWithStringBuilder(s string) string {
	// реализация с использование StringBuilder

	var strBuilder strings.Builder

	// руны нам позволят нормально переворачивать строку
	sliceWithRunes := []rune(s)

	for i := len(sliceWithRunes) - 1; i >= 0; i-- {
		strBuilder.WriteRune(sliceWithRunes[i])
	}

	return strBuilder.String()
}

func main() {
	inputStr := "главрыба — абырвалг - 007"

	println(reverse(inputStr))
	println(reverseWithStringBuilder(inputStr))
}
