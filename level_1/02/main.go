// 2. Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

func power2(n []int) {
	wg := &sync.WaitGroup{}

	for _, v := range n {
		wg.Add(1)

		go func() {
			fmt.Println(v * v)
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	power2(numbers)
}
