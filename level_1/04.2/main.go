// 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// В данной реализации мы останавливаем продьюсер, а не воркеры. Получилось компактнее.

const (
	countOfWorker = 5
	countOfJobs   = 21
)

func generateSlice() []int {
	result := make([]int, countOfJobs)

	for i := 0; i < countOfJobs; i++ {
		result[i] = rand.Int()
	}

	return result
}

func payload(v int) {
	// эмуляция нагрузки
	time.Sleep(time.Second * 3)
	fmt.Println("payload:", v)
}

func main() {
	defer timeout()()

	done := make(chan os.Signal, 1)
	ch := make(chan int)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < countOfWorker; i++ {
		go func() {
			for v := range ch {
				payload(v)
			}
		}()
	}

	sliceWithNumbers := generateSlice()
	fmt.Println("length of slice with numbers:", len(sliceWithNumbers))

	for _, v := range sliceWithNumbers {
		select {
		case <-done:
			log.Println("service was stopped with Ctrl+C")
			close(ch)

			return

		default:
			ch <- v
		}
	}
}

func timeout() func() {
	now := time.Now()

	return func() {
		elapsedTime := time.Since(now).Seconds()

		time.Sleep(time.Millisecond)
		fmt.Printf("timeout after %f seconds\n", elapsedTime)
	}
}
