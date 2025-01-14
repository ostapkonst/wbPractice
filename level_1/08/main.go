// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func toBits(n int64) string {
	var s string

	for n > 0 {
		s = fmt.Sprintf("%d", n%2) + s
		n /= 2
	}

	if s == "" {
		return "0"
	}

	return s
}

func setBit(n, i int64) int64 {
	bit := (int64(n) >> i) % 2 // какой бит под номером i
	if bit == 0 {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func printValue(source, bit int64) {
	fmt.Println(">>>")
	fmt.Println("\tsource:", toBits(source), "\n\tbit to change:", bit, "\n\tmodified:", toBits(setBit(source, bit)))
	fmt.Println("<<<\n")
}

func main() {
	printValue(0, 3)
	printValue(6, 2)  // перевод 1 в 0
	printValue(10, 2) // перевод 0 в 1
}
