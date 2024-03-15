package main

import "fmt"

func main() {
	var f, num int
	f, num = 1, 5

	for i := 1; i <= num; i++ {
		f *= i
	}
	fmt.Println(f)
}
