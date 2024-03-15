package main

import "fmt"

func main() {
	var f, num int
	f, num = 1, 5

	for ; 0 < num; num-- {
		f *= num
	}
	fmt.Println(f)
}
