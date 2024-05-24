// Перепишите приведённый выше пример со счётчиком из основного текста,
// но вместо примитивов из пакета atomic используйте условную переменную и
// попробуйте реализовать динамическую проверку достижения конечного значения счётчиком.
package main

import (
	"fmt"
	"sync"
)

func main() {
	const step int64 = 1
	const interationAmount int64 = 100

	var counter int64
	var c = sync.NewCond(&sync.Mutex{})

	increment := func() {
		for {
			c.L.Lock()
			c.Wait()
			counter = counter + step
			fmt.Println(counter)
			c.L.Unlock()
		}
	}

	go increment()
	go increment()
	go increment()
	go increment()
	go increment()
	go increment()
	go increment()
	go increment()

	for {
		c.L.Lock()
		if counter >= interationAmount {
			break
		}
		c.Signal()
		c.L.Unlock()
	}

	fmt.Println(counter)
}

//func main() {
//	var counter int64 = 0
//	var c = sync.NewCond(&sync.Mutex{})
//	increment := func(i int) {
//		atomic.AddInt64(&counter, step)
//		if i == interationAmount {
//			c.Signal()
//		}
//	}
//	for i := 1; i <= interationAmount; i++ {
//		go increment(i)
//	}
//	c.L.Lock()
//	c.Wait()
//	c.L.Unlock()
//	fmt.Println(counter)
//}
