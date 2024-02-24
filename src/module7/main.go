package main

import (
	"fmt"
	"log"
	"module7/calc"
)

func main() {
	// Внутри функции main необходимо из консоли считать 2 числа и оператор (как вы это делали в задании предыдущего модуля).
	var x, y float64
	var operation string
	if _, err := fmt.Scanln(&x); err != nil {
		log.Panic(err)
	}
	if _, err := fmt.Scanln(&operation); err != nil {
		log.Panic(err)
	}
	if _, err := fmt.Scanln(&y); err != nil {
		log.Panic(err)
	}

	fmt.Printf("\n%+v %+v %+v =\n", x, operation, y)

	// Затем создать экземпляр структуры calculator из пакета calc и вызвать метод Calculate, передав ему полученные из консоли значения.
	c := calc.NewCalulator()
	result := c.Calculate(x, y, operation)

	// Полученный из метода Calculate результат нужно распечатать в консоль.
	fmt.Println(result)
}
