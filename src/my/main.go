package main

import "fmt"

func main() {
	var num int64
	num = 7

	fmt.Println(FactRecursive(num))
}

func FactRecursive(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * FactRecursive(n-1)
}
