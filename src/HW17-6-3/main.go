package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Wait()
}

//Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал,
//вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).
