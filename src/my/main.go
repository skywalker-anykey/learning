package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3}
	var index int

	fmt.Scan(&index)

	if index >= len(numbers) || index < 0 {
		fmt.Println("prevented!")
	} else {
		fmt.Println(numbers[index])
	}
}
