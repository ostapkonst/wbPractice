// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func sliceToSet(a []string) map[string]bool {
	result := make(map[string]bool)

	for _, v := range a {
		result[v] = true
	}

	return result
}

func setToSlice(a map[string]bool) []string {
	result := make([]string, 0)

	for k := range a {
		result = append(result, k)
	}

	return result
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(setToSlice(sliceToSet(a)))
}
