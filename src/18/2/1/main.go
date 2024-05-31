package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for n := 0; n < 10; n++ {
				fmt.Println(i)
			}
		}()
	}
	wg.Wait()
}
