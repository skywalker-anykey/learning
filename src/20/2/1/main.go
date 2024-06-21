package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// Напишите код, реализующий пайплайн, работающий с целыми числами и состоящий из следующих стадий:
	pipe := bufferStage(filterStage2(filterStage1(generator())))
	wg.Add(1)
	reporter(&wg, pipe)
	wg.Wait()
}

// Для тестирования без ручного ввода
func generatorAuto() chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for i := -3; i < 15; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("send:", i)
			output <- i
		}
	}()
	return output
}

// Написать источник данных для конвейера. Непосредственным источником данных должна быть консоль
func generator() chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		fmt.Println("Программа принимает целые числа. Для выхода [Enter].")
		scanner := bufio.NewScanner(os.Stdin)
		var s string
		for {
			scanner.Scan()
			s = scanner.Text()
			if s == "" {
				fmt.Println("Выход. Получаем последние результаты.")
				return
			}
			// При написании источника данных подумайте о фильтрации нечисловых данных, которые можно ввести через консоль.
			// Как и где их фильтровать, решайте сами.
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("введено не число")
				continue
			}
			output <- i
		}
	}()
	return output
}

// Также написать код потребителя данных конвейера. Данные от конвейера можно направить снова в консоль построчно,
// сопроводив их каким-нибудь поясняющим текстом, например: «Получены данные ...»
func reporter(wg *sync.WaitGroup, input <-chan int) {

	go func() {
		defer wg.Done()
		for value := range input {
			fmt.Println("Получены данные: ", value)
		}
	}()

}

// Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа)
func filterStage1(input <-chan int) chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
			case value, ok := <-input:
				// Если канал закрыт, то данных больше не будет и сигнал к завершению работы рутины
				if !ok {
					// Закрываем следующий канал, чтобы оповестить следующую рутину о завершении
					return
				}
				// Обработка полученного значения и передача следующей рутине
				if value >= 0 {
					output <- value
				}
			}
		}
	}()

	return output
}

// Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0
func filterStage2(input <-chan int) chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
			case value, ok := <-input:
				// Если канал закрыт, то данных больше не будет и сигнал к завершению работы рутины
				if !ok {
					// Закрываем следующий канал, чтобы оповестить следующую рутину о завершении
					return
				}
				// Обработка полученного значения и передача следующей рутине
				if value%3 == 0 && value != 0 {
					output <- value
				}
			}
		}
	}()

	return output
}

// Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому, который был дан в качестве задания в 19 модуле.
// В этой стадии предусмотреть опустошение буфера (и соответственно, передачу этих данных, если они есть, дальше) с определённым интервалом
// во времени.
func bufferStage(input <-chan int) chan int {
	// Значения размера буфера и этого интервала времени сделать настраиваемыми (как мы делали: через константы или глобальные переменные).
	const (
		bufferSize    = 10
		bufferGetTime = time.Second * 15
	)
	output := make(chan int, bufferSize)
	r := ring.New(bufferSize)
	end := false
	m := sync.Mutex{}

	go func() {
		for {
			select {
			case value, ok := <-input:
				// Если канал закрыт, то данных больше не будет и сигнал к завершению работы рутины
				if !ok {
					// Закрываем следующий канал, чтобы оповестить следующую рутину о завершении
					end = true
					return
				}
				m.Lock()
				r.Value = value
				r = r.Next()
				m.Unlock()
			}
		}
	}()

	go func() {
		defer close(output)
		for {
			select {
			case <-time.After(bufferGetTime):
				m.Lock()
				// Читаем и очищаем "буфер"
				for j := 0; j < bufferSize; j++ {
					switch r.Value.(type) {
					case int:
						output <- r.Value.(int)
					}
					r.Value = nil
					r = r.Next()
				}
				m.Unlock()
			}
			// Выход из рутины если предыдущая рутина закрылась
			if end {
				return
			}
		}
	}()

	return output
}
