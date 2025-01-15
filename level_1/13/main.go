// 13. Поменять местами два числа без создания временной переменной.
package main

import "fmt"

// Реализовано оба варианта из статьи: https://www.programbeginner.ru/?p=99

func swap1(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b

	return a, b
}

func swap2(a, b int) (int, int) {
	a += b
	b = a - b
	a -= b

	return a, b
}

func main() {
	a := 5
	b := 7

	r1, r2 := swap1(a, b)
	r3, r4 := swap2(a, b)

	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Printf("a: %d, b: %d\n", r1, r2)
	fmt.Printf("a: %d, b: %d\n", r3, r4)
}
