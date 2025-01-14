// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func power2(n int) int {
	return n * n
}

func sum(n []int, op func(int) int) int {
	// написано из соображения, что операция `op` долгая.
	// суммирование в отдельной горутине чисто по фану.

	accumulator := make(chan int)
	wg := &sync.WaitGroup{}
	done := make(chan struct{})
	result := 0

	go func() {
		for v := range accumulator {
			result += v
		}

		done <- struct{}{}
	}()

	for _, v := range n {
		wg.Add(1)

		go func() {
			accumulator <- op(v)

			wg.Done()
		}()
	}

	wg.Wait()

	close(accumulator)

	<-done

	return result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(sum(numbers, power2))
}
