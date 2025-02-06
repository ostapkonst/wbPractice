// 23. Удалить i-ый элемент из слайса.
package main

import (
	"fmt"
)

// Сохраняем порядок элементов
func remOneElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// Игнорируем порядок элементов
func remOneElementV2(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(remOneElement(a, 3))

	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(remOneElementV2(b, 3))
}
