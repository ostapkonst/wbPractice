// 20. Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

// Переворачиваем слова за линейное время, используем строковый билдер, чтобы избежать лишних аллокаций

func reverseSentence(v string) string {
	if v == "" {
		return ""
	}

	strBuilder := strings.Builder{}

	j := len(v)

	for i := len(v) - 1; i >= 0; i-- {
		if v[i] == ' ' {
			strBuilder.WriteString(v[i+1:j] + " ")
			j = i
		}
	}

	strBuilder.WriteString(v[:j])

	return strBuilder.String()
}

func main() {
	inpString := "snow dog sun"
	fmt.Println(reverseSentence(inpString))
}
