package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time

	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2", v)
		default:
			t = <-ticker.C
			outputMessage := []byte("Время: ")
			outputMessage = t.AppendFormat(outputMessage, "15:04:05")
			fmt.Println(string(outputMessage))
		}
	}
}

//Напишите код, в котором имеются два канала сообщений из целых чисел так, чтобы приём сообщений всегда приводил
//к блокировке. Приёмом сообщений из обоих каналов будет заниматься главная горутина. Сделайте так, чтобы во время
//такого «бесконечного ожидания» сообщений выполнялась фоновая работа в виде вывода текущего времени в консоль.
