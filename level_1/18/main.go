// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync"
)

type Incrementor struct {
	value int
	mu    *sync.Mutex
}

func NewIncrementor() *Incrementor {
	return &Incrementor{
		value: 0,
		mu:    &sync.Mutex{},
	}
}

func (i *Incrementor) Increment() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.value++
}

func (i *Incrementor) Value() int {
	// можно было бы заюзать RWMutex, но тут и так пойдет
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.value
}

func main() {
	inc := NewIncrementor()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			inc.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("result after 1000 increments: %d\n", inc.Value())
}
