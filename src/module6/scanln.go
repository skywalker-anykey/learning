package main

import (
	"fmt"

	"module6/calc"
)

func main() {
	var x, y float64
	var operation string
	fmt.Scanln(&x)
	fmt.Scanln(&operation)
	fmt.Scanln(&y)

	calc.Calc()

	switch operation {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Ошибка: попытка деления на ноль")
			break
		}
		fmt.Println(x / y)
	default:
		fmt.Println("Ошибка: непредвиденная операция")
	}
}
