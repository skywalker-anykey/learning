package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	insertionSort(ar)
	fmt.Println(ar)
}

func insertionSort(ar []int) {
	// проходим со второго элемента до конца
	for i := 1; i < len(ar); i++ {

		// текущий элемент
		key := ar[i]
		// индекс начального элемента "слева"
		j := i - 1

		// пока слева есть элементы или элемент слева больше текущего
		for j >= 0 && ar[j] > key {
			// сдвигаем элемент
			ar[j+1] = ar[j]
			// следующий элемент "слева"
			j--
		}
		// устанавливаем текущий элемент на его позицию
		ar[j+1] = key
	}
}
