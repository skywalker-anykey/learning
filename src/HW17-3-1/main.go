package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

// Количество горутин
const numGoroutines = 10

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup
	increment := func() {
		defer wg.Done()
		for {
			if counter >= endCounterValue {
				return
			}
			atomic.AddInt64(&counter, step)
		}
	}
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go increment()
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
