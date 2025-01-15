// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import "fmt"

// Полезная статья: https://acmp.ru/article.asp?id_text=42846

func partition(a []int, left, right int) int {
	return a[(left+right)/2] // берем центральный эелемент
}

func quickSort(a []int, left, right int) {
	if left >= right {
		// закончили сортировку
		return
	}

	pivot := partition(a, left, right)
	i, j := left, right

	for i <= j {
		for a[i] < pivot {
			i++
		}

		for a[j] > pivot {
			j--
		}

		// выше мы ищем элементы, которые можно поменять местами
		if i <= j {
			// вот такой у нас swap
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	quickSort(a, left, j)
	quickSort(a, i, right)
}

func sort(a []int) []int {
	// вот так мы копируем слайсы, нужно копировать
	b := make([]int, len(a))
	copy(b, a)

	quickSort(b, 0, len(b)-1)

	return b
}

func main() {
	a := []int{1, 4, 2, 3, 5, 5, 6, 0, -1}

	fmt.Printf("source array: %v\n", a)
	fmt.Printf("sorted array: %v\n", sort(a))
}
