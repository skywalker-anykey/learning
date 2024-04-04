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
	ar := make([]int, 51)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println(ar)
	fmt.Println(bubbleSortRecursive(ar, len(ar)))
}

func bubbleSort(ar []int) {
	flag := false
	for !flag {
		flag = true
		for i := 1; i < len(ar); i++ {
			if ar[i-1] > ar[i] {
				ar[i-1], ar[i] = ar[i], ar[i-1]
				flag = false
			}
		}
	}
}

func bubbleSortReversed(ar []int) {
	i := 0
	j := len(ar) - 1

	for i < j {
		ar[i], ar[j] = ar[j], ar[i]
		i++
		j--
	}
}

func bubbleSortRecursive(ar []int, size int) []int {
	if size == 1 {
		return ar
	}

	for i := 1; i < size; i++ {
		if ar[i-1] > ar[i] {
			ar[i-1], ar[i] = ar[i], ar[i-1]
		}
	}

	bubbleSortRecursive(ar, size-1)

	return ar
}
