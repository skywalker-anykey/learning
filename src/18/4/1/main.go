// 18.4.1 Измените приведенный выше код по реализации двоичного семафора так, чтобы он описывал не двоичный семафор, а семафор подсчёта
package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	sem     chan int
	timeout time.Duration
}

func NewSemaphore(counter int, timeout time.Duration) *Semaphore {
	return &Semaphore{make(chan int, counter), timeout}
}

func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- 0:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("timeout. неудалось захватить семафор")
	}
}

func (s *Semaphore) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("timeout. неудалось освободить семафор")
	}
}

func main() {
	s := NewSemaphore(3, 3*time.Second)

	// Выполним 9 задач по 1 с работы каждая с ограничением 3 потока одновременно
	for i := 0; i < 9; i++ {
		go func(i int) {
			defer func() {
				if err := s.Release(); err != nil {
					fmt.Println(err)
				}
			}()
			if err := s.Acquire(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Work", i, "Start")
			time.Sleep(1 * time.Second)
			fmt.Println("Work", i, "Stop")
		}(i)
	}

	// Должны завершить работы примерно за 3 секунды
	time.Sleep(4 * time.Second)
}
