package main

import (
	"fmt"
	"sync"
)

func receive(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	// Это тоже работает
	//for _, ok := <-ch; ok; _, ok = <-ch {
	//	select {
	//	case <-ch:
	//	}
	//}
	for range ch {
	}
	return
}

func main() {
	// Количество отправляемых сообщений в каналы
	const countSend = 1000

	var countC1, countC2 int32
	wg := sync.WaitGroup{}

	c1 := make(chan int, 1000)
	c2 := make(chan int, 1000)

	go receive(c1, &wg)
	go receive(c2, &wg)

	for i := 0; i < countSend; i++ {
		select {
		case c1 <- i:
			countC1++
		case c1 <- i:
			countC1++
		case c2 <- i:
			countC2++
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
