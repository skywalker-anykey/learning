package main

import (
	"fmt"
	"time"
)

func main() {
	const timer = time.Second / 5
	var n, c int
	fmt.Scan(&n)

	go func() {
		for {
			fmt.Println(c)
			time.Sleep(timer)
		}
	}()

	for i := 0; i <= n; i++ {
		time.Sleep(1 * time.Second)
		c++
	}
}
