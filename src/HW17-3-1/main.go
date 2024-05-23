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

func main() {

	var counter int64 = 0
	var counterWG int = 0
	var wg sync.WaitGroup
	increment := func() {
		defer func() {
			counterWG--
			wg.Done()
		}()
		atomic.AddInt64(&counter, step)
	}
	// Не всегда вычисление этой переменной будет приводить к верному
	// результату в счётчике, но для правильных значений
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / step)
	for i := 1; i <= iterationCount; i++ {
		for counterWG > 10 {
			wg.Wait()
		}
		counterWG++
		wg.Add(1)
		go increment()
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
