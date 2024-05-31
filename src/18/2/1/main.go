// 18.2.1 Создайте программу, которая запускает 5 рутин, каждая из которых печатает свой порядковый номер 10 раз.
// Добиться такого результата за один вызов wg.Add.
//
// Используя старую версию Go
package main

import (
	"fmt"
	"sync"
)

func main() {
	h := make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for n := 0; n < 10; n++ {
				val := i
				fmt.Println(val)
				h[val] = h[val] + 1
			}
		}(i)
	}
	wg.Wait()
	for key, v := range h {
		fmt.Println(key, "выведено", v, "раз")
	}
}
