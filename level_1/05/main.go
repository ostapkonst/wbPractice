// 5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const timeout = time.Second * 5 // через сколько завершится программа

func publisher(ch chan<- int) {
	for {
		ch <- rand.Int()
	}
}

func subscriber(ch <-chan int) {
	// отмена не всегда отрабатываем вовремя, тут все зависит от того, с каким приоритетом берутся значения из каналов
	timer := time.After(timeout)

	for {
		select {
		case <-timer:
			return
		case v := <-ch:
			time.Sleep(time.Second)
			fmt.Println(v)
		}
	}
}

func main() {
	ch := make(chan int)
	go publisher(ch)
	subscriber(ch)
}
