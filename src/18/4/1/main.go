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
	s := NewSemaphore(10, 3*time.Second)

	if err := s.Acquire(); err != nil {
		fmt.Println(err)
	}

	if err := s.Release(); err != nil {
		fmt.Println(err)
	}
}
