// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Написана простейшая реализация требуемого функционала.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// пишет в первый канал случайные числа
	go func() {
		for {
			ch1 <- rand.Int() % 10

			time.Sleep(time.Second * 1)
		}
	}()

	// пишет во второй канал числа умноженные на 2
	go func() {
		for {
			ch2 <- <-ch1 * 2
		}
	}()

	// выводит числа из второго канала
	go func() {
		for {
			println(<-ch2)
		}
	}()

	<-done
}
