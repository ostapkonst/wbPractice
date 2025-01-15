// 11. Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

// Наивная реализация, лучше реализовать через hash-таблицу

func intersection(a, b []int) []int {
	result := make([]int, 0)

	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				result = append(result, v)
			}
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	c := intersection(a, b)
	fmt.Println(c)
}
