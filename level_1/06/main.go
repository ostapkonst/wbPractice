// 6. Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup // используется для ожидания завершения всех горутин

func payload(s string) {
	// нагрузка будет занимать 2 секунды
	time.Sleep(time.Second * 2)
	fmt.Println(s)
}

func cancelByContext(ctx context.Context) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			payload("cancelByContext")
		}
	}
}

func cancelByChannel(channel <-chan struct{}) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-channel:
			return
		default:
			payload("cancelByChannel")
		}
	}
}

func cancelByCancelFunc() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	defer wg.Done()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				payload("cancelByCancelFunc")
			}
		}
	}()

	return cancel
}

func cancelByTimeout(duration time.Duration) {
	timeout := time.After(duration)

	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-timeout:
			return
		default:
			payload("cancelByTimeout")
		}
	}
}

func cancelByAtomicBool(atomicBool *atomic.Bool) {
	wg.Add(1)
	defer wg.Done()

	for {
		if atomicBool.Load() {
			return
		}

		payload("cancelByAtomicBool")
	}
}

func main() {
	// отмена по контексту
	ctx, cancel := context.WithCancel(context.Background())
	go cancelByContext(ctx)
	time.Sleep(time.Second)
	cancel()

	// отмена по каналу
	channel := make(chan struct{})
	go cancelByChannel(channel)
	time.Sleep(time.Second)
	channel <- struct{}{}

	// отмена по функции отмены
	cancel = cancelByCancelFunc()

	time.Sleep(time.Second)
	cancel()

	// отмена по timeout
	go cancelByTimeout(time.Second)

	// отмена по atomicBool
	atomicBool := new(atomic.Bool)
	go cancelByAtomicBool(atomicBool)

	time.Sleep(time.Second)
	atomicBool.Store(true)

	// протестировали все способы
	wg.Wait()
}
