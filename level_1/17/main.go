// 17. Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"log"
	"sort"
)

func binarySearch(a []int, x int) bool {
	if len(a) == 0 {
		return false
	}

	middle := len(a) / 2

	log.Println("a:", a, "x:", x, "middle:", middle)

	if a[middle] == x {
		return true
	}

	if a[middle] < x {
		return binarySearch(a[middle+1:], x)
	} else {
		return binarySearch(a[:middle], x)
	}
}

func main() {
	a := []int{4, 3, 2, 1, 6}
	sort.Ints(a)

	// срез должен быть предварительно отсортирован
	result := binarySearch(a, -1)

	if result {
		println("found")
	} else {
		println("not found")
	}
}
