// 22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел a и b
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений a и b (больше 2^20)
	a.SetString("123456789012345678901234567890", 10) // Пример значения для a
	b.SetString("987654321098765432109876543210", 10) // Пример значения для b

	// Сложение
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	// Вычитание
	diff := new(big.Int)
	diff.Sub(a, b)
	fmt.Printf("Разность: %s\n", diff.String())

	// Умножение
	prod := new(big.Int)
	prod.Mul(a, b)
	fmt.Printf("Произведение: %s\n", prod.String())

	// Деление
	quotient := new(big.Int)
	quotient.Div(a, b)
	fmt.Printf("Частное: %s\n", quotient.String())
}
