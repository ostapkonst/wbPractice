// 25. Реализовать собственную функцию sleep.
package main

import (
	"context"
	"fmt"
	"time"
)

// Реализация на основе контекстов
func mySleep(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	}
}

// Реализация без контекстов
func mySleepV2(t time.Duration) {
	select {
	case <-time.After(t):
		return
	}
}

func main() {
	defer timeout()()

	mySleep(5 * time.Second)
	mySleepV2(5 * time.Second)
}

func timeout() func() {
	now := time.Now()

	return func() {
		fmt.Printf("Elapsed time: %v", time.Since(now))
	}
}
