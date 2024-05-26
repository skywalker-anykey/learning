package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	// Шаг наращивания счётчика
	step int64 = 1
	// Количество итераций на горутину
	counterGoroutine = 100
	// Количество горутин
	numGoroutines = 10
)

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < counterGoroutine; i++ {
				atomic.AddInt64(&counter, step)
			}
		}()
	}

	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
