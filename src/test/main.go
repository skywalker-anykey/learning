package main

import "fmt"

func main() {
	var a, b, c int = 2, 3, 4
	fmt.Println(ret(&a, &b, &c))
	fmt.Println(ret(&a, &b, &c))
}

func ret(x, y, z *int) (int, int, int) {
	*x = *x * 2
	*y = *y * 2
	*z = *z * 2
	return *x, *y, *z
}
