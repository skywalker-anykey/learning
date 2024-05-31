package main

import (
	"08/8/2/hw8_8_2"
	"fmt"
)

func main() {
	bmwDimensions := hw8_8_2.NewMDimensions(4886, 1938, 1745)
	autoBMW := hw8_8_2.NewBMW("X5", 250, 455, bmwDimensions)
	printAuto(autoBMW)

	mercedesDimensions := hw8_8_2.NewMDimensions(5160, 1855, 1444)
	autoMercedes := hw8_8_2.NewMercedes("w220", 250, 306, mercedesDimensions)
	printAuto(autoMercedes)

	dodgeDimensions := hw8_8_2.NewDDimensions(5027, 1923, 1460)
	autoDodge := hw8_8_2.NewDodge("Challenger", 260, 372, dodgeDimensions)
	printAuto(autoDodge)

}

func printAuto(a hw8_8_2.Auto) {
	fmt.Println("Brand: ", a.Brand())
	fmt.Println("Model: ", a.Model())
	fmt.Println("Max Speed: ", a.MaxSpeed())
	fmt.Println("Engine Power: ", a.EnginePower())

	fmt.Print("Dimensions: ")
	switch p := a.(type) {
	case hw8_8_2.Dodge:
		fmt.Println(p.Dimensions().Length().Get(hw8_8_2.Inch), "x", p.Dimensions().Width().Get(hw8_8_2.Inch), "x", p.Dimensions().Height().Get(hw8_8_2.Inch))
	case hw8_8_2.Bmw, hw8_8_2.Mercedes:
		fmt.Println(p.Dimensions().Length().Get(hw8_8_2.CM), "x", p.Dimensions().Width().Get(hw8_8_2.CM), "x", p.Dimensions().Height().Get(hw8_8_2.CM))
	default:
		fmt.Println("непредвиденный тип")
	}
	fmt.Println()
}
