package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func receive(ch chan int, count *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for _, ok := <-ch; ok; _, ok = <-ch {
		atomic.AddInt32(count, 1)
	}
	return
}

func main() {
	// Количество отправляемых сообщений в каналы
	const countSend = 1000

	var countC1, countC2 int32
	wg := sync.WaitGroup{}

	c1 := make(chan int)
	c2 := make(chan int)

	// Вероятность "2 к 1" получаем за счет дополнительной горутины, обрабатывающей канал c1
	go receive(c1, &countC1, &wg)
	go receive(c1, &countC1, &wg)
	go receive(c2, &countC2, &wg)

	for i := 0; i < countSend; i++ {
		select {
		case c1 <- i:
			//fmt.Println("c1 <-", i)
		case c2 <- i:
			//fmt.Println("c2 <-", i)
		}
	}
	close(c1)
	close(c2)
	wg.Wait()
	fmt.Println("Принятых сообщений первым каналом:", countC1)
	fmt.Println("Принятых сообщений вторым каналом:", countC2)
}

//Напишите код, в котором имеются два канала сообщений из целых чисел, так,чтобы приём сообщений из них никогда
//не приводил к блокировке и чтобы вероятность приёма сообщения из первого канала была выше в 2 раза, чем из второго.
//
//*Если хотите, можете написать код, который бы демонстрировал это соотношение.
