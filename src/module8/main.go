package main

import "fmt"

func main() {
	var input int
	var f int = 1
	fmt.Scanln(&input)

	for ; 0 < input; input-- {
		f = f * input
	}

	fmt.Println(f)
}
