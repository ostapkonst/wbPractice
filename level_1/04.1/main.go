// 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
package main

// Задачу можно было реализовать проще: через отмену записи в канал при Сtrl+C.
// В данной реализации мы останавливаем именно воркеры, а можно было и продьюсер остановить.

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"log"
)

const (
	countOfWorker = 5
	countOfJobs   = 11
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	done := &sync.WaitGroup{}

	done.Add(countOfWorker)

	for i := 0; i < countOfWorker; i++ {
		go func() {
			defer done.Done()

			for {
				select {
				case <-ctx.Done():
					log.Printf("worker %d: context was marked as done earlier, than user service has stopped\n", i)
					return
				case v, ok := <-ch:
					if !ok {
						log.Printf("worker %d: channel was closed, all work done\n", i)
						return
					}

					payload(v)
					wg.Done()
				}
			}
		}()
	}

	sliceWithNumbers := generateSlice()
	fmt.Println("length of slice with numbers:", len(sliceWithNumbers))

	go func() {
		defer stop()

		for _, randomNumber := range sliceWithNumbers {
			wg.Add(1)
			ch <- randomNumber
		}

		wg.Wait() // ожидаем что все воркеры завершили работу

		close(ch) // закрываем канал и сообщаем всем воркерам, что работа закончена
	}()

	<-ctx.Done() // ожидаем завершения работы программы

	done.Wait() // все воркеры завершились
}

func timeout() func() {
	now := time.Now()

	return func() {
		elapsedTime := time.Since(now).Seconds()

		time.Sleep(time.Millisecond)
		fmt.Printf("timeout after %f seconds\n", elapsedTime)
	}
}
