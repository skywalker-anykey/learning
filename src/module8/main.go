package main

import (
	"fmt"
	"module8/electronic"
)

func main() {
	apple := electronic.NewApplePhone("iPhone15", "iOS 17.0")
	android := electronic.NewAndroidPhone("POCO", "X6 Pro", "Android 15")
	radio := electronic.NewRadioPhone("Panasonic", "T1000", 10)

	printCharacteristics(&apple)
	printCharacteristics(&android)
	printCharacteristics(&radio)
}

// printCharacteristics должна выводить бренд, модель и тип телефона.
func printCharacteristics(p electronic.Phone) {
	fmt.Println("Brand:\t", p.Brand())
	fmt.Println("Model:\t", p.Model())
	fmt.Println("Type:\t", p.Type())
	switch p.Type() {
	case "station":
		fmt.Println("ButtonsCount:\t", p.ButtonsCount())
	case "smartphone":
		fmt.Println("OS:\t", p.OS())
	}
}
