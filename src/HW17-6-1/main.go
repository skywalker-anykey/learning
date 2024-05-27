package main

import (
	"fmt"
	"sync"
)

func receive(ch chan int, count *int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for range ch {
		*count++
	}
}

func main() {
	// Количество отправляемых сообщений в каналы
	const countSend = 100

	var countC1, countC2 int
	wg := sync.WaitGroup{}

	c1 := make(chan int)
	c2 := make(chan int)

	go receive(c1, &countC1, &wg)
	go receive(c1, &countC1, &wg)
	go receive(c2, &countC2, &wg)

	for i := 0; i < countSend; i++ {
		select {
		case c1 <- i:
		case c1 <- i:
		case c2 <- i:
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
