package main

import (
	"16/6/2/bank"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	client := bank.New()

	//Используя этот клиент, создайте консольное приложение, которое:
	//В момент старта запускает 10 горутин, каждая из которых с промежутком от 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.
	for i := 0; i < 10; i++ {
		go func() {
			for {
				cur := rand.IntN(10) + 1
				client.Deposit(cur)
				time.Sleep(time.Microsecond * time.Duration(rand.IntN(500)+500))
			}
		}()
	}

	//Так же запускается 5 горутин, которые с промежутком 0.5 секунд до 1 секунды снимают с клиента случайную сумму от 1 до 5.
	//Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
	for i := 0; i < 5; i++ {
		go func() {
			for {
				cur := rand.IntN(5) + 1
				err := client.Withdrawal(cur)
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(time.Microsecond * time.Duration(rand.IntN(500)+500))
			}
		}()
	}

	//Если пользователь введет в консоль слово balance — приложение выведет в консоль текущий баланс клиента.
	//
	//deposit — запрашивается сумма (целое число) — которая добавляется на счёт пользователя
	//
	//withdrawal — запрашивается сумма (целое число) — которая выводится со счёта пользователя.
	//Если запрашиваемая сумма больше текущего баланса пользователя, то пользователю должно быть показано сообщение о том,
	//что его баланс недостаточен (и, очевидно, операция не должна быть выполнена).
	//
	//Если пользователь введет слово exit — приложение завершит работу.
	//
	//При вводе любой другой команды приложение выведет сообщение "Unsupported command. You can use commands: balance, deposit, withdrawal, exit".
	for {
		switch ReadCMD() {
		case "deposit":
			client.Deposit(ReadAmount())
		case "withdraw":
			err := client.Withdrawal(ReadAmount())
			if err != nil {
				fmt.Println(err)
			}
		case "balance":
			fmt.Println(client.Balance())
		case "exit":
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}
}

func ReadCMD() string {
	var s string
	fmt.Print("Enter command: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		//log.Println(err)
		return ""
	}
	return s
}

func ReadAmount() int {
	var a int
	fmt.Print("Enter amount: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		//log.Println(err)
		return 0
	}
	return a
}
