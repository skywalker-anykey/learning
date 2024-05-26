package main

import (
	"fmt"
	"sync"
)

const step int64 = 1
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})

	// Запускаем 8 горутин
	for i := 0; i < 8; i++ {
		go func() {
			for {
				c.L.Lock()
				if counter == int64(interationAmount) {
					c.Signal()
				} else {
					counter = counter + step
				}
				c.L.Unlock()
			}
		}()
	}

	// Ждем "сигнал"
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
